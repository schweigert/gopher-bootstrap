package services

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type AnagramSuite struct {
	suite.Suite
}

func (suite *AnagramSuite) TestNew() {
	ser := NewAnagram()
	suite.NotNil(ser)
}

func (suite *AnagramSuite) TestCheck() {
	ser := NewAnagram()

	suite.True(ser.Check("abc", "bca"))
	suite.True(ser.Check("abc", "abc"))
	suite.True(ser.Check("abc", "cba"))
	suite.True(ser.Check("abc", "cab"))
}

func TestAnagramSuite(t *testing.T) {
	suite.Run(t, &AnagramSuite{})
}
