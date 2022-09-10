package test

import (
	model "github.com/wejick/alive/internal/model"
)

type Itest interface {
	GetTest(IDs []string, agent string, rows, offset int) ([]model.Test, error)
	GetTotalTest() (int64, error)
	AddTest(model.Test) error
}
