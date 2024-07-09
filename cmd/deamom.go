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

// конструктор демона
func New(period time.Duration, task func(fc *en.FileCache)) *PeriodicTask {
	return &PeriodicTask{
		period: period,
		task:   task,
	}
}

func (pt *PeriodicTask) RunMainTask(ctx context.Context, fc *en.FileCache) {
	ticker := time.NewTicker(pt.period)
	defer ticker.Stop()

	for { // цикл крутится до прерывания
		select {
		case <-ticker.C: //
			pt.task(fc)

		case <-ctx.Done(): //grace full shutdown

			pt.task(fc) //запись (попытка) остатков из КЭШ в файл

			log.Println("Stopping periodic task:", ctx.Err())
			return
		}
	}
}
