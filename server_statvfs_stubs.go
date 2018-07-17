// +build !darwin,!linux

package gosftp

import (
	"syscall"
)

func (p sshFxpExtendedPacketStatVFS) respond(svr *Server) error {
	return syscall.ENOTSUP
}
