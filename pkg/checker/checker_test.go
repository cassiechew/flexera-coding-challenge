package checker

import (
	"testing"
    "github.com/stretchr/testify/assert"

	"github.com/ryanchew3/flexera-coding-challenge/internal/model"
)

func TestGetKeys(t *testing.T) {
	toTest := map[int64][]*model.Computer{}
	add := &model.Computer{
		Id: 1,
		AppId: 2,
		Type: "desktop",
	}
	toTest[1] = []*model.Computer{add}

	exp := []int64{1}
	v := getKeys(toTest)
	assert.Equal(t, exp, v)
}


func TestProcessChecks(t *testing.T) {
	toTest := map[int64][]*model.Computer{}
	add := &model.Computer{
		Id: 1,
		AppId: 2,
		Type: "desktop",
	}
	toTest[1] = []*model.Computer{add}

	exp := 1
	v := ProcessChecks(toTest)
	assert.Equal(t, exp, v)
}