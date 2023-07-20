package struct

import "math"

// We're creating a new type Rectangle which is a struct.
// A struct is just a named collection of fields where you can store data.
type Rectangle struct {
    Width  float64
    Height float64
}

// The syntax for declaring methods is almost the same as functions and that's because they're so similar.
// The only difference is the syntax of the method receiver func (receiverName ReceiverType) MethodName(args).
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
    return 2 * (r.Width + r.Height)
}

type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return c.Radius * c.Radius * math.Pi // math.Pi is a constant in the math package
}

type Triangle struct {
    Base  float64
    Height float64
}

func (r Triangle) Area() float64 {
    return r.Base * r.Height * 0.5
}

// We're creating a new type just like we did with Rectangle and Circle but this time it is an interface
// rather than a struct.
type Shape interface {
    Area() float64
}
