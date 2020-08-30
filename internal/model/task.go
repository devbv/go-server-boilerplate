package model

import "time"

type Task struct {
	Name   string
	Detail string
	Due    *time.Time
}
