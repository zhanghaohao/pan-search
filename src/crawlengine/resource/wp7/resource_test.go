package wp7

import (
	"crawlengine/resource/common"
	"strconv"
	"testing"
	"util/logger"
)

func Test_getIDs(t *testing.T) {
	err := getIDs()
	if err != nil {
		t.Error(err)
	}
}

func Test_getMetaData(t *testing.T) {
	url := "https://wp7.net/share/file/1897417"
	resourceKind := Xinlang
	metaData, err := getMetaData(url, resourceKind)
	if err != nil {
		t.Error(err)
	}
	logger.Info.Printf("%+v", metaData)
}

func Test_assembleRedirectURL(t *testing.T) {
	url := "https://wp7.net/share/file/1897417"
	redirectURL := assembleRedirectURL(url)
	logger.Info.Printf(redirectURL)
}

func Test_getPanURL(t *testing.T) {
	redirectURL := "https://wp7.net/redirect/file?id=1897417"
	panURL, err := getPanURL(redirectURL)
	if err != nil {
		t.Error(err)
	}
	logger.Info.Println(panURL)
}

func Test_verifyPanURL(t *testing.T) {
	panURL := "http://vdisk.weibo.com/s/tckA9qQQQCB6a"
	isValid, resourceKind, err := verifyPanURL(panURL)
	if err != nil {
		t.Error(err)
	}
	if isValid == true {
		t.Errorf("want invalid, but get valid")
	}
	if resourceKind != Xinlang {
		t.Errorf("want xinglangpan, but get other")
	}
	panURL = "https://vdisk.weibo.com/s/vGB2PLMb-QTU"
	isValid, resourceKind, err = verifyPanURL(panURL)
	if err != nil {
		t.Error(err)
	}
	if isValid == false {
		t.Errorf("want valid, but get invalid")
	}
	if resourceKind != Xinlang {
		t.Errorf("want xinglangpan, but get other")
	}
	panURL = "https://pan.baidu.com/s/1miv1kj2"
	isValid, resourceKind, err = verifyPanURL(panURL)
	if err != nil {
		t.Error(err)
	}
	if isValid == false {
		t.Errorf("want valid, but get invalid")
	}
	if resourceKind != Baidu {
		t.Errorf("want xinglangpan, but get other")
	}
	panURL = "https://pan.baidu.com/share/link?shareid=4156794016&uk=2570243430&fid=916088277265528"
	isValid, resourceKind, err = verifyPanURL(panURL)
	if err != nil {
		t.Error(err)
	}
	if isValid == true {
		t.Errorf("want invalid, but get valid")
	}
	if resourceKind != Baidu {
		t.Errorf("want xinglangpan, but get other")
	}
}

func Test_writeMetaData(t *testing.T) {
	mysql, err := initData()
	if err != nil {
		t.Error(err)
	}
	metaData := common.BDP{
		Url:      "https://pan.baidu.com/s/1QSQupxQmDCxQ1TupCbG5eQ",
		Title:    "第一D血5：最后的血2019(1)",
		CTime:    "2019-12-08",
		Size:     strconv.Itoa(412940000),
		Category: "folder",
		Resource: string(Baidu),
	}
	err = mysql.writeMetaData(metaData)
	if err != nil {
		t.Error(err)
	}
}

func Test_handleSingleURL(t *testing.T) {
	mysql, err := initData()
	if err != nil {
		t.Error(err)
	}
	url := "https://wp7.net/share/file/1897417"
	err = mysql.handleSingleURL(url)
	if err != nil {
		t.Error(err)
	}
}

func Test_loadResources(t *testing.T) {
	err := loadResources()
	if err != nil {
		t.Error(err)
	}
}

