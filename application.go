package main

import (
	"context"
	"fmt"
	"os"
	"time"
)

type ApplicationConfig struct {
}

type signalHandler func(os.Signal) (bool, int)

type application struct {
	ctx    context.Context
	cancel context.CancelFunc
}

func NewApplication(ctx context.Context) *application {

	ctx, cancel := context.WithCancel(ctx)

	return &application{
		ctx:    ctx,
		cancel: cancel,
	}
}

func (a *application) grathefullShoutdown(ctx context.Context) {
}

func (a *application) cleanup(s os.Signal) (bool, int) {

	ctx, cancel := context.WithTimeout(a.ctx, time.Duration(30000)*time.Millisecond)
	defer cancel()

	a.grathefullShoutdown(ctx)

	return true, 0
}

func (a *application) terminate(s os.Signal) (bool, int) {
	fmt.Println("Use TERM signal for grecfull shoutdown. Exiting now.")
	a.cancel()

	return true, 0
}

func (a *application) reloadConfig(s os.Signal) (bool, int) {
	return false, 0
}
