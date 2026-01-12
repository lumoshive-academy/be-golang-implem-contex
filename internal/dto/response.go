package dto

type ResponseUser struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Pagination struct {
	CurrentPage  int `json:"current_page"`
	Limit        int `json:"limit"`
	TotalPages   int `json:"total_pages"`
	TotalRecords int `json:"total_records"`
}
