package trace

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrace(t *testing.T) {
	trace := GetTrace()
	assert.NotNil(t, trace)
}
