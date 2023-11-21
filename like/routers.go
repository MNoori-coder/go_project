package like

import "net/http"

func LikeRouters(mux *http.ServeMux) {

	// GET
	mux.HandleFunc("/like/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			ListLike(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// POST
	mux.HandleFunc("/like/insert/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			CreateLike(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// PUT
	mux.HandleFunc("/like/update/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPut {
			UpdateLike(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// DELETE
	mux.HandleFunc("/like/delete/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodDelete {
			DeleteLike(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// GET
	mux.HandleFunc("/like/id/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			RetrieveLike(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
}
