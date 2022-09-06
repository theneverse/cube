package uuid

import (
	"fmt"
	"testing"
)

func TestGenUUID(t *testing.T) {
	uuid := GenUUID()
	fmt.Println(uuid)
	fmt.Println(uuid.String())
	fmt.Println(uuid.Encode())
	fmt.Println(uuid.ConvertToBigInt())
}
