package orchestrator

import (
	"sync"
	"yandexGoCalc/internal/models"
)

type Orchestrator struct {
	expressions map[string]*models.Expression
	tasks       chan models.Task
	results     chan models.Result
	mu          sync.Mutex
}

func NewOrchestrator() *Orchestrator {
	return &Orchestrator{
		expressions: make(map[string]*models.Expression),
		tasks:       make(chan models.Task, 100),
		results:     make(chan models.Result, 100),
	}
}

func (o *Orchestrator) AddExpression(expr string) string {
	o.mu.Lock()
	defer o.mu.Unlock()

	id := generateID()
	o.expressions[id] = &models.Expression{
		ID:     id,
		Status: "pending",
		Expr:   expr,
	}

	// Разбиваем выражение на задачи и отправляем в канал задач
	tasks := parseExpression(expr)
	for _, task := range tasks {
		o.tasks <- task
	}

	return id
}

func (o *Orchestrator) GetTask() models.Task {
	return <-o.tasks
}

func (o *Orchestrator) SubmitResult(result models.Result) {
	o.results <- result
}

func (o *Orchestrator) ProcessResults() {
	for result := range o.results {
		o.mu.Lock()
		expr := o.expressions[result.ExpressionID]
		expr.Status = "completed"
		expr.Result = result.Value
		o.mu.Unlock()
	}
}

func parseExpression(expr string) []models.Task {
	// Логика разбора выражения на задачи
	return []models.Task{}
}

func generateID() string {
	// Логика генерации уникального ID
	return "unique-id"
}
