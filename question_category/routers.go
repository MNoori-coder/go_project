package question_category

import "net/http"

func QuestionCategoryRouters(mux *http.ServeMux) {

	// GET
	mux.HandleFunc("/question_category/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			ListQuestionCategory(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
}
