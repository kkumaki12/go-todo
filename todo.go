package main

import (
	"time"
)

type Todo struct {
	ID        int
	Content   string
	CreatedAt time.Time
}
