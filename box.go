package golang_united_school_homework

import (
	"errors"
	"fmt"
	"reflect"
)

var (
	errorCapacity       = errors.New("it goes out of the shapesCapacity range")
	errorIndex          = errors.New("shape by index went out of the range")
	errorExtractByIndex = errors.New("shape by index doesn't exist or index went out of the range ExtractByIndex")
	errorReplaceByIndex = errors.New("shape by index doesn't exist or index went out of the range ReplaceByIndex")
	errorNoCircle       = errors.New("circles are not exist in the list")
)

// box contains list of shapes and able to perform operations on them
type box struct {
	shapes         []Shape
	shapesCapacity int // Maximum quantity of shapes that can be inside the box.
}

// NewBox creates new instance of box
func NewBox(shapesCapacity int) *box {
	return &box{
		shapesCapacity: shapesCapacity,
	}
}

// AddShape adds shape to the box
// returns the error in case it goes out of the shapesCapacity range.
func (b *box) AddShape(shape Shape) error {
	if b.shapesCapacity <= len(b.shapes) {
		return fmt.Errorf("%w", errorCapacity)
	}
	b.shapes = append(b.shapes, shape)
	return nil
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	if i < len(b.shapes) {
		return b.shapes[i], nil
	}
	return nil, fmt.Errorf("%w", errorIndex)
}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	s, err := b.GetByIndex(i)
	if err != nil {
		return s, fmt.Errorf("%w", errorExtractByIndex)
	}
	b.shapes = append(b.shapes[:i], b.shapes[i+1:]...)
	return s, nil
}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	s, err := b.GetByIndex(i)
	if err != nil {
		return nil, fmt.Errorf("%w", errorReplaceByIndex)
	}
	b.shapes[i] = shape
	return s, nil

}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	var sum float64
	for i := 0; i < len(b.shapes); i++ {
		sum += b.shapes[i].CalcPerimeter()
	}
	return sum
}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	var sum float64
	for i := 0; i < len(b.shapes); i++ {
		sum += b.shapes[i].CalcArea()
	}
	return sum
}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	count := 0
	typeCircle := reflect.TypeOf(&Circle{})
	for i := 0; i < len(b.shapes); i++ {
		typeShape := reflect.TypeOf(b.shapes[i])
		if typeShape == typeCircle {
			count += 1
			b.shapes = append(b.shapes[:i], b.shapes[i+1:]...)
			i--
		}
	}
	if count == 0 {
		return fmt.Errorf("%w", errorNoCircle)
	}
	return nil
}
