package routes

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strings"
	"urlshortener/models"
	"urlshortener/utils"
)

func Init() *mux.Router {

	r := mux.NewRouter()

	return r
}

func logNecessary(r *http.Request) {
	log.Println("path", r.URL.Path)
	log.Println("scheme",r.URL.Scheme)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	logNecessary(r)
	log.Println("Index page get request")
	p := "vishal"
	utils.ExecuteTemplate(w, "index.html", p)
}

func registerURLHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	r.ParseForm()
	logNecessary(r)
	log.Println("Request to register a url received")
	for k, v := range r.Form {
		log.Println("key ",k)
		log.Println("value",strings.Join(v, ""))
	}
	models.RegisterURL(r.Form["originalurl"][0])
}

func urlGetHandler (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	//w.WriteHeader(http.StatusOK)
	logNecessary(r)
	params:=mux.Vars(r)
	url:=params["url"]
	log.Println("shortened url is ",url)
	original,err:=models.GetOriginalURL(url)
	if err!=nil{
		log.Println("url not found in routes.go",err)
	}else{
		log.Println("original url in routes.go", original)
		http.Redirect(w,r,"https://www."+original,http.StatusSeeOther)
	}
}

func Handle(r *mux.Router){
	r.HandleFunc("/", indexHandler)
	r.HandleFunc("/registerurl", registerURLHandler).Methods("POST")
	r.HandleFunc("/{url}",urlGetHandler).Methods("GET")
}