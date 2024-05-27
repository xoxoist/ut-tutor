package alpha

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

	Product struct {
		Name     string `json:"name"`
		Category string `json:"category"`
		Price    string `json:"price"`
	}

	Response struct {
		Message  string    `json:"message"`
		Status   int       `json:"status"`
		Products []Product `json:"products"`
	}
)
