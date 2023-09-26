package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

type Data struct {
	Data string `json:"data"`
}

func AddData(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		// Parse JSON data from the request
		var data Data
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&data); err != nil {
			http.Error(w, "Failed to parse JSON data", http.StatusBadRequest)
			return
		}

		// Insert the data into the database
		_, err := db.Exec(fmt.Sprintf("INSERT INTO %s (%s) VALUES %s"), "mytable", "data", data.Data)
		if err != nil {
			http.Error(w, "Failed to insert data into the database", http.StatusInternalServerError)
			return
		}

		fmt.Println("testing")

		w.WriteHeader(http.StatusCreated)
	}
}

func GetData(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		rows, err := db.Query("SELECT data FROM mytable")
		if err != nil {
			http.Error(w, "Failed to query data from the database", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var dataItems []Data
		for rows.Next() {
			var data *Data
			if err := rows.Scan(&data.Data); err != nil {
				http.Error(w, "Failed to scan data from the database", http.StatusInternalServerError)
				return
			}
			dataItems = append(dataItems, *data)
		}

		jsonData, err := json.Marshal(dataItems)
		if err != nil {
			http.Error(w, "Failed to marshal data to JSON", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)
	}
}
