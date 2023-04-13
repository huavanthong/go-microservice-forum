package common

import "time"

type EntityBase struct {
	ID               int        `json:"id"`
	CreatedBy        string     `json:"createdBy"`
	CreatedDate      time.Time  `json:"createdDate"`
	LastModifiedBy   string     `json:"lastModifiedBy"`
	LastModifiedDate *time.Time `json:"lastModifiedDate,omitempty"`
}
