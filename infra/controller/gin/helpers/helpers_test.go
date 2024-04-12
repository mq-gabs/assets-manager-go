package helpers

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGetIdFromParams(t *testing.T) {
	val := "123"
	uint16_val := 123

	p := gin.Param{
		Key:   "id",
		Value: val,
	}

	var ps gin.Params

	ps = append(ps, p)

	c := gin.Context{
		Params: ps,
	}

	id, err := GetIdFromParams(&c)

	if err != nil {
		t.Fatalf("Some error occurred: %v", err)
	}

	if id != uint16(uint16_val) {
		t.Errorf("Values do not match. Values(%q, %q)", id, uint16_val)
	}
}
