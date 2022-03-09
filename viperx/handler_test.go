package viperx

import (
	"fmt"
	"testing"
)

func TestNewViper(t *testing.T) {
	v := NewConfig("config")
	fmt.Println(v.Get("ok"))
}
