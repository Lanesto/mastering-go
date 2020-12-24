package sharedrpc

// MyFloats define pair of value
type MyFloats struct {
	A1, A2 float64
}

// MyInterface define interface to use between client and server
type MyInterface interface {
	Multiply(arguments *MyFloats, reply *float64) error
	Power(arguments *MyFloats, reply *float64) error
}
