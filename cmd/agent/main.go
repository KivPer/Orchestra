package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"yandexGoCalc/internal/agent"
)

func main() {
	// Получаем переменные среды
	orchestratorURL := os.Getenv("ORCHESTRATOR_URL")
	if orchestratorURL == "" {
		orchestratorURL = "http://localhost:8080" // Значение по умолчанию
	}

	computingPowerStr := os.Getenv("COMPUTING_POWER")
	if computingPowerStr == "" {
		computingPowerStr = "4" // Значение по умолчанию
	}
	computingPower, err := strconv.Atoi(computingPowerStr)
	if err != nil {
		log.Fatalf("Invalid COMPUTING_POWER value: %v", err)
	}

	// Создаем и запускаем агента
	ag := agent.NewAgent(orchestratorURL, computingPower)
	fmt.Println("Agent is running with computing power:", computingPower)
	ag.Start()

	// Бесконечный цикл для поддержания работы агента
	select {}
}
