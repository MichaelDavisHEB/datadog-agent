package packets

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func countPoolSize(p *PoolManager) int {

	i := 0
	p.refs.Range(func(key, value interface{}) bool {
		i++

		return true
	})

	return i
}

func TestPoolManager(t *testing.T) {

	pool := NewPool(1024)
	manager := NewPoolManager(pool)

	// passthru mode by default
	assert.True(t, manager.IsPassthru())

	packet := manager.Get()
	manager.Put(packet)
	assert.Equal(t, 0, countPoolSize(manager))

	// passthru mode disabled
	manager.SetPassthru(false)
	assert.False(t, manager.IsPassthru())
	packet = manager.Get()
	manager.Put(packet)
	assert.Equal(t, 1, countPoolSize(manager))
	// second put should remove it from reference tracker and return buffer to pool
	manager.Put(packet)
	assert.Equal(t, 0, countPoolSize(manager))

	for i := 0; i < 10; i++ {
		packet = manager.Get()
		manager.Put(packet)
	}
	assert.Equal(t, 10, countPoolSize(manager))
	manager.Flush()
	assert.Equal(t, 0, countPoolSize(manager))

}
