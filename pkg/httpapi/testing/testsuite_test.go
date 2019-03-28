package testing

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

func Test(t *testing.T) {
	suite.Run(t, &TestSuite{})
}
