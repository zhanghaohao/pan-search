package weixin

import (
	"util/httpclient"
	"util/logger"
	"encoding/json"
	"github.com/garyburd/redigo/redis"
	"fmt"
	"plugin/db"
	"time"
	"net/http"
	"github.com/gorilla/websocket"
	"crypto/sha1"
	"io"
	"strings"
	"sort"
	"io/ioutil"
	"encoding/xml"
	"regexp"
)

const (
	AccessTokenKey = "accessToken"
)

var WebSocketCount = 0

var TicketChan = make(chan string)

type accessToken struct {
	AccessToken 				string 						`json:"access_token"`
	ExpiresIn 					int 						`json:"expires_in"` 			// 2 hours by default
}

type ticketRespBody struct {
	Ticket 						string 						`json:"ticket"`
	ExpireSeconds 				int64 						`json:"expire_seconds"`
	Url 						string 						`json:"url"`
}

type ticketRespBodyError struct {
	ErrCode 					int64 						`json:"errcode"`
	ErrMsg 						string 						`json:"errmsg"`
}

type ticketPostBody struct {
	ExpireSeconds 				int64 						`json:"expire_seconds"`
	ActionName 					string 						`json:"action_name"`
	ActionInfo 					ticketPostBodyScene 		`json:"action_info"`
}

type ticketPostBodyScene struct {
	Scene 						ticketPostBodySceneID 		`json:"scene"`
}

type ticketPostBodySceneID struct {
	SceneID 					int64 						`json:"scene_id"`
}

type qrcode struct {
	id 							int
	tmpTicket 					string
	createAt					string
	validMinutes 				int
}

type qrcodeScanEvent struct {
	FromUser 					string 						`xml:"FromUserName"`
	CreateTime 					string 						`xml:"CreateTime"`
	EventKey 					string 						`xml:"EventKey"`
	Ticket 						string 						`xml:"Ticket"`
	MsgType						string 						`xml:"MsgType"`
	Event 						string 						`xml:"Event"`
}

type WeiXinConfig struct {
	AppID 		string 				`yaml:"appID"`
	AppSecret	string 				`yaml:"appSecret"`
	Token  		string 				`yaml:"token"`
}

type weixin struct {
	mysqlEngine 				*db.MysqlEngine
	redisPool 					*db.RedisPool
	weixinConfig 				*WeiXinConfig
}

type WeiXinFactory interface {
	GetTmpTicket(w http.ResponseWriter, r *http.Request)
	EventReceiver(w http.ResponseWriter, r *http.Request)
	WebSocketHandler(w http.ResponseWriter, r *http.Request)
}

func New(mysqlEngine *db.MysqlEngine, redisPool *db.RedisPool, weixinConfig *WeiXinConfig) WeiXinFactory {
	return &weixin{
		mysqlEngine: mysqlEngine,
		redisPool: redisPool,
		weixinConfig: weixinConfig,
	}
}

func (c *weixin) genAccessToken() (ak *accessToken, err error) {
	ak = new(accessToken)
	method := "GET"
	url := "https://api.weixin.qq.com/cgi-bin/token"
	header := map[string]string{"Content-Type": "application/json"}
	params := map[string]string{
		"grant_type": "client_credential",
		"appid": c.weixinConfig.AppID,
		"secret": c.weixinConfig.AppSecret,
	}
	request, err := httpclient.NewRequest(url, header, method, params, "")
	if err != nil {
		return
	}
	resp, err := httpclient.NewClient().DoForBody(request)
	if err != nil {
		logger.Error.Println(err)
		return
	}
	err = json.Unmarshal(resp, ak)
	if err != nil {
		logger.Error.Println(err)
		logger.Error.Println(string(resp))
		return
	}
	if len(ak.AccessToken) == 0 {
		err = fmt.Errorf(string(resp))
		logger.Error.Println(err)
		return
	}
	return
}

func (c *weixin) setAccessToken(ak *accessToken) (err error) {
	cli := c.redisPool.Get()
	defer cli.Close()
	akByte, err := json.Marshal(*ak)
	if err != nil {
		logger.Error.Println(err)
		return
	}
	_, err = cli.Do("SET", AccessTokenKey, string(akByte))
	if err != nil {
		logger.Error.Println(err)
		return
	}
	_, err = cli.Do("EXPIRE", AccessTokenKey, ak.ExpiresIn)
	if err != nil {
		logger.Error.Println(err)
		return
	}
	return
}

func (c *weixin) delAccessToken() (err error)  {
	cli := c.redisPool.Get()
	defer cli.Close()
	_, err = cli.Do("DEL", AccessTokenKey)
	if err != nil {
		logger.Error.Println(err)
		return
	}
	return
}

func (c *weixin) readAccessToken() (ak *accessToken, err error) {
	ak = new(accessToken)
	cli := c.redisPool.Get()
	defer cli.Close()
	value, err := redis.Bytes(cli.Do("GET", AccessTokenKey))
	if err != nil {
		err = fmt.Errorf("access token has expired")
		logger.Info.Println(err)
		return
	}
	err = json.Unmarshal(value, ak)
	if err != nil {
		logger.Error.Println(err)
		return
	}
	return
}

func genTmpTicket(accessToken string, sceneID int64) (ticket *ticketRespBody, err error) {
	// generate new temp ticket
	ticket = new(ticketRespBody)
	url := "https://api.weixin.qq.com/cgi-bin/qrcode/create"
	method := "POST"
	header := map[string]string{"Content-Type": "application/json"}
	params := map[string]string{
		"access_token": accessToken,
	}
	postBody := ticketPostBody{
		ExpireSeconds: 600, 			// 10 minutes
		ActionName: "QR_SCENE",
		ActionInfo: ticketPostBodyScene{
			Scene: ticketPostBodySceneID{
				SceneID: sceneID,
			},
		},
	}
	postBodyByte, err := json.Marshal(postBody)
	if err != nil {
		logger.Error.Println(err)
		return
	}
	request, err := httpclient.NewRequest(url, header, method, params, string(postBodyByte))
	if err != nil {
		return
	}
	resp, err := httpclient.NewClient().DoForBody(request)
	if err != nil {
		logger.Error.Println(err)
		return
	}
	//logger.Info.Println(string(resp))
	err = json.Unmarshal(resp, ticket)
	if err != nil {
		logger.Error.Println(err)
		return
	}
	// the response may be error
	if len(ticket.Ticket) == 0 {
		err = fmt.Errorf(string(resp))
		logger.Error.Println(err)
		return
	}
	return
}

func (c *weixin) newSceneID() (sceneID *int64, err error) {
	/*
	set ticket as expired.
	insert a new record, and return the ID.
	  */
	var qrcodes []db.Qrcode
	err = c.mysqlEngine.Where("isExpired = 0").Cols("id", "tmpTicket", "createAt", "validMinutes").Find(&qrcodes)
	if err != nil {
		logger.Error.Println(err)
		return
	}
	for _, qrcode := range qrcodes {
		// skip record with tmpTicket empty
		if len(qrcode.TmpTicket) == 0 {
			continue
		}
		// set ticket as expired if expired
		createTime, err := time.Parse("2006-01-02 15:04:05", qrcode.CreateAt)
		if err != nil {
			logger.Error.Println(err)
			return nil, nil
		}
		if createTime.Add(time.Minute * time.Duration(qrcode.ValidMinutes)).After(time.Now()) == true {
			qrcode.IsExpired = true
			_, err := c.mysqlEngine.ID(qrcode.Id).Update(&qrcode)
			if err != nil {
				logger.Error.Println(err)
				return nil, nil
			}
		}
	}
	// insert a new record
	createAt := time.Now().Format("2006-01-02 15:04:05")
	record := &db.Qrcode{
		CreateAt: createAt,
		IsScanned: false,
		IsExpired: false,
		IsLocked: false,
	}
	_, err = c.mysqlEngine.InsertOne(record)
	if err != nil {
		logger.Error.Println(err)
		return
	}
	lastInsertID := record.Id
	sceneID = &lastInsertID
	return
}

func (c *weixin) getAccessToken() (ak *accessToken, err error) {
	/*
	get access token from redis first,
	generate access token if cannot get from redis, then set in redis.
	  */
	ak, err = c.readAccessToken()
	if err != nil {
		// generate access token
		ak, err = c.genAccessToken()
		if err != nil {
			return nil, err
		}
		//logger.Info.Println(*ak)
		err := c.setAccessToken(ak)
		if err != nil {
			logger.Error.Println(err)
			return nil, err
		}
	}
	return
}

func (c *weixin) resetAccessTokenByErr(errBody error) (err error) {
	/*
	if error match the following, delete access token in redis
	"errcode":40001,"errmsg":"invalid credential, access_token is invalid or not latest hints
	  */
	isMatch, err := regexp.Match("invalid credential", []byte(errBody.Error()))
	if err != nil {
		logger.Error.Println(err)
		return
	}
	if isMatch == true {
		err = c.delAccessToken()
		if err != nil {
			logger.Error.Println(err)
			return err
		}
	} else {
		return errBody
	}
	return
}

func (c *weixin) getTmpTicket() (ticket string, err error) {
	/*
	generate sceneID first, then generate ticket, write into mysql in the end.
	 */
	sceneID, err := c.newSceneID()
	if err != nil {
		return
	}
	logger.Info.Println("sceneID: ", *sceneID)
	// get access token
	ak, err := c.getAccessToken()
	if err != nil {
		return
	}
	tmpTicket, err := genTmpTicket(ak.AccessToken, *sceneID)
	if err != nil {
		err = c.resetAccessTokenByErr(err)
		if err != nil {
			logger.Error.Println(err)
			return "", err
		} else {
			return c.getTmpTicket()
		}
	}
	//logger.Info.Printf("%+v", *tmpTicket)
	// write ticket into mysql
	validMinutes := tmpTicket.ExpireSeconds/60
	record := &db.Qrcode{
		TmpTicket: tmpTicket.Ticket,
		ValidMinutes: int(validMinutes),
		IsLocked: true,
	}
	_, err = c.mysqlEngine.ID(*sceneID).Update(record)
	if err != nil {
		logger.Error.Println(err)
		return
	}
	ticket = tmpTicket.Ticket
	return
}

func (c *weixin) GetTmpTicket(w http.ResponseWriter, r *http.Request)  {
	ticket, err := c.getTmpTicket()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("获取不到ticket"))
		return
	}
	w.Write([]byte(ticket))
	return
}

func (c *weixin) verifySignature(signature, timestamp, nonce string) (bool) {
	token := c.weixinConfig.Token
	tmpSlice := []string{timestamp, nonce, token}
	sort.Strings(tmpSlice)
	tmpStr := strings.Join(tmpSlice, "")
	//logger.Info.Println(tmpStr)
	t := sha1.New()
	io.WriteString(t, tmpStr)
	tmpStr = fmt.Sprintf("%x", t.Sum(nil))
	//logger.Info.Println(tmpStr)
	if tmpStr == signature {
		return true
	} else {
		return false
	}
}

func (c *weixin) EventReceiver(w http.ResponseWriter, r *http.Request)  {
	/*
	verify this url is valid
	  */
	if r.Method == "GET"  {
		logger.Info.Println("this is GET")
		signature := r.URL.Query().Get("signature")
		timestamp := r.URL.Query().Get("timestamp")
		nonce := r.URL.Query().Get("nonce")
		echoStr := r.URL.Query().Get("echostr")
		logger.Info.Println("signature: ", signature)
		logger.Info.Println("timestamp: ", timestamp)
		logger.Info.Println("nonce: ", nonce)
		logger.Info.Println("echostr: ", echoStr)
		ifValid := c.verifySignature(signature, timestamp, nonce)
		if ifValid == true {
			logger.Info.Println("verify success")
			w.Write([]byte(echoStr))
			return
		} else {
			logger.Info.Println("verify failed")
			w.Write([]byte("这是攻击行为，请立即停止，否则将上报国家信息安全部门！"))
			return
		}
	}
	/*
	handle event when temp qrcode is scanned by user
	  */
	if r.Method == "POST" {
		logger.Info.Println("this is POST")
		body, err := ioutil.ReadAll(r.Body)
		//_, err := r.Body.Read(body)
		if err != nil {
			logger.Error.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("无法读取POST参数"))
			return
		}
		//logger.Info.Println(string(body))
		var postParams qrcodeScanEvent
		err = xml.Unmarshal(body, &postParams)
		if err != nil {
			logger.Error.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("无法序列化POST参数"))
			return
		}
		// donot handle if the event is not qrcode scan
		if postParams.MsgType != "event" {
			logger.Info.Printf("%+v", postParams)
			return
		}
		if postParams.Event != "subscribe" && postParams.Event != "unsubscribe" && postParams.Event != "SCAN" {
			logger.Info.Printf("%+v", postParams)
			return
		}
		logger.Info.Println(postParams.Event)
		// record params
		record := &db.Qrcode{
			FromUser: postParams.FromUser,
			ScannedAt: postParams.CreateTime,
			IsScanned: true,
			IsLocked: false,
		}
		_, err = c.mysqlEngine.Update(record, &db.Qrcode{TmpTicket: postParams.Ticket})
		if err != nil {
			logger.Error.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		// send ticket to global channel, so that other function can handle it
		TicketChan <- postParams.Ticket
		w.WriteHeader(http.StatusOK)
		return
	}
}

func (c *weixin) WebSocketHandler(w http.ResponseWriter, r *http.Request) {
	var upgrader = websocket.Upgrader{
		ReadBufferSize:   1024,
		WriteBufferSize:  1024,
		HandshakeTimeout: 5 * time.Second,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logger.Error.Println(err)
		return
	}
	_, msg, err := ws.ReadMessage()
	if err != nil {
		logger.Error.Println(err)
		return
	}
	WebSocketCount ++
	defer func() {
		WebSocketCount --
	}()
	wantedTicket := string(msg)
	//logger.Info.Println("wantedTicket: ", wantedTicket)
	var ticketCounter = make(map[string]int)
	for {
		select {
		case ticket := <- TicketChan:
			// if this goroutine gets the same ticket more than one time, then abanden the ticket. because the related websocket connection was timeout.
			//logger.Info.Println("ticket: ", ticket)
			ticketCounter[ticket] ++
			//logger.Info.Println("ticketCounter: ", ticketCounter)
			if ticketCounter[ticket] > 1 {
				continue
			} else {
				// throw it back to channel and wait for some time, so that other goroutine can get it.
				if WebSocketCount > 1 {
					TicketChan <- ticket
				}
			}
			if ticket != wantedTicket {
				//logger.Info.Println("ticket is not wanted")
				time.Sleep(100 * time.Millisecond)
			} else {
				//logger.Info.Println("ticket is wanted")
				err := ws.WriteMessage(websocket.TextMessage, []byte("scanned"))
				if err != nil {
					logger.Error.Println(err)
				}
				return
			}
		case <- time.After(10 * time.Minute):
			logger.Info.Println("websocket connection timeout.")
			logger.Info.Println(wantedTicket)
			err = ws.Close()
			if err != nil {
				logger.Error.Println(err)
			}
			return
		}
	}
}