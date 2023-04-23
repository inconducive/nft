package main

import (
	"log"
	"sync"
	"time"

	api "github.com/inconducive/nft/api"
	"golang.org/x/time/rate"
)

var (
	Caresults = []string{}
)

func main() {
	cr, newline := "\n", "\n"
	limiter := rate.NewLimiter(rate.Every(1*time.Second), 3)
	log.Println("limiter is \n", limiter)
	for i := 0; i < len(api.CAs); i++ {
		// runtime.GOMAXPROCS(5)
		// go api.GetFloorPrice(api.APIkey, api.CAs[i])
		// for j := 1; j < rateLimit; j++ {
		Caresults = append(Caresults, api.GetFloorPrice(api.APIkey, api.CAs[i]))
		log.Println("Fetching for: Contract Address #", i+1, " <> ", api.CAs[i], "\n", api.GetFloorPrice(api.APIkey, api.CAs[i]))

		if (i+1)%(api.RateLimit-1) == 0 {
			log.Println(newline, "sleeping for 1 Second... due to the 4 Requests/ Second rate limit you imposed", cr, newline)
			time.Sleep(time.Second * 1)
		}
	}
}

type Limiter interface {
	Wait()
}

type limiter struct {
	tick    time.Duration
	count   uint
	entries []time.Time
	index   uint
	mutex   sync.Mutex
}

func NewLimiter(tick time.Duration, count uint) Limiter {
	l := limiter{
		tick:  tick,
		count: count,
		index: 0,
	}
	l.entries = make([]time.Time, count)
	before := time.Now().Add(-2 * tick)
	for i, _ := range l.entries {
		l.entries[i] = before
	}
	return &l
}
func (l *limiter) Wait() {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	last := &l.entries[l.index]
	next := last.Add(l.tick)
	now := time.Now()
	if now.Before(next) {
		time.Sleep(next.Sub(now))
	}
	*last = time.Now()
	l.index = l.index + 1
	if l.index == l.count {
		l.index = 0
	}
}
