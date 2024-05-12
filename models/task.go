package models

import (
	"gorm.io/gorm"
)

const (
	STATUS_DONE    = "done"
	STATUS_IN_WORK = "in_work"
	STATUS_CREATED = "created"
)

type Task struct {
	gorm.Model
	Name   string
	Status string
}
