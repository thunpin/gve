package network

import (
	"testing"

	"github.com/thunpin/gve/pkg/domain"
)

func TestMustPrintWarningWhenTheHostContainsAsterisk(t *testing.T) {
	element := testCreateEmptyElement()
	element.Items[0].Spec.Hosts = append(element.Items[0].Spec.Hosts, "*.google.com")
}

func testCreateEmptyElement() domain.Element {
	element := domain.Element{}
	item := domain.Item{}
	item.Kind = "ServiceEntry"
	item.Spec = domain.Spec{make([]string, 0), make([]domain.Port, 0)}

	element.Items = append(element.Items, item)
	return element
}
