package like

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

func ListLike(w http.ResponseWriter, r *http.Request) {
	var likes []SchemaLike

	db, err := sql.Open("sqlite3", "my_db.db")
	if err != nil {
		http.Error(w, "Can not connect to db\n"+err.Error(), http.StatusBadRequest)
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM like")
	if err != nil {
		http.Error(w, "Can not execute query\n"+err.Error(), http.StatusBadRequest)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var like SchemaLike
		err := rows.Scan(&like.ID, &like.UserID, &like.TicketID, &like.IsLike)
		if err != nil {
			http.Error(w, "Can not scan row data\n"+err.Error(), http.StatusBadRequest)
			return
		}

		likes = append(likes, like)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, "Error during iteration\n"+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := map[string]interface{}{
		"data":    likes,
		"success": true,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Error encoding response to JSON\n", http.StatusInternalServerError)
		return
	}
}
