package main

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/kardianos/service"
	"github.com/mlctrez/servicego"
)

func main() {

	servicego.Run(&goRoutines{})
}

type goRoutines struct {
	servicego.Defaults
	context context.Context
	cancel  func()
	wg      *sync.WaitGroup
	db      *sampleDatabase
}

func (g *goRoutines) Start(s service.Service) error {

	defer g.logEntryExit("Start")()

	g.db = &sampleDatabase{}
	g.context, g.cancel = context.WithCancel(context.Background())
	g.wg = &sync.WaitGroup{}

	g.wg.Add(2)
	go g.routineA()
	go g.routineB()

	return nil
}

func (g *goRoutines) Stop(s service.Service) error {
	defer g.logEntryExit("Stop")()
	g.shutdown()
	return nil
}

func (g *goRoutines) shutdown() {

	g.cancel()
	g.wg.Wait()
}

func (g *goRoutines) logEntryExit(name string) func() {
	g.Log().Infof("entry %s", name)
	return func() {
		g.Log().Infof("exit  %s", name)
	}
}

func (g *goRoutines) routineA() {

	defer g.logEntryExit("routineA")()
	defer g.wg.Done()

	finished := false
	ticker := time.NewTicker(5 * time.Second)
	for !finished {
		select {
		case <-ticker.C:
			g.Log().Info("routineA running")
		case <-g.context.Done():
			finished = true
		}
	}
}

func (g *goRoutines) routineB() {

	defer g.logEntryExit("routineB")()
	defer g.wg.Done()

	time.Sleep(1 * time.Second)

	// Start would be a better place for db connections but this example
	// shows how to cancel and exit from within a goroutine
	err := g.db.Connect()
	if err != nil {

		g.Log().Error("unable to connect to database")

		// de-increment wait group or shutdown() will block forever
		g.wg.Done()
		g.shutdown()

		os.Exit(1)
	}

	<-g.context.Done()
	g.db.Shutdown()
}

// demo object for example lifecycle management of a DB
type sampleDatabase struct{}

func (s *sampleDatabase) Connect() error {
	_, failConnect := os.LookupEnv("FAIL_CONNECT")
	if failConnect {
		return fmt.Errorf("simulate db connect failure")
	}
	return nil
}
func (s *sampleDatabase) Shutdown() error { return nil }
