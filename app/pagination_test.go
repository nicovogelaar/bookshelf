package app

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPagination(t *testing.T) {
	p := newPagination(5, 10)

	assert.Equal(t, 10, p.limit())
	assert.Equal(t, 40, p.offset())
}
