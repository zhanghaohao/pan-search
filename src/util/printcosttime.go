package util

import (
	"time"
	logger "util/logger"
	"runtime"
)

func PrintCostTime(start time.Time)  {
	cost := time.Since(start)
	pc, _, _, _ := runtime.Caller(1)
	callerFunc := runtime.FuncForPC(pc).Name()
	logger.Info.Printf("%s cost: %s", callerFunc, cost)
}
