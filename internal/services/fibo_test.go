package services

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type FiboSuite struct {
	suite.Suite
}

func (suite *FiboSuite) TestNew() {
	suite.NotNil(NewFibo())
}

func (suite *FiboSuite) TestFibo() {
	suite.Equal(uint(0), NewFibo().Fibo(0))
	suite.Equal(uint(1), NewFibo().Fibo(1))
	suite.Equal(uint(1), NewFibo().Fibo(2))
}

func TestFiboSuite(t *testing.T) {
	suite.Run(t, &FiboSuite{})
}
