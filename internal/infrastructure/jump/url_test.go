package jump

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type URLTests struct {
	suite.Suite
}

func (self *URLTests) SetupTest() {}

func (self *URLTests) TestSetPage() {
	url := URL("/test?page={{.Page}}")
	page := 10

	url = url.SetPage(page)

	self.Contains(url, "/test?page=10")
}

func TestJumpClientTests(t *testing.T) {
	suite.Run(t, new(URLTests))
}
