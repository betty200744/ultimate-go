package factory_method

import "fmt"

/*

 更灵活的创建对象的方法

The Factory Method pattern is a design pattern used to define a runtime interface for creating an object. It’s called a factory because it creates various types of objects without necessarily knowing what kind of object it creates or how to create it.

Purpose
Allows the sub-classes to choose the type of objects to create at runtime
It provides a simple way of extending the family of objects with minor changes in application code.
Promotes the loose-coupling by eliminating the need to bind application-specific structs into the code
*/

type IGun interface {
	getName() string
	getPower() int
}

type Gun struct {
	name  string
	power int
}

func (g *Gun) getName() string {
	return g.name
}
func (g *Gun) getPower() int {
	return g.power
}

type ak47 struct {
	Gun
}
type m16 struct {
	Gun
}

func GetGun(gunType string) (IGun, error) {
	switch gunType {
	case "ak47":
		return &ak47{Gun{
			name:  "ak47",
			power: 9,
		}}, nil
	case "m16":
		return &m16{Gun{
			name:  "m16",
			power: 10,
		}}, nil
	default:
		return &ak47{Gun{
			name:  "ak47",
			power: 9,
		}}, nil
	}
}
func PrintDetails(g IGun) {
	fmt.Printf("Gun: %s \n", g.getName())
	fmt.Printf("Power: %d \n", g.getPower())
}
