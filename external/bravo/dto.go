package bravo

type (
	Query struct {
		Keyword string `json:"keyword"`
		Value   string `json:"value"`
	}

	Request struct {
		Query Query `json:"query"`
		Limit int   `json:"limit"`
		Page  int   `json:"page"`
	}

	Store struct {
		Name    string `json:"name"`
		Address string `json:"address"`
		Phone   string `json:"phone"`
	}

	Response struct {
		Message string  `json:"message"`
		Status  int     `json:"status"`
		Stores  []Store `json:"stores"`
	}
)
