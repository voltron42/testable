package yaml

import (
	"../adt"
	"gopkg.in/yaml.v2"
)

func NewYaml() iyaml.Yaml {
	return &Yaml{}
}

type Yaml struct{}

func (y *Yaml) Marshal(in interface{}) (out []byte, err error) {
	return yaml.Marshal(in)
}

func (y *Yaml) Unmarshal(in []byte, out interface{}) (err error) {
	return yaml.Unmarshal(in, out)
}
