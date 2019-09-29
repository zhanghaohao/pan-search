package movie

import (
	"net/http"
	"html/template"
	logger "util/log"
)

func Index(w http.ResponseWriter, r *http.Request)  {
	t, err := template.ParseFiles("src/template/movieindex.html", "src/template/public/header.html", "src/template/public/footer.html", "src/template/public/navi.html", "src/template/public/pansearcher.html", "src/template/public/magnetsearcher.html")
	if err != nil {
		logger.Error.Println("Canot find movieindex.html, ", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Page not found"))
		return
	}
	t.ExecuteTemplate(w, "movieindex", nil)
}
