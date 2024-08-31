package model

import "time"

type Restaurant struct {
	ID             string
	Name           string
	Address        string
	NearestStation string
	Genre          string
	URL            string
	PostedAt       time.Time
}
