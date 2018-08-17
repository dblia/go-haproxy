/*

  Unit tests for the utils.go module.

*/
package haproxy

import "testing"

func TestIsValidIPv4Addr(t *testing.T) {
	var validIPv4Addrs = []string{
		"127.0.0.1:80",
		"1.2.3.4:80",
		"localhost:80",
	}

	for _, ipAddr := range validIPv4Addrs {
		assertNil(isValidIPv4Addr(ipAddr), t)
	}
}

func TestNotIsValidIPv4Addr(t *testing.T) {
	var invalidIPv4Addrs = map[string]string{
		"[127.0.0.1]:80":      "Invalid IPv4 address: '[127.0.0.1]'",
		"1:2.3.4:80":          "Invalid address format: '1:2.3.4:80'",
		"koko.example.com:80": "Invalid IPv4 address: 'koko.example.com'", // will be supported soon
	}

	for ipAddr, retVal := range invalidIPv4Addrs {
		assertEqual(isValidIPv4Addr(ipAddr).Error(), retVal, t)
	}
}

func TestIsValidIPv6Addr(t *testing.T) {
	var validIPv6Addrs = []string{
		"[::]:80",
		"[::1]:80",
		"[fc00::]:80",
		"[2a03:e40:2a:400::97]:80",
	}

	for _, ipAddr := range validIPv6Addrs {
		assertNil(isValidIPv6Addr(ipAddr), t)
	}
}

func TestNotIsValidIPv6Addr(t *testing.T) {
	var invalidIPv6Addrs = map[string]string{
		"localhost:80":   "Invalid address format: 'localhost:80'",
		"[localhost]:80": "Invalid address format: '[localhost]:80'",
		"[10.0.0.42]:80": "Invalid address format: '[10.0.0.42]:80'",
		"[fc00::g]:80":   "Invalid address format: '[fc00::g]:80'",
	}

	for ipAddr, retVal := range invalidIPv6Addrs {
		assertEqual(isValidIPv6Addr(ipAddr).Error(), retVal, t)
	}
}
