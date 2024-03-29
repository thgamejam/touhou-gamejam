package cache

import (
	"context"
	"github.com/stretchr/testify/assert"
	"service/pkg/conf"
	"testing"
)

func TestNewCache(t *testing.T) {
	c := conf.NewPkgTestConfService()
	cache, err := NewCache(c.Service)
	assert.NoError(t, err)
	assert.NotNil(t, cache)
}

func TestCache_SetANDGet(t *testing.T) {
	c := conf.NewPkgTestConfService()
	cache, _ := NewCache(c.Service)
	ctx := context.Background()

	var g1 map[string]string

	ok, err := cache.Get(ctx, "a", g1)
	assert.NoError(t, err)
	assert.False(t, ok)

	m := map[string]string{
		"a": "a",
	}

	err = cache.Set(ctx, "aaa", m, 0)
	assert.NoError(t, err)
	var r map[string]string
	ok, err = cache.Get(ctx, "aaa", &r)
	assert.NoError(t, err)
	assert.True(t, ok)
	assert.Equal(t, m, r)
}
