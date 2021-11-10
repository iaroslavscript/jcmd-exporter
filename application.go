package main

import (
	"context"
	"fmt"
	"os"
	"time"
)

type Application struct {
	ctx    context.Context
	cancel context.CancelFunc
}

func NewApplication(ctx context.Context) *Application {

	ctx, cancel := context.WithCancel(ctx)

	return &Application{
		ctx:    ctx,
		cancel: cancel,
	}
}

func (a *Application) grathefullShoutdown(ctx context.Context) {
}

func (a *Application) cleanup(s os.Signal) (bool, int) {

	ctx, cancel := context.WithTimeout(a.ctx, time.Duration(30000)*time.Millisecond)
	defer cancel()

	a.grathefullShoutdown(ctx)

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
