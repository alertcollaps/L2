package main

import (
	"fmt"
	"testing"
)

func Test_unBox(t *testing.T) {
	data := []struct {
		A      string
		Result string
		err    error
	}{
		{
			A:      "Hello_Мир",
			Result: "Hello_Мир",
			err:    nil,
		},
		{
			A:      "Hh2",
			Result: "Hhh",
			err:    nil,
		},
		{
			A:      "Y\\23",
			Result: "Y222",
			err:    nil,
		},
		{
			A:      "\\\\\\",
			Result: "",
			err:    fmt.Errorf("No enouth symbols after \\"),
		},
		{
			A:      "\\\\",
			Result: "\\",
			err:    nil,
		},
		{
			A:      "qwe\\4\\5",
			Result: "qwe45",
			err:    nil,
		},
		{
			A:      "qwe\\45",
			Result: "qwe44444",
			err:    nil,
		},
		{
			A:      "qwe\\\\5",
			Result: "qwe\\\\\\\\\\",
			err:    nil,
		},
	}
	for _, testCase := range data {
		res, err := unBox(testCase.A)
		if res != testCase.Result {
			t.Errorf("Expected %v, got %v", testCase.Result, res)
		}
		if err != nil && err.Error() != testCase.err.Error() {
			t.Errorf("Expected %v, got %v", testCase.err, err)
		}
	}
}
