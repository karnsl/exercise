package register

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type cliMock struct {
	title string
}

func (c cliMock) httpGet(url string) ([]byte, error) {
	return []byte(c.title), nil
}

func Test_userExists(t *testing.T) {

	response := `{"title":"testUser"}`
	c := cliMock{
		title: response,
	}

	assert.True(t, userExists(c, "testUser", "url"))
}
