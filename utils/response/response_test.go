package response

import "testing"

func TestNew(t *testing.T) {
	msg := "This is a test message"
	r := New(msg)

	if r.Message != msg {
		t.Errorf("Mesages do not match. Messages(%q, %q)", msg, r.Message)
	}
}
