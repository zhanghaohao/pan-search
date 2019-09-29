package main

import (
	"net/http"
	"net"
	"os"
	"router"
	logger "util/log"
	"util/coreconfig"
)

func main()  {
	_, err := coreconfig.YamlParser()
	if err != nil {
		panic(err)
	}
	mux := http.NewServeMux()
	mux = router.Router(mux)
	server := &http.Server{Handler: mux}
	l, err := net.Listen("tcp4", ":8080")
	if err != nil {
		logger.Error.Printf("%v", err)
	}
	logger.Info.Printf("Starting Baidupan Search ...")
	err = server.Serve(l)
	if err != nil {
		logger.Error.Printf("%v", err)
		os.Exit(1)
	}
}
