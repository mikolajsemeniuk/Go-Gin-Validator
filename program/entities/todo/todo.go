package todo

import "time"

type Todo struct {
	Id          int64     `json:"id" validate:"min=1,max=10"`
	Title       string    `json:"title" validate:"nonzero"`
	Description string    `json:"description"`
	Email       string    `json:"email" validate:"regexp=^[0-9a-z]+@[0-9a-z]+(\\.[0-9a-z]+)+$"`
	Date        time.Time `json:"time"`
	Completed   bool      `json:"completed"`
}
