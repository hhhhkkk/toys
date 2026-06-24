package job

import (
	"context"
	"fmt"
	"sync"
)

type IJobGroup interface {
	AddJobs(jobs ...IJob)
	Run(ctx context.Context) error
}

// IJob 定义 Job 接口
type IJob interface {
	// Run 执行任务，context 用于优雅关闭
	Run(ctx context.Context) error
	// Name 返回任务名称
	Name() string
}

type JobGroup struct {
	jobs []IJob
	name string
	lock sync.RWMutex
}

var _ IJobGroup = (*JobGroup)(nil)

func NewJobGroup(name string) *JobGroup {
	return &JobGroup{
		name: name,
		jobs: make([]IJob, 0),
	}
}

func (jg *JobGroup) AddJobs(jobs ...IJob) {
	jg.lock.Lock()
	defer jg.lock.Unlock()
	jg.jobs = append(jg.jobs, jobs...)
}

func (jg *JobGroup) Name() string {
	return jg.name
}

func (jg *JobGroup) Run(ctx context.Context) error {
	var wg sync.WaitGroup
	defer wg.Wait()

	for _, job := range jg.jobs {
		wg.Add(1)
		go func(job IJob) {
			defer wg.Done()
			err := job.Run(ctx)
			if err != nil {
				fmt.Printf("job [%s] failed: %v\n", job.Name(), err)
			}
		}(job)
	}
	fmt.Printf("All [%s] jobs started", jg.Name())
	<-ctx.Done()
	// 等待所有 job 完成
	fmt.Printf("All [%s] jobs completed\n", jg.Name())
	return nil
}
