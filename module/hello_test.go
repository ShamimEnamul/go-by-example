package module

import "testing"

func TestHello(t *testing.T) {

	t.Run("basic test", func(t *testing.T) {
		got := hello()
		want := "bye"
		if got != want {
			t.Errorf("Test failed! want %s  and got %s", want, got)
		}
	})
}
