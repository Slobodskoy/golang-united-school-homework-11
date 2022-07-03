package batch

import (
	"sync"
	"time"
)

type user struct {
	ID int64
}

func getOne(id int64) user {
	time.Sleep(time.Millisecond * 100)
	return user{ID: id}
}

func getBatch(n int64, pool int64) (res []user) {
	wg := sync.WaitGroup{}
	sem := make(chan struct{}, pool)
	res = make([]user, n)
	for i := int64(0); i < n; i++ {
		wg.Add(1)
		go func(i int64) {
			sem <- struct{}{}
			res[i] = getOne(i)
			<-sem
			wg.Done()
		}(i)
	}
	wg.Wait()
	return
}
