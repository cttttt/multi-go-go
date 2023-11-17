package hello_test

import (
	"testing"

	"github.com/cttttt/multi-go-go/pkg/hello"
	"github.com/stretchr/testify/assert"
)

func TestHello(t *testing.T) {
	assert.Equal(t, "Hello", hello.Hello())
}
