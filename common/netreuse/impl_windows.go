package netreuse

import (
	"context"
	"net"
)

func dial(ctx context.Context, dialer net.Dialer, network, address string) (net.Conn, error) {
	return dialer.DialContext(ctx, network, address)
}

// on windows, we just use the regular functions. sources
// vary on this-- some claim port reuse behavior is on by default
// on some windows systems. So we try. may the force be with you.
func available() bool {
	return true
}
