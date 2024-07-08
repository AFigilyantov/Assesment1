package main

import (
	"context"
	"log"
	"time"

	en "asssement1.ru/entities"
)

// PeriodicTask represents a struct which handles the periodic execution of a function
type PeriodicTask struct {
	period time.Duration
	task   func(fc *en.FileCache)
}

func New(period time.Duration, task func(fc *en.FileCache)) *PeriodicTask {
	return &PeriodicTask{
		period: period,
		task:   task,
	}
}

func (pt *PeriodicTask) RunMainTask(ctx context.Context, fc *en.FileCache) {
	ticker := time.NewTicker(pt.period)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			pt.task(fc)

		case <-ctx.Done():

			pt.task(fc)

			log.Println("Stopping periodic task:", ctx.Err())
			return
		}
	}
}
