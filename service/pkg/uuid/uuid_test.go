package uuid

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUUID_New(t *testing.T) {
	uuid := New()
	t.Logf("TestUUID_New:  New log.uuid:=%v\n", uuid.Format())
}

func TestUUID_NewUUID1(t *testing.T) {
	uuid, err := NewUUID1()
	assert.NoError(t, err)
	t.Logf("TestUUID_NewUUID1:  NewUUID1 log.uuid:=%v\n", uuid.Format())
}

func TestUUID_NewOrderedUUID(t *testing.T) {
	uuid, err := NewOrderedUUID()
	assert.NoError(t, err)
	t.Logf("TestUUID_NewOrderedUUID:  NewOrderedUUID log.uuid:=%v\n", uuid.Format())
}

func TestUUID_Parse(t *testing.T) {
	uuid := New()
	uuid1, err := NewUUID1()
	orderedUUID, err := NewOrderedUUID()
	assert.NoError(t, err)

	var parseUUID UUID
	parseUUID, err = Parse(uuid.String())
	assert.NoError(t, err)
	t.Logf("TestUUID_Parse:  Parse-uuid log.uuid:=%v\n", parseUUID.Format())

	parseUUID, err = Parse(uuid1.String())
	assert.NoError(t, err)
	t.Logf("TestUUID_Parse:  Parse-uuid-1 log.uuid:=%v\n", parseUUID.Format())

	parseUUID, err = Parse(orderedUUID.String())
	assert.NoError(t, err)
	t.Logf("TestUUID_Parse:  Parse-uuid-ordered log.uuid:=%v\n", parseUUID.Format())
}
