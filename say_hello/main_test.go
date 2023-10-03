package say_hello

import (
	"fmt"
	"testing"
)

func TestMain(m *testing.M) {
	// Before starting
	fmt.Println("Before starting testings...")

	m.Run() // running all test

	// After starting
	fmt.Println("After starting testings...")
}