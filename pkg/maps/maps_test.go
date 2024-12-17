package maps

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStoreAndLoad(t *testing.T) {
	m := Map[string, int]{}

	m.Store("key1", 10)
	m.Store("key2", 20)

	value, ok := m.Load("key1")
	assert.True(t, ok)
	assert.Equal(t, 10, value)

	value, ok = m.Load("key2")
	assert.True(t, ok)
	assert.Equal(t, 20, value)

	value, ok = m.Load("nonexistent")
	assert.False(t, ok)
	assert.Equal(t, 0, value) // default zero value of int
}

func TestDelete(t *testing.T) {
	m := Map[string, int]{}

	m.Store("key1", 10)
	m.Store("key2", 20)

	m.Delete("key1")

	_, ok := m.Load("key1")
	assert.False(t, ok)

	value, ok := m.Load("key2")
	assert.True(t, ok)
	assert.Equal(t, 20, value)
}

func TestRange(t *testing.T) {
	m := Map[string, int]{}

	m.Store("key1", 10)
	m.Store("key2", 20)
	m.Store("key3", 30)

	sum := 0
	m.Range(func(key string, value int) bool {
		sum += value
		return true
	})

	assert.Equal(t, 60, sum)
}

func TestClear(t *testing.T) {
	m := Map[string, int]{}

	m.Store("key1", 10)
	m.Store("key2", 20)
	m.Store("key3", 30)

	m.Clear()

	_, ok := m.Load("key1")
	assert.False(t, ok)

	_, ok = m.Load("key2")
	assert.False(t, ok)

	_, ok = m.Load("key3")
	assert.False(t, ok)
}

func TestMarshalJSON(t *testing.T) {
	m := Map[string, int]{}

	m.Store("key1", 10)
	m.Store("key2", 20)

	jsonData, err := json.Marshal(&m)
	assert.NoError(t, err)

	expectedJSON := `{"key1":10,"key2":20}`
	assert.JSONEq(t, expectedJSON, string(jsonData))
}

func TestUnmarshalJSON(t *testing.T) {
	m := Map[string, int]{}

	jsonData := `{"key1":10,"key2":20}`
	err := json.Unmarshal([]byte(jsonData), &m)
	assert.NoError(t, err)

	value, ok := m.Load("key1")
	assert.True(t, ok)
	assert.Equal(t, 10, value)

	value, ok = m.Load("key2")
	assert.True(t, ok)
	assert.Equal(t, 20, value)
}

func TestMarshalAndUnmarshalJSON(t *testing.T) {
	m := Map[string, int]{}

	m.Store("key1", 10)
	m.Store("key2", 20)

	// Marshal to JSON
	jsonData, err := json.Marshal(&m)
	assert.NoError(t, err)

	// Unmarshal back to map
	var newMap Map[string, int]
	err = json.Unmarshal(jsonData, &newMap)
	assert.NoError(t, err)

	// Validate contents of unmarshalled map
	value, ok := newMap.Load("key1")
	assert.True(t, ok)
	assert.Equal(t, 10, value)

	value, ok = newMap.Load("key2")
	assert.True(t, ok)
	assert.Equal(t, 20, value)
}
