package test

import (
	modelTest "github.com/wejick/alive/internal/model/test"
	repoTest "github.com/wejick/alive/internal/repo/test"
)

type Test struct {
	testRepo repoTest.Itest
}

func New(testRepo repoTest.Itest) *Test {
	return &Test{
		testRepo: testRepo,
	}
}

//GetTest get test data by id or by pagination.
// if IDs are provided, rows and page will be ignored
// page started from 1
func (T *Test) GetTest(IDs []string, rows, page int64) (testlist []modelTest.Test, err error) {
	offset := rows * page
	testlist, err = T.testRepo.GetTest(IDs, int(rows), int(offset))

	return
}

func (T *Test) GetTotalTest() (total int64, err error) {
	return T.testRepo.GetTotalTest()
}

func (T *Test) AddTest(test modelTest.Test) (err error) {
	return T.testRepo.AddTest(test)
}
