package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type RequestData struct {
	Message string `json:"message"`
}

type ResponseData struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func main() {
	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Only POST requests are allowed", http.StatusMethodNotAllowed)
			return
		}

		var requestData RequestData
		err := json.NewDecoder(r.Body).Decode(&requestData)
		if err != nil {
			http.Error(w, "Error decoding JSON", http.StatusBadRequest)
			return
		}

		if requestData.Message == "" {
			responseData := ResponseData{
				Status:  "400",
				Message: "Некорректное JSON-сообщение",
			}
			http.Error(w, "Invalid JSON message", http.StatusBadRequest)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(responseData)
			return
		}

		fmt.Printf("Received message: %s\n", requestData.Message)

		responseData := ResponseData{
			Status:  "success",
			Message: "Данные успешно приняты",
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(responseData)
	})

	fmt.Println("Server is running on :8080")
	http.ListenAndServe(":8080", nil)
}
