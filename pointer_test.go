package pointer

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

// ------------------------------------------------- --------------------------------------------------------------------

func TestFalsePointer(t *testing.T) {
	pointer := FalsePointer()
	assert.NotNil(t, pointer)
	assert.False(t, *pointer)
}

func ExampleFalsePointer() {
	pointer := FalsePointer()
	fmt.Printf("%T %v", pointer, *pointer)
	// Output:
	// *bool false
}

// ------------------------------------------------- --------------------------------------------------------------------

func TestFromPointer(t *testing.T) {
	v1 := true
	v2 := &v1
	pointer := FromPointer(v2)
	assert.Equal(t, true, pointer)
}

// ------------------------------------------------- --------------------------------------------------------------------

func TestFromPointerOrDefault(t *testing.T) {
	v1 := true
	v2 := &v1
	pointer := FromPointer(v2)
	assert.Equal(t, true, pointer)
}

// ------------------------------------------------- --------------------------------------------------------------------

func TestToPointer(t *testing.T) {
	v := 1
	pointer := ToPointer(v)
	t.Log(pointer)
}

// ------------------------------------------------- --------------------------------------------------------------------

func TestToPointerOrNil(t *testing.T) {
	orNil := ToPointerOrNil(0)
	t.Log(orNil)
}

// ------------------------------------------------- --------------------------------------------------------------------

func TestTruePointer(t *testing.T) {
	pointer := TruePointer()
	assert.NotNil(t, pointer)
	assert.True(t, *pointer)
}

func ExampleTruePointer() {
	pointer := TruePointer()
	fmt.Printf("%T %v", pointer, *pointer)
	// Output:
	// *bool true
}

// ------------------------------------------------- --------------------------------------------------------------------
