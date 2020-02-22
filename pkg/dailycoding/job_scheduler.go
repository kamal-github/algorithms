package dailycoding

import "time"

// Use rather Priority Queue and pop the job when its time is hit.

// alternative solution - not sure if it works :P
type Functions []CallFunc

type CallFunc struct {
	f            func()
	callAfter    int
	msLeftToCall int
}

//f1, 100
//f2, 100
//f3, 120
//f4 90
func (fs *Functions) JobScheduler(f func(), n int) {
	*fs = append(*fs, CallFunc{
		f:            f,
		callAfter:    n,
		msLeftToCall: n,
	})
}

func (fs *Functions) RunJob(ch <-chan time.Time) {
	for range ch {
		for _, cf := range *fs {
			if cf.msLeftToCall == 0 {
				cf.f()
				cf.msLeftToCall = cf.callAfter
				continue
			}

			cf.msLeftToCall--
		}
	}
}
