package task

import (
	model "github.com/multiverse-os/maglev/model"
)

type Record struct {
	model.Record

	Name        string
	Description string
}
