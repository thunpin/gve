package parser

import "testing"

func TestReturnErrorWhenUsingInvalidYmlFormat(t *testing.T) {
	var data = `
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

	_, err := Yaml(data)
	if err == nil {
		t.Fail()
	}
}

func TestMustReturnsEmptyItemsInElement(t *testing.T) {
	var data = `
	invalid: tes
	item:
	`

	element, err := Yaml(data)
	if err != nil && len(element.Items) != 0 {
		t.Fail()
	}
}
