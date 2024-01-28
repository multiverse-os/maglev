package task

import (
	model "github.com/multiverse-os/webframe/model"
)

type Record struct {
	model.Record

	Name        string
	Description string
}
