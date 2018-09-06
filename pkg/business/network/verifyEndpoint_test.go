package network

import (
	"testing"

	"github.com/thunpin/gve/pkg/domain"
)

func TestMustReturnWarningWhenTheHostContainsAsterisk(t *testing.T) {
	element := testCreateEmptyElement()
	element.Items[0].Spec.Hosts = append(element.Items[0].Spec.Hosts, "*.google.com")
	err := VerifyEndpoint(element)
	if err != nil {
		t.Fail()
	}
}

func TestMustReturnErrorWhenItemHostIsInvalid(t *testing.T) {
	element := testCreateEmptyElement()
	element.Items[0].Spec.Hosts = append(element.Items[0].Spec.Hosts, "oogle.com")
	element.Items[0].Spec.Hosts = append(element.Items[0].Spec.Hosts, "google.com")
	element.Items[0].Spec.Ports = append(element.Items[0].Spec.Ports, domain.Port{"80", "HTTP"})
	err := VerifyEndpoint(element)
	if err == nil {
		t.Fail()
	}
}

func TestMustReturnValidResultsWhenTheElementIsValid(t *testing.T) {
	element := testCreateEmptyElement()
	element.Items[0].Spec.Hosts = append(element.Items[0].Spec.Hosts, "google.com")
	element.Items[0].Spec.Ports = append(element.Items[0].Spec.Ports, domain.Port{"443", "HTTPS"})
	err := VerifyEndpoint(element)
	if err != nil {
		t.Fail()
	}
}

func testCreateEmptyElement() domain.Element {
	element := domain.Element{}
	item := domain.Item{}
	item.Kind = "ServiceEntry"
	item.Spec = domain.Spec{make([]string, 0), make([]domain.Port, 0)}

	element.Items = append(element.Items, item)
	return element
}
