package parser

import (
	"testing"
)

func TestYamReturnEmptyWhenItemsIsNotFound(t *testing.T) {
	data := `
    invalid: tes
    item:
    - apiversion: 001
      kind: ServiceEntry
      spec:
        hosts:
        - github.com
        ports:
        - number: 443
          protocol: HTTPS
    - apiversion: 002
      kind: ServiceEntry
      spec:
        hosts:
        - google.com
        ports:
        - number: 443
          protocol: HTTPS
     `

	element, err := Yaml(data)
	if err != nil && len(element.Items) != 0 {
		t.Fail()
	}
}

func TestYamlMustReturnElementWithTwoItens(t *testing.T) {
	data := `
    items:
    - apiversion: 001
      kind: ServiceEntry
      spec:
        hosts:
        - google.com
        ports:
        - number: 443
          protocol: HTTPS
    - apiversion: 002
      kind: ServiceEntry
      spec:
        hosts:
        - github.com
        ports:
        - number: 443
          protocol: HTTPS
    `
	element, err := Yaml(data)
	if err != nil && len(element.Items) != 2 {
		t.Fail()
	}
}

func TestYamlMustReturnElementWithTwoItensFirstWithOnePortSecondWithTwoPorts(t *testing.T) {
	data := `
    items:
    - apiversion: 001
      kind: ServiceEntry
      spec:
        hosts:
        - google.com
        ports:
        - number: 443
          protocol: HTTPS
    - apiversion: 002
      kind: ServiceEntry
      spec:
        hosts:
        - github.com
        ports:
        - number: 80
          protocol: HTTP
        - number: 443
          protocol: HTTPS
    `
	element, err := Yaml(data)
	if err != nil || len(element.Items) != 2 ||
		len(element.Items[0].Spec.Ports) != 1 ||
		len(element.Items[1].Spec.Ports) != 2 {

		t.Fail()
	}
}
