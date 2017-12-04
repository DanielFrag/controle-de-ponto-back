package dto

import "time"

//DateRegister struct data transfer object
type DateRegister struct {
	Timestamp   time.Time `json:"timeStamp"`
	Description string    `json:"description"`
}
