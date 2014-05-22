package hex

// Hexagon (6-sided polygons) using cubic coordinates
type Hex struct {
  x, y, z int
}

// Converts hexagon cubic coordinates to axial ones
func (h *Hex) AxialCoordinates() (r, q int) {
  r = h.z
  q = h.x
  return
}

// Creates an hexagon cubic coordinates from axial coordinates
func NewFromAxialCoordinates(r, q int) *Hex {
  h := new(Hex)
  h.x = q
  h.z = r
  h.y = -q-r
  return h
}

// Finds adjacent hexagons
func (h *Hex) NeighborCoordinates() ([6][3]int) {
  neighbors := [6][3]int{
    {+1, -1,  0}, {+1,  0, -1}, { 0, +1, -1},
    {-1, +1,  0}, {-1,  0, +1}, { 0, -1, +1},
  }
  var coords [6][3]int
  for i, value := range neighbors {
    coords[i] = [3]int{h.x + value[0], h.y + value[1], h.z + value[2]}
  }
  return coords
}
