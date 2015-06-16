package iyaml

type Yaml interface {
	Marshal(in interface{}) (out []byte, err error)
	Unmarshal(in []byte, out interface{}) (err error)
}
