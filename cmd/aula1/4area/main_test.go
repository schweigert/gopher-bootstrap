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
