package xb360ctrl
import (
  "testing"
)

func TestMarshall(t *testing.T) {
  var e Xbc_event
  e.Time = 1
  e.Value = 2
  e.EventType = 3
  e.Number = 4

  data, _ := e.MarshalBinary()

  var e2 Xbc_event
  e2.UnMarshalBinary(data)
  if e2.Time != e.Time {
    t.Error("Expected ", e.Time ,", got ", e2.Time)
  }
  if e2.Value != e.Value {
    t.Error("Expected ", e.Value, ", got ", e2.Value)
  }
  if e2.EventType !=  e.EventType {
    t.Error("Expected ", e.EventType, ", got ", e2.EventType)
  }
  if e2.Number != e.Number {
    t.Error("Expected ", e.Number, ", got ", e2.Number)
  }

  e.Time = 0xffffffff
  e.Value = 0xff
  e.EventType = 0xff
  e.Number = 0xff

  data, _ = e.MarshalBinary()
  e2.UnMarshalBinary(data)
  if e2.Time != e.Time {
    t.Error("Expected ", e.Time ,", got ", e2.Time)
  }
  if e2.Value != e.Value {
    t.Error("Expected ", e.Value, ", got ", e2.Value)
  }
  if e2.EventType !=  e.EventType {
    t.Error("Expected ", e.EventType, ", got ", e2.EventType)
  }
  if e2.Number != e.Number {
    t.Error("Expected ", e.Number, ", got ", e2.Number)
  }

  // Negative
  e.Time = 0xffffffff
  e.Value = -0xff
  e.EventType = 0xff
  e.Number = 0xff

  data, _ = e.MarshalBinary()
  e2.UnMarshalBinary(data)
  if e2.Time != e.Time {
    t.Error("Expected ", e.Time ,", got ", e2.Time)
  }
  if e2.Value != e.Value {
    t.Error("Expected ", e.Value, ", got ", e2.Value)
  }
  if e2.EventType !=  e.EventType {
    t.Error("Expected ", e.EventType, ", got ", e2.EventType)
  }
  if e2.Number != e.Number {
    t.Error("Expected ", e.Number, ", got ", e2.Number)
  }
}
