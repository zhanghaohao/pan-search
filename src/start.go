package main

import (
	"net/http"
	"net"
	"os"
	"route"
	"util/logger"
)

func main()  {
	mux := http.NewServeMux()
	mux, err := route.RegisterRoutes(mux)
	if err != nil {
		panic(err)
	}
	server := &http.Server{Handler: mux}
	l, err := net.Listen("tcp4", ":8080")
	if err != nil {
		logger.Error.Printf("%v", err)
	}
	logger.Info.Printf("starting service...")
	err = server.Serve(l)
	if err != nil {
		logger.Error.Printf("%v", err)
		os.Exit(1)
	}
}
