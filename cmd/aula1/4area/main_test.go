package main

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type MainSuite struct {
	suite.Suite
}

func (suite *MainSuite) TestMain() {
	suite.NotPanics(main)
}

func TestMainSuite(t *testing.T) {
	suite.Run(t, &MainSuite{})
}

func BenchmarkMain(b *testing.B) {
	for n := 0; n < b.N; n++ {
		main()
	}
}
