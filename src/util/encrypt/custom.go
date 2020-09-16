package encrypt

import (
	"strconv"
	"strings"
	"util/logger"
)

const (
	salt int64 = 964
)

type mapper struct {
	num 				string
	str 				string
}

var table = []mapper{
	{
		num: "0",
		str: "e",
	},
	{
		num: "1",
		str: "l",
	},
	{
		num: "2",
		str: "q",
	},
	{
		num: "3",
		str: "h",
	},
	{
		num: "4",
		str: "v",
	},
	{
		num: "5",
		str: "b",
	},
	{
		num: "6",
		str: "y",
	},
	{
		num: "7",
		str: "c",
	},
	{
		num: "8",
		str: "n",
	},
	{
		num: "9",
		str: "i",
	},
}

func forwardConv(num string) (str string) {
	for _, e := range table {
		if num == e.num {
			return e.str
		}
	}
	return
}

func backwardConv(str string) (num string) {
	for _, e := range table {
		if str == e.str {
			return e.num
		}
	}
	return
}

func EncryptID(rawID int64) (cookedID string, err error) {
	//id, err := strconv.ParseInt(rawID, 10, 64)
	//if err != nil {
	//	logger.Error.Println(err)
	//	return
	//}
	saltID := rawID * salt
	//logger.Info.Println(saltID)
	idStr := strconv.FormatInt(saltID,10)
	//logger.Info.Println(idStr)
	var convertedIDStr []string
	for i:=0;i<len(idStr);i++ {
		//logger.Info.Println(string(idStr[i]))
		m := forwardConv(string(idStr[i]))
		convertedIDStr = append(convertedIDStr, m)
	}
	cookedID = strings.Join(convertedIDStr, "")
	return
}

func DecryptID(rawID string) (cookedID string, err error) {
	var numSlice []string
	for i:=0;i<len(rawID);i++ {
		num := backwardConv(string(rawID[i]))
		numSlice = append(numSlice, num)
	}
	numStr := strings.Join(numSlice, "")
	id, err := strconv.ParseInt(numStr, 10, 64)
	if err != nil {
		logger.Error.Printf("invalid id %s, error: %v", rawID, err)
		return
	}
	idStr := id/salt
	cookedID = strconv.FormatInt(idStr, 10)
	return
}