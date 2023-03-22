package cache

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBigCacheOptions(t *testing.T) {
	InitBigCacheConfig()
	expect := "This is a Test string"
	BigCache.Set("test", expect)
	result, _ := BigCache.Get("test")
	assert.Equal(t, result.(string), expect)
}
