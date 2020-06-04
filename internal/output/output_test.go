package output

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSaveNewFileWithDest(t *testing.T) {
	assert := assert.New(t)

	err := Save("testdata", "file", "Test")
	assert.NoError(err)

	bb, err := ioutil.ReadFile("./testdata/file.tf")
	assert.NoError(err)
	assert.Equal(string(bb), "Test")
}

func TestReplaceFile(t *testing.T) {
	assert := assert.New(t)

	err := Save("testdata", "file.replace", "Test")
	assert.NoError(err)

	bb, err := ioutil.ReadFile("./testdata/file.replace.tf")
	assert.NoError(err)
	assert.Equal(string(bb), "Test")

	err = Save("testdata", "file.replace", "Test Replace")
	assert.NoError(err)

	bb, err = ioutil.ReadFile("./testdata/file.replace.tf")
	assert.NoError(err)
	assert.Equal(string(bb), "Test Replace")
}
