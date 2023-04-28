package dtos

type Task struct {
	ID          int64  `json:id`
	Description string `json:description`
	Board       int64  `json:board`
}
