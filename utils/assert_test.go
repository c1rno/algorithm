package utils

import "testing"

func TestAssertOk(t *testing.T) {
	Assert("true", true)
}

func TestAssertFail(t *testing.T) {
	defer func() {
		r := recover()
		if r.(string) == "Assertion: false" {
			return
		}

		t.Fatalf("Assert fails: '%v'", r)
	}()

	Assert("false", false)
}
