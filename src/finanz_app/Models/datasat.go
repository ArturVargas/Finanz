package models

import (
	"time"
)

// Obligaciones ...
type Obligaciones struct {
	Id          uint      `json:"id"`
	Description string    `json:"description"`
	Idregimen   uint      `json:"regimen"`
	Created_at  time.Time `json:"created_at"`
	Updated_at  time.Time `json:"updated_at"`
}

// Deducciones con su estructura
type Deducciones struct {
	Id          uint      `json:"id"`
	Description string    `json:"description"`
	Idregimen   uint      `json:"regimen"`
	Created_at  time.Time `json:"created_at"`
	Updated_at  time.Time `json:"updated_at"`
	Type        string    `json:"type"`
}
