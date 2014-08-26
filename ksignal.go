package kutil

import (
	"os"
	"os/signal"
	"syscall"
)

const (
	// More invented values for signals
	FORHUP  = syscall.Signal(0x1) //reload
	FORINT  = syscall.Signal(0x2)
	FORQUIT = syscall.Signal(0x3)
	FORILL  = syscall.Signal(0x4)
	FORTRAP = syscall.Signal(0x5)
	FORABRT = syscall.Signal(0x6)
	FORBUS  = syscall.Signal(0x7)
	FORFPE  = syscall.Signal(0x8)
	FORKILL = syscall.Signal(0x9)
	FORSEGV = syscall.Signal(0xb)
	FORPIPE = syscall.Signal(0xd)
	FORALRM = syscall.Signal(0xe)
	GORTERM = syscall.Signal(0xf)
)

type SignalNofity chan os.Signal

func (sn *SignalNofity) Wait() os.Signal {
	if *sn == nil {
		*sn = make(chan os.Signal, 1)
		signal.Notify(*sn, FORHUP, FORINT, FORQUIT, FORILL, FORTRAP, FORABRT, FORBUS, FORFPE, FORKILL, FORSEGV, FORPIPE, FORALRM, GORTERM)
	}
	s := <-*sn
	return s
}
