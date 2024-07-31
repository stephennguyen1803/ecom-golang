package basic

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAddOne(t *testing.T) {
	// var input, output = 1, 3

	// actual := AddOne(input)
	// if actual != output {
	// 	t.Errorf("AddOne(%d), ouput = %d, actual = %d", input, output, actual)
	// }

	assert.Equal(t, 3, AddOne(1), "AddOne(1) should return 2")
	assert.NotEqual(t, 1, AddOne(1))
}

func TestRequire(t *testing.T) {
	require.Equal(t, 3, AddOne(1))
	fmt.Println("Not Executed")
}

func TestAssert(t *testing.T) {
	assert.Equal(t, 1, 2)
	fmt.Println("Executed")
}
