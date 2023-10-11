package dtos

type EventsIndexRequest struct {
	Limit int64     `json:"limit"`
	Page  int64     `json:"page"`
	Name  string    `json:"name"`
	Ids   *[]string `json:"ids"`
}
