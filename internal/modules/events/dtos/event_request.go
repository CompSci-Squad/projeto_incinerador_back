package dtos

type Place struct {
	Street     string `json:"street"`
	PostalCode string `json:"postal_code"`
	Number     string `json:"number"`
}

type EventsIndexRequest struct {
	Limit int64     `json:"limit"`
	Page  int64     `json:"page"`
	Name  string    `json:"name"`
	Ids   *[]string `json:"ids"`
}

type EventSingleRequest struct {
	Id string `json:"id"`
}

type EventUpdateRequest struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	DateTime string `json:"date"`
	Place    *Place `json:"place"`
}

type EventStoreRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	DateTime    string `json:"date"`
	Place       *Place `json:"place"`
}
