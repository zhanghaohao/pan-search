package log

import (
	"log"
	"os"
)

var (
	Error *log.Logger
	Warn *log.Logger
	Info *log.Logger
)

func init()  {
	Error = log.New(os.Stderr, "[Error]", log.Ldate | log.Ltime | log.Lshortfile)
	Warn = log.New(os.Stdout, "[Warn]", log.Ldate | log.Ltime | log.Lshortfile)
	Info = log.New(os.Stdout, "[Info]", log.Ldate | log.Ltime | log.Lshortfile)

}
