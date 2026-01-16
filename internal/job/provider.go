package job

import "github.com/google/wire"

func NewJobGroupProvider() []IJobGroup {
	jg := []IJobGroup{
		NewJobGroup("default"),
	}

	// append other
	bizJobGroup := NewJobGroup("biz")

	bizJobGroup.AddJobs()

	return jg
}

var ProviderSet = wire.NewSet(NewJobGroupProvider)
