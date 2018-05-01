package battle

import (
  "testing"
)


func TestAdd(t *testing.T) {
  var v int
  v = Add(1, 6)
  if v != 7 {
    t.Error("Expected 6, got ", v)
  }
}