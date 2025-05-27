package utils

import (
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRepeatCheck_CallsFunction(t *testing.T) {
	var count int32

	RepeatCheck(10*time.Millisecond, 50*time.Millisecond, func() {
		atomic.AddInt32(&count, 1)
	})

	assert.GreaterOrEqual(t, count, int32(4))
}
