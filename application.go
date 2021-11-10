package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"
)

type Application struct {
	ctx        context.Context
	cancel     context.CancelFunc
	srv        http.Server
	inShutdown bool
}

func NewApplication(ctx context.Context) *Application {

	ctx, cancel := context.WithCancel(ctx)

	return &Application{
		ctx:        ctx,
		cancel:     cancel,
		inShutdown: false,
	}
}

func (a *Application) grathefullShutdown(ctx context.Context) error {
	fmt.Println("enter grathefull shutdown state")
	a.inShutdown = true
	err := a.srv.Shutdown(ctx)

	fmt.Println("exit grathefull shutdown state with error", err)
	return err
}

func (a *Application) cleanup(s os.Signal) (bool, int) {

	ctx, cancel := context.WithTimeout(a.ctx, time.Duration(30000)*time.Millisecond)
	defer cancel()

	a.grathefullShutdown(ctx)

	return true, 0
}

func (a *Application) terminate(s os.Signal) (bool, int) {
	fmt.Println("Use TERM signal for grecfull shoutdown. Exiting now.")
	a.cancel()

	return true, 0
}

func (a *Application) reloadConfig(s os.Signal) (bool, int) {
	return false, 0
}
