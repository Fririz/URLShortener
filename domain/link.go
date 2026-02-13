package domain

type Link struct {
	ID        int    `json:"id"`
	URL       string `json:"url"`
	Slug      string `json:"slug"`
	CreatedAt string `json:"created_at"`
	Visits    int    `json:"visits"`
}
