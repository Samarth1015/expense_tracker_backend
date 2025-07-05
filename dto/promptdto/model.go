package promptdto

type ExpenseData struct {
	Amount      float64 `json:"amount"`
	Category    string  `json:"category"`
	Date        string  `json:"date"`
	Description string  `json:"description"`
}
