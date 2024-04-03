package exception

import "testing"

func TestNew(t *testing.T) {
	exception := New("Not Found", 404)

	if exception.StatusCode != 404 || exception == nil {
		t.Errorf("Cannot create exception")
	}
}
