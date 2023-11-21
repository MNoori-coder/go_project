package question

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

func ListQuestion(w http.ResponseWriter, r *http.Request) {
	var questions []SchemaQuestion

	db, err := sql.Open("sqlite3", "my_db.db")
	if err != nil {
		http.Error(w, "Can not connect to db\n"+err.Error(), http.StatusBadRequest)
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM question")
	if err != nil {
		http.Error(w, "Can not execute query\n"+err.Error(), http.StatusBadRequest)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var question SchemaQuestion
		err := rows.Scan(&question.ID, &question.Title, &question.Content)
		if err != nil {
			http.Error(w, "Can not scan row data\n"+err.Error(), http.StatusBadRequest)
			return
		}

		questions = append(questions, question)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, "Error during iteration\n"+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := map[string]interface{}{
		"data":    questions,
		"success": true,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Error encoding response to JSON\n", http.StatusInternalServerError)
		return
	}
}

func CreateQuestion(w http.ResponseWriter, r *http.Request) {
	var question SchemaQuestion
	err := json.NewDecoder(r.Body).Decode(&question)
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

	_, err = db.Exec("INSERT INTO question (title, content) VALUES (?, ?)", question.Title, question.Content)
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

func UpdateQuestion(w http.ResponseWriter, r *http.Request) {
	vars := strings.Split(strings.TrimSuffix(r.URL.Path, "/"), "/")
	questionIDStr := vars[len(vars)-1]

	questionID, err := strconv.Atoi(questionIDStr)
	if err != nil {
		http.Error(w, "Invalid ticket ID\n"+err.Error(), http.StatusBadRequest)
		return
	}

	var question SchemaQuestion
	err = json.NewDecoder(r.Body).Decode(&question)
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

	_, err = db.Exec("UPDATE question SET title=?, content=? WHERE id=?", question.Title, question.Content, questionID)
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

func DeleteQuestion(w http.ResponseWriter, r *http.Request) {
	vars := strings.Split(strings.TrimSuffix(r.URL.Path, "/"), "/")
	questionIDStr := vars[len(vars)-1]

	questionID, err := strconv.Atoi(questionIDStr)
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

	_, err = db.Exec("DELETE FROM question WHERE id=?", questionID)
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

func RetrieveQuestion(w http.ResponseWriter, r *http.Request) {
	vars := strings.Split(strings.TrimSuffix(r.URL.Path, "/"), "/")
	questionIDStr := vars[len(vars)-1]

	questionID, err := strconv.Atoi(questionIDStr)
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

	row := db.QueryRow("SELECT * FROM question WHERE id=?", questionID)

	var question SchemaQuestion
	err = row.Scan(&question.ID, &question.Title, &question.Content)
	if err != nil {
		http.Error(w, "Error retrieving ticket\n"+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := map[string]interface{}{
		"data":    question,
		"success": true,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Error encoding response to JSON\n", http.StatusInternalServerError)
		return
	}
}
