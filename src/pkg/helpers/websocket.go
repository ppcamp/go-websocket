package helpers

import (
	"fmt"
	"net"
	"strings"
)

func WebsocketAddress(addr net.Addr) (string, error) {
	id := strings.Split(addr.String(), "]:")
	if len(id) < 2 {
		return "", fmt.Errorf("some error with the remote address")
	}
	return id[1], nil
}
