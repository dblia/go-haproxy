/*

  Utility functions for the HAProxy client.

*/
package haproxy

import "fmt"
import "net"
import "regexp"
import "strconv"
import "strings"

// Regular expression definining a valid IPv6 address format.
var ipv6Expr = regexp.MustCompile(`^\[(?P<ip>[:0-9a-f]+)\]:(?P<port>[0-9]+)`)

// Check is the given expression has a valid IPv6 format.
// Example of valid IPv6 format: `[2a03:e40:2a:400::97]:8000`
func isValidIPv6Addr(ipAddr string) (err error) {
	matched := ipv6Expr.FindStringSubmatch(ipAddr)
	if matched == nil {
		err = fmt.Errorf("Invalid address format: '%s'", ipAddr)
	} else if e := net.ParseIP(matched[1]); e == nil {
		err = fmt.Errorf("Invalid IPv6 address: '%s'", matched[1])
	} else if _, e := strconv.Atoi(matched[2]); e != nil {
		err = fmt.Errorf("Invalid port: '%s'", matched[2])
	}
	return err
}

// Check is the given expression has a valid IPv4 format.
// Example of valid IPv4 format: `1.2.3.4:8000`
func isValidIPv4Addr(ipAddr string) (err error) {
	data := strings.Split(ipAddr, ":")
	if len(data) != 2 {
		err = fmt.Errorf("Invalid address format: '%s'", ipAddr)
	} else if e := net.ParseIP(data[0]); e == nil {
		err = fmt.Errorf("Invalid IPv4 address: '%s'", data[0])
	} else if _, e := strconv.Atoi(data[1]); e != nil {
		err = fmt.Errorf("Invalid port: '%s'", data[1])
	}
	return err
}
