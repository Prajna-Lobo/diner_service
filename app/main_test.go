package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadFromFile(t *testing.T) {
	cases := map[string]struct {
		path        string
		linebreak   int
		expectedErr error
	}{
		"Should successfully open the file from the given valid path and read first 3 lines": {
			path:        "../testdata/test_log_valid.csv",
			linebreak:   2,
			expectedErr: nil,
		},
		"Should successfully open the file from the given valid path and throw error when duplicate entry found": {
			path:        "../testdata/test_log_duplicate.csv",
			linebreak:   2,
			expectedErr: errors.New("duplicate entry"),
		},
	}

	for caseTitle, tc := range cases {
		t.Run(caseTitle, func(t *testing.T) {
			err := ReadFromFile(tc.path, tc.linebreak)

			assert.Equal(t, tc.expectedErr, err)
		})
	}
}
