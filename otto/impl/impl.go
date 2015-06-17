package otto

import (
	"../adt"
	"github.com/robertkrimen/otto"
)

func NewFactory() iotto.Factory {
	return &Factory{}
}

type Factory struct {
}

func (f *Factory) New() iotto.Otto {
	return &Otto{otto.New()}
}

func (f *Factory) Run(src interface{}) (iotto.Otto, iotto.Value, error) {
	o, v, err := otto.Run(src)
	return &Otto{o}, &Value{v}, err
}

func (f *Factory) FalseValue() iotto.Value {
	return &Value{otto.FalseValue()}
}

func (f *Factory) NaNValue() iotto.Value {
  return &Value{otto.NaNValue()}
}

func (f *Factory) NullValue() iotto.Value {
  return &Value{otto.NullValue()}
  NullValue

func (f *Factory) ToValue(value interface{}) (iotto.Value, error) {
  v, err := otto.ToValue(value)
  return &Value{v}, err
}

func (f *Factory) TrueValue() iotto.Value {
  return &Value{otto.TrueValue()}
}

func (f *Factory) UndefinedValue() iotto.Value {
  return &Value{otto.UndefinedValue()}
}

type Otto struct{
  o *otto.Otto
}

func (o *Otto) Interupt() chan func() {

}

func (o *Otto) Call(source string, this interface{}, argumentList ...interface{}) (iotto.Value, error) {

}

func (o *Otto) Compile(filename string, src interface{}) (iotto.Script, error) {

}

func (o *Otto) Copy() iotto.Otto {

}

func (o *Otto) Get(name string) (iotto.Value, error) {

}

func (o *Otto) Object(source string) (iotto.Object, error) {

}

func (o *Otto) Run(src interface{}) (iotto.Value, error) {

}

func (o *Otto) Set(name string, value interface{}) error {

}

func (o *Otto) ToValue(value interface{}) (iotto.Value, error) {

}

type Script struct {
  script *otto.Script
}

func (s *Script) String() string {
  return s.script.String()
}

type Value struct{
  value *otto.Value
}

func (v *Value) Call(this iotto.Value, argumentList ...interface{}) (iotto.Value, error) {
  self, ok := this.(Value)
  if !ok {
    panic("This is not a Value.")
  }
  out, err := v.value.Call(self.value)
}

func (v *Value) Class() string {

}

func (v *Value) Export() (interface{}, error) {

}

func (v *Value) IsBoolean() bool {

}

func (v *Value) IsDefined() bool {

}

func (v *Value) IsFunction() bool {

}

func (v *Value) IsNaN() bool {

}

func (v *Value) IsNull() bool {

}

func (v *Value) IsNumber() bool {

}

func (v *Value) IsObject() bool {

}

func (v *Value) IsPrimitive() bool {

}

func (v *Value) IsString() bool {

}

func (v *Value) IsUndefined() bool {

}

func (v *Value) Object() Object {

}

func (v *Value) String() string {

}

func (v *Value) ToBoolean() (bool, error) {

}

func (v *Value) ToFloat() (float64, error) {

}

func (v *Value) ToInteger() (int64, error) {

}

func (v *Value) ToString() (string, error) {

}

type Object struct{
  obj *otto.Object
}

func (o *Object) Call(name string, argumentList ...interface{}) (iotto.Value, error) {

}

func (o *Object) Class() string {

}

func (o *Object) Get(name string) (iotto.Value, error) {

}

func (o *Object) Keys() []string {

}

func (o *Object) Set(name string, value interface{}) error {

}

func (o *Object) Value() iotto.Value {

}

type FunctionCall struct {
  fn *otto.FunctionCall
}

func (f *FunctionCall) This() iotto.Value {

}

func (f *FunctionCall) Argument(index int) iotto.Value {

}

func (f *FunctionCall) ArgumentList() []iotto.Value {

}

func (f *FunctionCall) Otto() iotto.Otto {

}
