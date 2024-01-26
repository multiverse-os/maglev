package task

import (
	model "github.com/multiverse-os/webkit/model"
)

type Record struct {
	model.Record

	Name        string
	Description string
}
