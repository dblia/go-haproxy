/*

  Helper functions for testing the HAProxy client.

*/
package haproxy

import "testing"

func assertEqual(a, b string, t *testing.T) {
	if a != b {
		t.Fatalf("Expected %#v and %#v to be equal.", a, b)
	}
}

func assertNotEqual(a, b string, t *testing.T) {
	if a == b {
		t.Fatalf("Expected %#v and %#v to not be equal.", a, b)
	}
}

func assertNil(e error, t *testing.T) {
	if e != nil {
		t.Fatalf("Expected %#v to be nil.", e)
	}
}
