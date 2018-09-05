package reader

import (
	"strings"
	"testing"
)

func TestReturnErrorIfThePathIsInvalid(t *testing.T) {
	_, err := FromFile("blah.yaml")
	if err == nil {
		t.Fail()
	}
}

func TestReturnTheDataFileWhenThePathIsValid(t *testing.T) {
	data, err := FromFile(dir + "../../github.com/thunpin/gve/istio.yaml")
	if err != nil {
		t.Fail()
	}

	if !strings.Contains(data, "apiVersion") {
		t.Fail()
	}
}
