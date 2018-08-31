package parser

import (
	"github.com/thunpin/gve/pkg/domain"
	yaml "gopkg.in/yaml.v2"
)

func Yaml(data string) (domain.Element, error) {
	element := domain.Element{}
	err := yaml.Unmarshal([]byte(data), &element)

	return element, err
}
