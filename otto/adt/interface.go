package iotto

import "github.com/robertkrimen/otto"

var ErrVersion = otto.ErrVersion

type Error interface {
	Error() string
	String() string
}
type FunctionCall interface {
	This() Value
	Argument(index int) Value
	ArgumentList() []Value
	Otto() *Otto
}
type Object interface {
	Call(name string, argumentList ...interface{}) (Value, error)
	Class() string
	Get(name string) (Value, error)
	Keys() []string
	Set(name string, value interface{}) error
	Value() Value
}
type Factory interface {
	New() *Otto
	Run(src interface{}) (*Otto, Value, error)
	FalseValue() Value
	NaNValue() Value
	NullValue() Value
	ToValue(value interface{}) (Value, error)
	TrueValue() Value
	UndefinedValue() Value
}
type Otto interface {
	Interupt() chan func()
	Call(source string, this interface{}, argumentList ...interface{}) (Value, error)
	Compile(filename string, src interface{}) (*Script, error)
	Copy() *Otto
	Get(name string) (Value, error)
	Object(source string) (*Object, error)
	Run(src interface{}) (Value, error)
	Set(name string, value interface{}) error
	ToValue(value interface{}) (Value, error)
}
type Script interface {
	String() string
}
type Value interface {
	Call(this Value, argumentList ...interface{}) (Value, error)
	Class() string
	Export() (interface{}, error)
	IsBoolean() bool
	IsDefined() bool
	IsFunction() bool
	IsNaN() bool
	IsNull() bool
	IsNumber() bool
	IsObject() bool
	IsPrimitive() bool
	IsString() bool
	IsUndefined() bool
	Object() *Object
	String() string
	ToBoolean() (bool, error)
	ToFloat() (float64, error)
	ToInteger() (int64, error)
	ToString() (string, error)
}
