package core

import (
	"fmt"
	"testing"
)

func TestNewAOIManager(t *testing.T) {
	manager := NewAOIManager(0, 250, 5, 0, 250, 5)
	fmt.Println(manager)
}
