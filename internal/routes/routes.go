package routes

import ( 
    "html/template"
    "net/http"
)

func NewRouter() http.Handler {
    mux := http.NewServeMux()
    mux.HandleFunc("/", indexHandler)

    mux.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("../../public"))))

    return mux

}
func indexHandler(w http.ResponseWriter, r *http.Request) {
    tmpl := template.Must(template.ParseFiles("../../templates/index.html"))
    tmpl.Execute(w, nil)
}
