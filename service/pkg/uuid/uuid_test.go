package uuid

import "testing"

func TestUUID_New(t *testing.T) {
	uuid := New()
	t.Logf("TestUUID_New:  New log.uuid:=%v\n", uuid.Format())
}

func TestUUID_NewUUID1(t *testing.T) {
	uuid, err := NewUUID1()
	if err != nil {
		t.Fatalf("TestUUID_NewUUID1:  NewUUID1 failed!err:=%v", err)
	}
	t.Logf("TestUUID_NewUUID1:  NewUUID1 log.uuid:=%v\n", uuid.Format())
}

func TestUUID_NewOrderedUUID(t *testing.T) {
	uuid, err := NewOrderedUUID()
	if err != nil {
		t.Fatalf("TestUUID_NewOrderedUUID:  NewOrderedUUID failed!err:=%v", err)
	}
	t.Logf("TestUUID_NewOrderedUUID:  NewOrderedUUID log.uuid:=%v\n", uuid.Format())
}

func TestUUID_Parse(t *testing.T) {
	uuid := New()
	uuid1, err := NewUUID1()
	orderedUUID, err := NewOrderedUUID()

	var parseUUID UUID
	parseUUID, err = Parse(uuid.String())
	if err != nil {
		t.Fatalf("TestUUID_Parse:  Parse-uuid failed!err:=%v uuid:=%v", err, uuid.String())
	}
	t.Logf("TestUUID_Parse:  Parse-uuid log.uuid:=%v\n", parseUUID.Format())

	parseUUID, err = Parse(uuid1.String())
	if err != nil {
		t.Fatalf("TestUUID_Parse:  Parse-uuid-1 failed!err:=%v uuid:=%v", err, uuid1.String())
	}
	t.Logf("TestUUID_Parse:  Parse-uuid-1 log.uuid:=%v\n", parseUUID.Format())

	parseUUID, err = Parse(orderedUUID.String())
	if err != nil {
		t.Fatalf("TestUUID_Parse:  Parse-uuid-ordered failed!err:=%v uuid:=%v", err, orderedUUID.String())
	}
	t.Logf("TestUUID_Parse:  Parse-uuid-ordered log.uuid:=%v\n", parseUUID.Format())
}
