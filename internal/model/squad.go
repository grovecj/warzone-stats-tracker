package model

import "time"

type Squad struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Members   []Player  `json:"members,omitempty"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
