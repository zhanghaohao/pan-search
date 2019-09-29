package download

import (
	"net/http"
	"html/template"
	logger "util/log"
)

func Download(w http.ResponseWriter, r *http.Request)  {
	t, err := template.ParseFiles("src/template/download.html", "src/template/public/header.html", "src/template/public/footer.html", "src/template/public/navi.html", "src/template/public/pansearcher.html")
	if err != nil {
		logger.Error.Println("Canot find download.html, ", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Page not found"))
		return
	}
	t.ExecuteTemplate(w, "download", nil)
	return
}
