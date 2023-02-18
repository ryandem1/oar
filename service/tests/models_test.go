package tests

import (
	"testing"
)

var Fake = newFaker()

// TestValidTestsPassValidate will ensure that the Test.Validate() function correctly accepts valid Test objects.
func TestValidTestsPassValidate(t *testing.T) {
	validTests := multiple(15, Fake.test)

	for _, validTest := range validTests {
		err := validTest.Validate()
		if err != nil {
			t.Error(err)
		}
	}
}
