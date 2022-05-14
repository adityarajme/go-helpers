package golang_helpers

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
	"time"
)

//Struct for the template
type Pages struct {
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

func SetCORSHeaders(w http.ResponseWriter) http.ResponseWriter {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Authorization, Accept, Content-Type, Content-Length, Accept-Encoding, X-Auth-Token,  Access-Control-Allow-Origin")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	return w
}

//HTML templates
var tmpl_page = template.Must(template.ParseFiles("/var/www/location_to_html/page/index.html", "/var/www/location_to_html/page/list.html", "/var/www/location_to_html/page/edit.html"))

//Code to start the server and handle the requests
func main() {
	adminf := http.NewServeMux()
	adminf.HandleFunc("/api/pages", Page_Index)
	adminf.HandleFunc("/api/page_edit", Page_Edit)
	adminf.HandleFunc("/api/page_insert", Page_Insert)
	adminf.HandleFunc("/api/page_update", Page_Update)
	adminf.HandleFunc("/api/page_delete", Page_Delete)

	srv := &http.Server{
		ReadTimeout:       3 * time.Second,
		WriteTimeout:      6 * time.Second,
		Addr:              "localhost:9000",
		IdleTimeout:       15 * time.Second,
		ReadHeaderTimeout: 3 * time.Second,
		Handler:           adminf,
	}

	log.Panic(srv.ListenAndServe())
}

//indivisual pages
func Page_Index(w http.ResponseWriter, r *http.Request) {
	res := []Pages{}
	err := tmpl_page.ExecuteTemplate(w, "Page_Index", res)
	if err != nil {
		log.Println(err)
	}
}

func Page_Edit(w http.ResponseWriter, r *http.Request) {
	res := Pages{}
	err := tmpl_page.ExecuteTemplate(w, "Page_Edit", res)
	if err != nil {
		log.Println(err)
	}
}

func Page_Insert(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Page_Insert")
	if err != nil {
		log.Println(err)
	}
}

func Page_Update(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Page_Update")
	if err != nil {
		log.Println(err)
	}
}

func Page_Delete(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Page_Delete")
	if err != nil {
		log.Println(err)
	}
}
