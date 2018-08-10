/*

  Module implementing the API with HAProxy's UNIX Socket interface.

*/

package haproxy

import "fmt"
import "io/ioutil"
import "net"
import "strings"
import "time"

// A HAProxyClient instance represents a HAProxy connection object.
type HAProxyClient struct {
	AfNet   string
	Address string
	Timeout time.Duration
	conn    net.Conn
}

// Connects to the address on the given network type.
// The library can listen to a either a TCP port or a UNIX socket.
func (h *HAProxyClient) connect() (err error) {
	switch h.AfNet {
	case AFNET_UNIX:
	case AFNET_TCP4:
		err = isValidIPv4Addr(h.Address)
	case AFNET_TCP6:
		err = isValidIPv6Addr(h.Address)
	default:
		err = fmt.Errorf("Supported network types are '%s', '%s', and '%s'\n",
			AFNET_UNIX, AFNET_TCP4, AFNET_TCP6)
	}
	if err == nil {
		h.conn, err = net.DialTimeout(h.AfNet, h.Address, h.Timeout)
	}
	return err
}

// Send the given command to HAProxy over the UNIX admin socket.
func (h *HAProxyClient) SendCommand(cmd string) (string, error) {
	if err := h.connect(); err != nil {
		return EMPTY_RESP, err
	}
	defer h.conn.Close()

	//. ensure a newline char will be always appended to the command
	command := []byte(strings.Replace(cmd, "\n", "", -1) + "\n")

	h.conn.SetWriteDeadline(time.Now().Add(h.Timeout))
	if _, err := h.conn.Write(command); err != nil {
		return EMPTY_RESP, err
	}

	h.conn.SetReadDeadline(time.Now().Add(h.Timeout))
	response, err := ioutil.ReadAll(h.conn)
	if err != nil {
		return EMPTY_RESP, err
	}

	return string(response), nil
}
