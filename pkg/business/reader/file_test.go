package reader

import (
	"os"
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
	gopath := os.Getenv("GOPATH")
	data, err := FromFile(gopath + "/src/github.com/thunpin/gve/istio.yaml")
	if err != nil {
		t.Fail()
	}

	if !strings.Contains(data, "apiVersion") {
		t.Fail()
	}
}
