package request

type ExpenseReq struct {
	Amount      uint   `json:"amount"`
	Category    string `json:"category"`
	Date        string `json:"date,omitempty"`
	Description string `json:"description"`
}
