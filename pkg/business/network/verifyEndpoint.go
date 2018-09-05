package network

import (
	"strings"

	"github.com/thunpin/gve/pkg/domain"
)

const KindTypeToTestNetwork = "serviceentry"

func VerifyEndpoint(element domain.Element) ([]Result, error) {
	results := make([]Result, 0)
	for _, item := range element.Items {
		if strings.ToLower(item.Kind) == KindTypeToTestNetwork {
			_, err := verifySpec(item.Spec)
			if err != nil {
				return nil, err
			}
		}
	}
	return results, nil
}

func verifySpec(spec domain.Spec) ([]Result, error) {
	results := make([]Result, 0)
	for _, host := range spec.Hosts {
		if strings.Contains(host, "*") {
			results = append(results, Warning{host})
		} else {
			returnedResults, err := verifyHost(host, spec.Ports)
			if err != nil {
				return nil, err
			}
			results = append(results, returnedResults...)
		}
	}

	return results, nil
}

func verifyHost(host string, ports []domain.Port) ([]Result, error) {
	return nil, nil
}
