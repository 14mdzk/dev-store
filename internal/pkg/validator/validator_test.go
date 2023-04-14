package validator

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidator_Check(t *testing.T) {
	type DummyRequest struct {
		Name        string
		Description string
	}

	type TestCase struct {
		Name        string
		InvalidData bool
		ReqBody     string
	}

	cases := []TestCase{
		{
			Name:        "When name not presence",
			InvalidData: true,
			ReqBody:     "{'description': 'some descriptions'}",
		},
		{
			Name:        "When description not presence",
			InvalidData: true,
			ReqBody:     "{'name': 'some name'}",
		},
	}

	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			var req DummyRequest
			_ = json.Unmarshal([]byte(tc.ReqBody), &req)
			isInvalid := Check(req)
			assert.Equal(t, tc.InvalidData, isInvalid)
		})
	}
}
