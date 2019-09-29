package router

import (
	"net/http"
	"baidupan/pansearch"
	"baidupan/download"
	"magnet/magnetsearch"
	"movie"
)

func Router(mux *http.ServeMux) *http.ServeMux {
	mux.HandleFunc("/", pansearch.Index)
	mux.Handle("/static/", http.FileServer(http.Dir("./src")))
	mux.HandleFunc("/movie/index.html", movie.Index)
	mux.HandleFunc("/pan/index.html", pansearch.Index)
	mux.HandleFunc("/pan/search/", pansearch.Search)
	mux.HandleFunc("/pan/search/getrbdps", pansearch.GetBdps)
	mux.HandleFunc("/pan/search/getid", pansearch.GetId)
	mux.HandleFunc("/pan/search/bdp/", pansearch.SearchBdp)
	mux.HandleFunc("/pan/download", download.Download)
	mux.HandleFunc("/magnet/index.html", magnetsearch.Index)
	mux.HandleFunc("/magnet/search/", magnetsearch.Search)
	mux.HandleFunc("/magnet/getmagnets", magnetsearch.GetMagnets)
	return mux
}
