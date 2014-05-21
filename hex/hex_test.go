package hex

import "testing"

func TestAxialCoordinates(t *testing.T) {
  h := Hex{x: 0, y: 1, z: -1}
  r, q := h.AxialCoordinates()
  if q != h.x {
    t.Error("Expected q to equal x(0), got", q)
  }
  if r != h.z {
    t.Error("Expected r to equal z(-1), got", r)
  }
}

func TestNewFromAxialCoordinates(t *testing.T) {
  r, q := -1, 0
  h := NewFromAxialCoordinates(r, q)
  if h.x != q {
    t.Error("Expected x to equal q(0), got", h.x)
  }
  if h.z != r {
    t.Error("Expected z to equal r(-1), got", h.z)
  }
  if h.y != 1 {
    t.Error("Expected y to equal q(1), got", h.y)
  }
}
