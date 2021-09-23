package generator

func NewIntGenerator() func() int {
	var next int
	return func() int {
		next++
		return next
	}
}

type VertexID int

func NewVertexIDGeneratorV1() func() VertexID {
	var next int
	return func() VertexID {
		next++
		return VertexID(next)
	}
}

func NewVertexIDGenerator() func() VertexID {
	gen := NewIntGenerator()
	return func() VertexID {
		return VertexID(gen())
	}
}
