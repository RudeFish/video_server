package utils

import (
	"testing"
	"fmt"
)

func TestNewUUID(t *testing.T) {
	uuid, err := NewUUID()
	if err != nil{
		t.Error(err)
	}
	fmt.Println(uuid)
}
