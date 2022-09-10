package test

import (
	modelTest "github.com/wejick/alive/internal/model/test"
)

type Itest interface {
	GetTest(IDs []string, agent string, rows, offset int) ([]modelTest.Test, error)
	GetTotalTest() (int64, error)
	AddTest(modelTest.Test) error
}
