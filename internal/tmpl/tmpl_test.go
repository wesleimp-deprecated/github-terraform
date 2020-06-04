package tmpl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInvalidTemplate(t *testing.T) {
	assert := assert.New(t)
	_, err := New().Apply("{{{.Foo}")
	assert.EqualError(err, "template: tmpl:1: unexpected \"{\" in command")
}

func TestWithFields(t *testing.T) {
	assert := assert.New(t)

	out, err := New().WithFields(Fields{"Hey": "Heeey Test"}).Apply("{{ .Hey }}")
	assert.NoError(err)
	assert.Equal("Heeey Test", out)
}
