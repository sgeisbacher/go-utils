package fileutils

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestAppendSlash(t *testing.T) {
	RegisterTestingT(t)

	testDataTable := []struct {
		path     string
		expected string
	}{
		{"./foo/bar", "./foo/bar/"},
		{"./foo/bar/", "./foo/bar/"},
		{"   ./foo/bar   ", "./foo/bar/"},
	}

	for _, testData := range testDataTable {
		Expect(AppendSlash(testData.path)).To(Equal(testData.expected))
	}
}
