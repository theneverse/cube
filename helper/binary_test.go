package helper

import (
	"fmt"
	"testing"
)

func TestIntToBytes(t *testing.T) {
	fmt.Println(Uint32ToBytes(257))
	fmt.Println(Uint32ToBytes(1))
	fmt.Println(Uint16ToBytes(1004))
	fmt.Println(Uint16ToBytes(65530))
}
