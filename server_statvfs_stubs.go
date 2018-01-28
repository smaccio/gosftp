// +build !darwin,!linux gccgo

package gosftp

import (
	"syscall"
)

func (p sshFxpExtendedPacketStatVFS) respond(svr *Server) error {
	return syscall.ENOTSUP
}
