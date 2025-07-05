package models

import (
	"time"
)

type Status string

const (
	StatusDone    Status = "done"
	StatusFinish  Status = "finish"
	StatusRunning Status = "running"
)

type Schedule struct {
	PetshopId   int        `json:"petshop_id"`
	Id          int        `json:"id"`
	Pet         Pet        `json:"pet"`
	Sku         string     `json:"sku"`
	Services    []Services `json:"services"`
	Observation string     `json:"observation"`
	Status      Status     `json:"status"`
	Total       float64    `json:"total"`
	DateTime    time.Time  `json:"date_time"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}
