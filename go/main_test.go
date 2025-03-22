package sample

import "testing"

func TestMain(t *testing.T) {
	t.Run("main test case", func(t *testing.T) {
		got := "random string"
		if got != "random string" {
			t.Error("This dummy test should never fail.")
		}
	})
}