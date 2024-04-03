package exception

import "testing"

func TestNew(t *testing.T) {
	exception := New("Not Found", 404)

	if exception.statusCode != 404 || exception == nil {
		t.Errorf("Cannot create exception")
	}
}
