package svrman

import (
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var (
	defaultSeverMan = NewServerMan()
)

type ServerMan struct {
	svrs []Server
}

type Server interface {
	Serve() error
	Stop() error
}

func NewServerMan() *ServerMan {
	return &ServerMan{
		svrs: make([]Server, 0, 1),
	}
}

func (s *ServerMan) RegisterServer(server Server) {
	s.svrs = append(s.svrs, server)
}

func (s *ServerMan) Start() (err error) {
	wg := sync.WaitGroup{}
	done := make(chan struct{})
	errChan := make(chan error)
	go func() {
		if err := handleSysSignal(); nil != err {
			errChan <- err
		}
	}()
	wg.Add(len(s.svrs))
	for _, svr := range s.svrs {
		go func(s Server) {
			defer wg.Done()
			if err = s.Serve(); nil != err {
				errChan <- err
			}
		}(svr)
	}
	go func() {
		wg.Wait()
		close(done)
	}()
	select {
	case err = <-errChan:
		return
	case <-done:
		return
	}
}

func (s *ServerMan) Stop() (err error) {
	for _, svr := range s.svrs {
		if err = svr.Stop(); nil != err {
			return
		}
	}
	return
}

func RegisterServer(server Server) {
	defaultSeverMan.RegisterServer(server)
}

func Start() error {
	return defaultSeverMan.Start()
}

func Stop() error {
	return defaultSeverMan.Stop()
}

func handleSysSignal() error {
	sChan := make(chan os.Signal)
	for {
		signal.Notify(sChan, os.Interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
		sig := <-sChan
		switch sig {
		case os.Interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
			return Stop()
		}
	}
}
