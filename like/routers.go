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
}
