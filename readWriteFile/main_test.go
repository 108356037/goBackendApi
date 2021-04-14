package main

import "testing"

func BasicTest(t *testing.T) error {
	if err := testWrite("dcc.txt", "moomoo dcc"); err != nil {
		t.Errorf("the write operation failed, err: %v\n", err)
	}
	return nil
}
