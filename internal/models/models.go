package models

type Expression struct {
	ID     string
	Status string
	Expr   string
	Result float64
}

type Task struct {
	ID            string
	ExpressionID  string
	Arg1          float64
	Arg2          float64
	Operation     string
	OperationTime int
}

type Result struct {
	ExpressionID string
	Value        float64
}
