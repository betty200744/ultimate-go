package abstract_factory

import "testing"

func TestNewSimpleShapeFactory(t *testing.T) {
	simple := NewSimpleShapeFactory()
	straightShape := simple.CreateStraightShape()
	straightShape.Draw()
	curvedShape := simple.CreateCurvedShape()
	curvedShape.Draw()
}
func TestNewRobustShapeFactory(t *testing.T) {
	robust := NewRobustShapeFactory()
	straightShape := robust.CreateStraightShape()
	straightShape.Draw()
	curvedShape := robust.CreateCurvedShape()
	curvedShape.Draw()
}
