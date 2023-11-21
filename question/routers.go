package question

import "net/http"

func QuestionRouters(mux *http.ServeMux) {

	// GET
	mux.HandleFunc("/question/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			ListQuestion(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
}
