package weixin

import (
	"testing"
	"util/logger"
)

func Test_genTmpTicket(t *testing.T) {
	ak := ""
	sceneID := int64(1)
	ticket, err := genTmpTicket(ak, sceneID)
	if err != nil {
		t.Error(err)
	}
	logger.Info.Printf("%+v", ticket)
}

func TestWebSocketHandler(t *testing.T) {
	var ticketCounter = make(map[string]int)
	for i := 0; i < 10; i++ {
		ticketCounter["a"]++
		logger.Info.Println(ticketCounter)
	}
}
