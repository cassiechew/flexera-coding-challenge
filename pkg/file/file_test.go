package file

import (
	"testing"
    "github.com/stretchr/testify/assert"
)

type testcase struct {
	name string
	test func(*testing.T)
}

func TestReadFile(t *testing.T) {
	for _, c := range []testcase{
		{"Fails to run when no file path is given", noFileInput},
		{"Fails to run when a file path that does not exist is given", notExistingFileInput},
		{"Fails to read when file has invalid data in computerID field", invalidReadComputerId},
		{"Fails to read when file has invalid data in appID field", invalidReadAppId},
		{"Successfully reads valid file", validRead},
	} {
		t.Run(c.name, c.test)
	}
}

func noFileInput(t *testing.T) {
	o, e := ReadFile("")
	assert.Nil(t, o)
	assert.Error(t, e)
}

func notExistingFileInput(t *testing.T) {
	o, e := ReadFile("notafile")
	assert.Nil(t, o)
	assert.Error(t, e)
}

func invalidReadComputerId(t *testing.T) {
	file := "./testdata/valid_data_id.csv"
	o, e := ReadFile(file)
	assert.Nil(t, o)
	assert.Error(t, e)
}

func invalidReadAppId(t *testing.T) {
	file := "./testdata/invalid_data_aid.csv"
	o, e := ReadFile(file)
	assert.Nil(t, o)
	assert.Error(t, e)
}

func validRead(t *testing.T) {
	file := "./testdata/valid_data.csv"
	o, e := ReadFile(file)
	assert.Equal(t, 4, len(o))
	assert.NoError(t, e)
}