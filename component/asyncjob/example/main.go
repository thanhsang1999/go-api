package main

import (
	"context"
	"errors"
	"go-api/component/asyncjob"
	"log"
	"time"
)

func main() {
	job1 := asyncjob.NewJob(func(ctx context.Context) error {
		time.Sleep(1 * time.Second)
		log.Println("I am job 1")
		//return nil
		return errors.New("something went wrong job 1")
	})
	job2 := asyncjob.NewJob(func(ctx context.Context) error {
		time.Sleep(2 * time.Second)
		log.Println("I am job 2")
		return nil
		//return errors.New("something went wrong job 1")
	})
	job3 := asyncjob.NewJob(func(ctx context.Context) error {
		time.Sleep(12 * time.Second)
		log.Println("I am job 3")
		return nil
		//return errors.New("something went wrong job 1")
	})
	group := asyncjob.NewGroup(true, job1, job2, job3)
	if err := group.Run(context.Background()); err != nil {
		log.Println(err)
	}
}
