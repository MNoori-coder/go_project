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

	// POST
	mux.HandleFunc("/question/insert/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			CreateQuestion(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// PUT
	mux.HandleFunc("/question/update/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPut {
			UpdateQuestion(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// DELETE
	mux.HandleFunc("/question/delete/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodDelete {
			DeleteQuestion(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// GET
	mux.HandleFunc("/question/id/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			RetrieveQuestion(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
}
