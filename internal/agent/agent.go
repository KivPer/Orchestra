package agent

import (
	"time"
	"yandexGoCalc/internal/models"
)

type Agent struct {
	orchestratorURL string
	computingPower  int
}

func NewAgent(orchestratorURL string, computingPower int) *Agent {
	return &Agent{
		orchestratorURL: orchestratorURL,
		computingPower:  computingPower,
	}
}

func (a *Agent) Start() {
	for i := 0; i < a.computingPower; i++ {
		go a.worker()
	}
}

func (a *Agent) worker() {
	for {
		task := a.fetchTask()
		if task.ID == "" {
			time.Sleep(1 * time.Second)
			continue
		}

		result := a.executeTask(task)
		a.submitResult(result)
	}
}

func (a *Agent) fetchTask() models.Task {
	// Логика получения задачи от оркестратора
	return models.Task{}
}

func (a *Agent) executeTask(task models.Task) models.Result {
	// Логика выполнения задачи
	time.Sleep(time.Duration(task.OperationTime) * time.Millisecond)
	return models.Result{
		ExpressionID: task.ExpressionID,
		Value:        task.Arg1 + task.Arg2, // Пример вычисления
	}
}

func (a *Agent) submitResult(result models.Result) {
	// Логика отправки результата оркестратору
}
