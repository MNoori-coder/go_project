package question_category

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

func ListQuestionCategory(w http.ResponseWriter, r *http.Request) {
	var question_categories []SchemaQuestionCategory

	db, err := sql.Open("sqlite3", "my_db.db")
	if err != nil {
		http.Error(w, "Can not connect to db\n"+err.Error(), http.StatusBadRequest)
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM question_category")
	if err != nil {
		http.Error(w, "Can not execute query\n"+err.Error(), http.StatusBadRequest)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var question_category SchemaQuestionCategory
		err := rows.Scan(&question_category.ID, &question_category.Name, &question_category.Content)
		if err != nil {
			http.Error(w, "Can not scan row data\n"+err.Error(), http.StatusBadRequest)
			return
		}

		question_categories = append(question_categories, question_category)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, "Error during iteration\n"+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := map[string]interface{}{
		"data":    question_categories,
		"success": true,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Error encoding response to JSON\n", http.StatusInternalServerError)
		return
	}
}
