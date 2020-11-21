package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
)

type MainSuite struct {
	suite.Suite
}

func (suite *MainSuite) TestMain() {
	defer os.Setenv("PORT", os.Getenv("PORT"))
	os.Setenv("PORT", "error")

	suite.NotPanics(main)
}

func TestMainSuite(t *testing.T) {
	suite.Run(t, &MainSuite{})
}
