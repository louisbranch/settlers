package hex

import "testing"

func TestAxialCoordinates(t *testing.T) {
  h := Hex{x: 0, y: 1, z: -1}
  r, q := h.AxialCoordinates()
  if q != h.x {
    t.Error("Expected q to equal", h.x ,"got", q)
  }
  if r != h.z {
    t.Error("Expected r to equal", h.z, "got", r)
  }
}

func TestNewFromAxialCoordinates(t *testing.T) {
  r, q := -1, 0
  h := NewFromAxialCoordinates(r, q)
  if h.x != q {
    t.Error("Expected x to equal", q ,"got", h.x)
  }
  if h.z != r {
    t.Error("Expected z to equal", r ,"got", h.z)
  }
  if h.y != 1 {
    t.Error("Expected y to equal", 1 ,"got", h.y)
  }
}

func TestNeighborCoordinates(t *testing.T) {
  h := Hex{x: 0, y: 1, z: -1}
  coords := [6][3]int{
    {+1,  0, -1}, {+1, +1, -2}, { 0, +2, -2},
    {-1, +2, -1}, {-1, +1,  0}, { 0,  0,  0},
  }
  neighbors := h.NeighborCoordinates()
  if coords != neighbors {
    t.Error("Expected", coords, "got ", neighbors)
  }
}
