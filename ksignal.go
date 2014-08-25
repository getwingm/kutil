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

type SignalProc func(os.Signal, ...interface{}) bool

func WaitForSignalEx(spfn SignalProc, p ...interface{}) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, FORHUP, FORINT, FORQUIT, FORILL, FORTRAP, FORABRT, FORBUS, FORFPE, FORKILL, FORSEGV, FORPIPE, FORALRM, GORTERM)
	bLoop := true
	for bLoop {
		s := <-c
		bLoop = spfn(s, p...)
	}
}

var sc chan os.Signal

func WaitForSignal() os.Signal {
	if sc == nil {
		sc = make(chan os.Signal, 1)
		signal.Notify(sc, FORHUP, FORINT, FORQUIT, FORILL, FORTRAP, FORABRT, FORBUS, FORFPE, FORKILL, FORSEGV, FORPIPE, FORALRM, GORTERM)
	}
	s := <-sc
	return s
}
