package deck

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 20 {
		t.Errorf("Wrong deck card quantity, should be 20 but got %d\n", len(d))
	}
}
