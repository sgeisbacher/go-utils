package sliceutils

import "testing"
import . "github.com/onsi/gomega"

func TestContains(t *testing.T) {
	RegisterTestingT(t)

	tableTestData := []struct {
		slice    []string
		search   string
		expected bool
	}{
		{[]string{"apple", "banana", "pineapple"}, "apple", true},
		{[]string{"apple", "banana", "pineapple"}, "banana", true},
		{[]string{"apple", "banana", "pineapple"}, "strawberry", false},
		{nil, "apple", false},
	}

	for _, testData := range tableTestData {
		Expect(Contains(testData.slice, testData.search)).To(Equal(testData.expected))
	}
}
