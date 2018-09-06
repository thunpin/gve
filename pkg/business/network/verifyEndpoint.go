package network

import (
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"

	"github.com/thunpin/gve/pkg/domain"
)

const KindTypeToTestNetwork = "serviceentry"

func VerifyEndpoint(element domain.Element) error {
	for _, item := range element.Items {
		if strings.ToLower(item.Kind) == KindTypeToTestNetwork {
			err := verifySpec(item.Spec)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func verifySpec(spec domain.Spec) error {
	for _, host := range spec.Hosts {
		if strings.Contains(host, "*") {
			log.Println("[WARNING] Skipping:" + host)
		} else {
			err := verifyHost(host, spec.Ports)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func verifyHost(host string, ports []domain.Port) error {
	for _, port := range ports {
		var err error
		switch strings.ToLower(port.Protocol) {
		case "http":
			err = httpConn("http://" + host + ":" + port.Number)
		case "https":
			err = httpConn("https://" + host + ":" + port.Number)
		case "tcp":
			err = tcpConn(host, port.Number)
		default:
			return errors.New(fmt.Sprintf("invalid port protocol: %s of host %s. Only accept HTTP/HTTPS/TCP", port, host))
		}

		addressToLog := port.Protocol + ":" + host + ":" + port.Number
		if err != nil {
			fmt.Println(err)
			return errors.New(addressToLog)
		}
	}
	return nil
}

func httpConn(address string) error {
	data, err := http.Get(address)
	if err != nil {
		return err
	}
	if (data.StatusCode / 100) == 5 {
		return errors.New(address)
	}
	return nil
}

func tcpConn(host string, port string) error {
	_, err := net.Dial("tcp", host+":"+port)
	return err
}
