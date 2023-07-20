package asyncjob

import (
	"context"
	"go-api/common"
	"log"
	"sync"
)

type group struct {
	isConcurrent bool
	jobs         []Job
	wg           *sync.WaitGroup
}

func NewGroup(isConcurrent bool, jobs ...Job) *group {
	return &group{isConcurrent: isConcurrent, jobs: jobs, wg: new(sync.WaitGroup)}
}
func (g *group) Run(ctx context.Context) error {
	g.wg.Add(len(g.jobs))
	errChan := make(chan error, len(g.jobs))
	for i, _ := range g.jobs {
		if g.isConcurrent {
			go func(job Job) {
				defer common.AppRecover()
				errChan <- g.runJob(ctx, job)
				g.wg.Done()
			}(g.jobs[i])
			continue
		}
		job := g.jobs[i]
		err := g.runJob(ctx, job)
		if err != nil {
			if !g.isConcurrent {
				return err
			}
			log.Println(err)
		}
		errChan <- err
		g.wg.Done()
	}
	g.wg.Wait()
	var err error
	for i := 0; i < len(g.jobs); i++ {
		if v := <-errChan; err != nil {
			if !g.isConcurrent {
				return v
			}
			err = v
		}
	}

	return err
}
func (g *group) runJob(ctx context.Context, job Job) error {
	if err := job.Execute(ctx); err != nil {
		for {
			log.Println(err)
			if job.State() == StateRetryFailed {
				return err
			}
			if job.Retry(ctx) == nil {
				return nil
			}
		}
	}
	return nil
}
