package model

import "time"

type User struct {
	ID    int
	Count int
	Date  time.Time
}
