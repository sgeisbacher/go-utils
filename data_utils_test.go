package utils

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestToIDReplacesSpacesWithDash(t *testing.T) {
	RegisterTestingT(t)

	var tableTestData = []struct {
		str      string
		expected string
	}{
		{"Name of a Product", "name-of-a-product"},
		{" Product's name   ", "products-name"},
		{" Äpfel sind öfters grün ! ", "aepfel-sind-oefters-gruen"},
	}
	for _, testData := range tableTestData {
		Expect(ToID(testData.str)).To(Equal(testData.expected))
	}
}
