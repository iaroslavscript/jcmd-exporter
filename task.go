package main

import (
	"context"
	"fmt"
	"os/exec"
	"regexp"
	"time"
)

func RunTasks(ctx context.Context, p *regexp.Regexp, tasks []*JcmdTask) {

	for i := range tasks {
		go func(task *JcmdTask) error {

			ticker := time.NewTicker(500 * time.Millisecond)
			defer ticker.Stop()

			for {
				select {
				case <-ctx.Done():
					return ctx.Err() // TODO do we need to return error in goroutine?
				case t := <-ticker.C:
					fmt.Println("Tick at", t)
					output, err := CallJcmd(
						ctx,
						time.Duration(task.TimeoutMs)*time.Millisecond,
						task.PathJcmd,
						task.MainClass,
						task.SubSystem,
					)

					if err != nil {
						fmt.Println("error", err)
						continue
					}
					parse_response(output, p, task.Metrics)

					fmt.Println("Finished", t)
				}
			}
		}(tasks[i])
	}
}

func CallJcmd(ctx context.Context, timeout time.Duration, app string, mainClass string, arg1 string) (string, error) {

	// TODO do we need "select { case <-ctx.Done()" here ???

	ctx, close := context.WithTimeout(ctx, timeout)
	defer close()

	cmd := exec.CommandContext(ctx, app, mainClass, arg1)
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	return string(stdout), nil
}
