// Command vscp provides a scp-like utility for copying files over VM
// sockets.  It is meant to show example usage of package vsock, but is
// also useful in scenarios where a virtual machine does not have
// networking configured, but VM sockets are available.
package main

import (
	"log"

	"github.com/mdlayher/vsock"
)

func main() {
	l, err := vsock.ListenUDP(8080, nil)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer l.Close()
}
