package say_hello

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

/*
tidak boleh ada return value pada unit test
*/

func TestSayHello(t *testing.T) {
	hello := SayHello()
	if hello != "Hello World" {
		t.Error("Result must be 'Hello World'")
	}
}

func TestSayHelloWithAssert(t *testing.T) {
	hello := SayHello()
	assert.Equal(t, "Hello World", hello, "Result must be 'Hello World'")
}

func TestSayHelloWithRequire(t *testing.T) {
	hello := SayHello()
	require.Equal(t, "Hello World", hello, "Result must be 'Hello World'")
}

func TestIgnoreTestForLinux(t *testing.T) {
	if runtime.GOOS == "linux" {
		t.Skip("Ignore test for linux Os")
	}
	hello := SayHello()
	assert.Equal(t, "Hello World", hello, "Result must be 'Hello World'")
}

/* Sub Test */
func TestSubTest(t *testing.T){
	hello := SayHello()
	t.Run("SubTest1", func(t *testing.T){
        assert.Equal(t, "Hello World", hello, "Result must be 'Hello World'")
    })

	t.Run("SubTest2", func(t *testing.T){
        assert.Equal(t, "Hello World", hello, "Result must be 'Hello World'")
    })
}


func TestTableTest(t *testing.T) {
	
	tests := []struct {
		name  string
        request string
        expected string
	}{
		{
			name:  "Test Return Said",
            request: "Said",
            expected: "Hi Said",
		},
		{
			name:  "Test Return Al",
            request: "Al",
            expected: "Hi Al",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			say_hello_to := SayHelloTo(test.request)
			assert.Equal(t, test.expected, say_hello_to, "Result not Equals")
		})
	}

}


