package abstract_factory

import "fmt"

type Shape interface {
	Draw()
}
type circle struct {
}

func (c *circle) Draw() {
	fmt.Println("i am circle")
}

type square struct {
}

func (s *square) Draw() {
	fmt.Println("i am square")
}

type ellipse struct {
}

func (e *ellipse) Draw() {
	fmt.Println("i am ellipse")
}

type rectangle struct {
}

func (r *rectangle) Draw() {
	fmt.Println("i am rectangle")
}

// ShapeFactory is an interface for a factory which can be used
// to create curved and straight shapes.
type ShapeFactory interface {
	CreateCurvedShape() Shape
	CreateStraightShape() Shape
}

type simpleShapeFactory struct {
}

// NewSimpleShapeFactory creates a new simpleShapeFactory.
func NewSimpleShapeFactory() ShapeFactory {
	return &simpleShapeFactory{}
}

func (s *simpleShapeFactory) CreateCurvedShape() Shape {
	return &circle{}
}

func (s *simpleShapeFactory) CreateStraightShape() Shape {
	return &square{}
}

type robustShapeFactory struct {
}

// NewRobustShapeFactory creates a new robustShapeFactory.
func NewRobustShapeFactory() ShapeFactory {
	return &robustShapeFactory{}
}

func (s *robustShapeFactory) CreateCurvedShape() Shape {
	return &ellipse{}
}

func (s *robustShapeFactory) CreateStraightShape() Shape {
	return &rectangle{}
}
