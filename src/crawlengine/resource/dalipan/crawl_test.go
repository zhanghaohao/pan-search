package dalipan

import (
	"crawlengine/resource/common"
	"testing"
	"util/logger"
)

func Test_getMetaData(t *testing.T) {
	keyword := "美女"
	pageNumber := 1
	o := constructIPGenerator()
	ids, err := o.getIDs(keyword, pageNumber)
	if err != nil {
		t.Error(err)
	}
	for _, id := range ids {
		logger.Info.Println(id)
		metaData, err := o.getMetaData(id)
		if err != nil {
			t.Error(err)
		}
		logger.Info.Printf("%+v", metaData)
	}
}

func Test_getMetaDataPerKeyword(t *testing.T) {
	keyword := "美女"
	o, err := construct()
	if err != nil {
		t.Error(err)
	}
	err = o.getMetaDataPerKeyword(keyword)
	if err != nil {
		t.Error(err)
	}
}

func Test_getIDs(t *testing.T) {
	keyword := "美女"
	pageNumber := 1
	o := constructIPGenerator()
	ids, err := o.getIDs(keyword, pageNumber)
	if err != nil {
		t.Error(err)
	}
	logger.Info.Println(ids)
}

func Test_hasExisted(t *testing.T) {
	o, err := constructMysqlEngine()
	if err != nil {
		t.Error(err)
	}
	tests := []struct{
		metaData 	*common.BDP
		wanted 		bool
	}{
		{
			metaData: &common.BDP{
				Url:      "https://pan.baidu.com/s/1cHVyxq",
				Title:    "GTA专卖QQ1668618000",
				CTime:    "2016-07-02",
				Size:     "307566748287",
				Resource: "百度网盘",
			},
			wanted: true,
		},
		{
			metaData: &common.BDP{
				Url:   "https://pan.baidu.com/s/14FeFD",
				Title: "成都地铁１号线多站点临时停车　因乘客争抢上车.doc",
			},
			wanted: false,
		},
		{
			metaData: &common.BDP{
				Url: "https://pan.baidu.com/s/1dDTzTNr",
			},
			wanted: true,
		},
	}
	for _, e := range tests {
		exist, err := o.hasExisted(e.metaData)
		if err != nil {
			t.Error(err)
		}
		if exist != e.wanted {
			t.Errorf("wanted %v, but got %v", e.wanted, exist)
		}
	}

}

func Test_writeToMysql(t *testing.T) {
	url := "https://pan.baidu.com/s/1c0EPG1i"
	metaDatas := []*common.BDP{
		{
			Url:   url,
			Title: "test",
		},
	}
	o, err := constructMysqlEngine()
	if err != nil {
		t.Error(err)
	}
	err = o.writeToMysql(metaDatas)
	if err != nil {
		t.Error(err)
	}
}

func Test_ipGenerator_next(t *testing.T) {
	ipGenerator := constructIPGenerator()
	for i := 0; i < 500000; i++ {
		ip := ipGenerator.next()
		logger.Info.Println(i, ip)
	}
}

func Test_run(t *testing.T) {
	err := run()
	if err != nil {
		t.Error(err)
	}
}


