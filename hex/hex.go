package hex

type Hex struct {
  x, y, z int
}

func (h *Hex) AxialCoordinates() (r, q int) {
  r = h.z
  q = h.x
  return
}

func NewFromAxialCoordinates(r, q int) *Hex {
  h := new(Hex)
  h.x = q
  h.z = r
  h.y = -q-r
  return h
}
