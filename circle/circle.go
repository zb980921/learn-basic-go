package circle

import (
	"fmt"
	"math"
)

type Circle struct {
	radius float64
}

func NewCircle(radius float64) *Circle {
	c := new(Circle)
	c.radius = radius
	return c
}

func (c *Circle) GetRadius() float64 {
	return c.radius
}

func (c *Circle) SetRadius(newRadius float64) {
	c.radius = newRadius
}

func (c *Circle) GetArea() float64 {
	return math.Pi * c.radius * c.radius
}

func (c *Circle) ToString() string {
	return fmt.Sprintf("半径: %f米, 面积: %f平方米", c.radius, c.GetArea())
}
