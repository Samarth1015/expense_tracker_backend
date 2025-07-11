package request

import "time"

type ExpenseReq struct {
	Amount      uint   `json:"amount"`
	Category    string `json:"category"`
	Date        string `json:"date,omitempty"`
	Description string `json:"description"`
}

type UpdateExpenseReq struct {
	ExpenseId   int       `json:"expense_id"`
	Amount      uint      `json:"amount"`
	Category    string    `json:"category"`
	Date        time.Time `json:"date,omitempty"`
	Description string    `json:"description"`
}
