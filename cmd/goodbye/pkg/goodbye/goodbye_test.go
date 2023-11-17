package goodbye_test

import (
	"testing"

	"github.com/cttttt/multi-go-go/cmd/goodbye/pkg/goodbye"
	"github.com/stretchr/testify/assert"
)

func TestGoodbye(t *testing.T) {
	assert.Equal(t, "Goodbye", goodbye.Goodbye())
}
