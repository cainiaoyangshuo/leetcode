package _68validIPAddress

import (
	"fmt"
	"testing"
)

func TestValidIPAddress(t *testing.T) {
	ip := "1.0.1."
	res := ValidIPAddress(ip)
	fmt.Printf("%s\n", res)
}
