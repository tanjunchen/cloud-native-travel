package util

import (
	"strings"

	"github.com/ghodss/yaml"
	"github.com/gogo/protobuf/types"
	"github.com/golang/protobuf/jsonpb"
)

// JSON2Struct json string transform types.Struct
func JSON2Struct(str string) (*types.Struct, error) {
	result := &types.Struct{}

	m := jsonpb.Unmarshaler{}
	err := m.Unmarshal(strings.NewReader(str), result)
	return result, err
}

// YAML2Struct yaml string transform types.Struct
func YAML2Struct(str string) (*types.Struct, error) {
	b, err := yaml.YAMLToJSON([]byte(str))
	if err != nil {
		return nil, err
	}
	return JSON2Struct(string(b))
}
