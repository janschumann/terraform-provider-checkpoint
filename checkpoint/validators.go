package checkpoint

import (
	"fmt"
	"net"
)

// represents an IP4 address - it adds an error otherwise
func validateIP4Address(v interface{}, k string) (ws []string, errors []error) {
	value := v.(string)
	for i := 0; i < len(value); i++ {
		switch value[i] {
		case '.':
			ip := net.ParseIP(value)
			if ip == nil {
				errors = append(errors, fmt.Errorf(
					"%q must contain a valid IPv4 address", k))
			}
			return
		case ':':
			errors = append(errors, fmt.Errorf(
				"%q must contain a valid IPv4 address but looks like IPv6", k))
			return
		}
	}

	return
}

// represents an IP6 address - it adds an error otherwise
func validateIP6Address(v interface{}, k string) (ws []string, errors []error) {
	value := v.(string)
	for i := 0; i < len(value); i++ {
		switch value[i] {
		case '.':
			ip := net.ParseIP(value)
			if ip == nil {
				errors = append(errors, fmt.Errorf(
					"%q must contain a valid IPv4 address", k))
			}
			return
		case ':':
			errors = append(errors, fmt.Errorf(
				"%q must contain a valid IPv4 address but looks like IPv6", k))
			return
		}
	}

	return
}
