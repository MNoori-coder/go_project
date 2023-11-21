package like

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
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

func CreateLike(w http.ResponseWriter, r *http.Request) {
	var like SchemaLike
	err := json.NewDecoder(r.Body).Decode(&like)
	if err != nil {
		http.Error(w, "Error decoding request body\n"+err.Error(), http.StatusBadRequest)
		return
	}

	db, err := sql.Open("sqlite3", "my_db.db")
	if err != nil {
		http.Error(w, "Can not connect to db\n"+err.Error(), http.StatusBadRequest)
		return
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO like (user_id, ticket_id, is_like) VALUES (?, ?, ?)", like.UserID, like.TicketID, like.IsLike)
	if err != nil {
		http.Error(w, "Error executing insert query\n"+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := map[string]interface{}{
		"success": true,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Error encoding response to JSON\n", http.StatusInternalServerError)
		return
	}
}

func UpdateLike(w http.ResponseWriter, r *http.Request) {
	vars := strings.Split(strings.TrimSuffix(r.URL.Path, "/"), "/")
	likeIDStr := vars[len(vars)-1]

	likeID, err := strconv.Atoi(likeIDStr)
	if err != nil {
		http.Error(w, "Invalid ticket ID\n"+err.Error(), http.StatusBadRequest)
		return
	}

	var like SchemaLike
	err = json.NewDecoder(r.Body).Decode(&like)
	if err != nil {
		http.Error(w, "Error decoding request body\n"+err.Error(), http.StatusBadRequest)
		return
	}

	db, err := sql.Open("sqlite3", "my_db.db")
	if err != nil {
		http.Error(w, "Can not connect to db\n"+err.Error(), http.StatusBadRequest)
		return
	}
	defer db.Close()

	_, err = db.Exec("UPDATE like SET user_id=?, ticket_id=?, is_like=? WHERE id=?", like.UserID, like.TicketID, like.IsLike, likeID)
	if err != nil {
		http.Error(w, "Error executing update query\n"+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := map[string]interface{}{
		"success": true,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Error encoding response to JSON\n", http.StatusInternalServerError)
		return
	}
}

func DeleteLike(w http.ResponseWriter, r *http.Request) {
	vars := strings.Split(strings.TrimSuffix(r.URL.Path, "/"), "/")
	likeIDStr := vars[len(vars)-1]

	likeID, err := strconv.Atoi(likeIDStr)
	if err != nil {
		http.Error(w, "Invalid ticket ID\n"+err.Error(), http.StatusBadRequest)
		return
	}

	db, err := sql.Open("sqlite3", "my_db.db")
	if err != nil {
		http.Error(w, "Can not connect to db\n"+err.Error(), http.StatusBadRequest)
		return
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM like WHERE id=?", likeID)
	if err != nil {
		http.Error(w, "Error executing delete query\n"+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := map[string]interface{}{
		"success": true,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Error encoding response to JSON\n", http.StatusInternalServerError)
		return
	}
}

func RetrieveLike(w http.ResponseWriter, r *http.Request) {
	vars := strings.Split(strings.TrimSuffix(r.URL.Path, "/"), "/")
	likeIDStr := vars[len(vars)-1]

	likeID, err := strconv.Atoi(likeIDStr)
	if err != nil {
		http.Error(w, "Invalid ticket ID\n"+err.Error(), http.StatusBadRequest)
		return
	}

	db, err := sql.Open("sqlite3", "my_db.db")
	if err != nil {
		http.Error(w, "Can not connect to db\n"+err.Error(), http.StatusBadRequest)
		return
	}
	defer db.Close()

	row := db.QueryRow("SELECT * FROM like WHERE id=?", likeID)

	var like SchemaLike
	err = row.Scan(&like.ID, &like.UserID, &like.TicketID, &like.IsLike)
	if err != nil {
		http.Error(w, "Error retrieving ticket\n"+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := map[string]interface{}{
		"data":    like,
		"success": true,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Error encoding response to JSON\n", http.StatusInternalServerError)
		return
	}
}
