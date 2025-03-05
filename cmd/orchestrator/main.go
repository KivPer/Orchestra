package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"yandexGoCalc/internal/api"
	"yandexGoCalc/internal/models"
	"yandexGoCalc/internal/orchestrator"
)

func main() {
	// Создаем экземпляр оркестратора
	orc := orchestrator.NewOrchestrator()

	// Запускаем горутину для обработки результатов
	go orc.ProcessResults()

	// Регистрируем HTTP-обработчики
	http.HandleFunc("/api/v1/calculate", func(w http.ResponseWriter, r *http.Request) {
		api.CalculateHandler(w, r, orc)
	})

	http.HandleFunc("/api/v1/expressions", func(w http.ResponseWriter, r *http.Request) {
		// Логика для получения списка выражений
	})

	http.HandleFunc("/api/v1/expressions/", func(w http.ResponseWriter, r *http.Request) {
		// Логика для получения выражения по ID
	})

	http.HandleFunc("/internal/task", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			// Логика для выдачи задачи агенту
			task := orc.GetTask()
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(task)
		case http.MethodPost:
			// Логика для приема результата от агента
			var result models.Result
			if err := json.NewDecoder(r.Body).Decode(&result); err != nil {
				http.Error(w, `{"error": "Bad request"}`, http.StatusBadRequest)
				return
			}
			orc.SubmitResult(result)
			w.WriteHeader(http.StatusOK)
		default:
			http.Error(w, `{"error": "Method not allowed"}`, http.StatusMethodNotAllowed)
		}
	})

	// Запуск сервера
	fmt.Println("Orchestrator is running on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
