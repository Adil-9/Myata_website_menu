package render

import(
	"text/template"
	"net/http"
	"fmt"
)

func Homepage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		errorHandler(w, r, http.StatusMethodNotAllowed)
		return
	}
	tmpl, err := template.ParseFiles("webpages/main.html")
	if err != nil {
		errorHandler(w, r, http.StatusInternalServerError)
		return
	}
	if err := tmpl.Execute(w, ""); err != nil {
		errorHandler(w, r, http.StatusInternalServerError)
		return
	}
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		fmt.Fprint(w, "404 page not found")
	} else if status == http.StatusBadRequest {
		fmt.Fprint(w, "400 bad request")
	} else if status == http.StatusInternalServerError {
		fmt.Fprint(w, "500 internal server error")
	} else if status == http.StatusMethodNotAllowed {
		fmt.Fprint(w, "Method not allowed")
	}
}