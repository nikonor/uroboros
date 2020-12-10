package uroboros

import (
	"strconv"
	"strings"
	"sync"
	"time"
)

type Uroboros struct {
	sync.Mutex
	col    int
	period int64
	head   int
	u      []int64
}

func (u *Uroboros) String() string {
	var s []string
	for i := 0; i < len(u.u); i++ {
		s = append(s, strconv.FormatInt(u.u[i]/1000000, 10))
	}
	return "head=" + strconv.Itoa(u.head) + ", data=[" + strings.Join(s, ",") + "]"
}

func New(col int, period time.Duration) *Uroboros {
	return &Uroboros{
		col:    col,
		period: period.Nanoseconds() / 1000,
		u:      make([]int64, col),
	}
}

func (u *Uroboros) Can(t time.Time) bool {
	tt := t.UnixNano() / 1000

	u.Lock()
	defer u.Unlock()

	if u.u[u.head] == 0 || u.u[u.head]+u.period <= tt {
		u.u[u.head] = tt
	} else {
		return false
	}

	u.setNext()
	return true
}

func (u *Uroboros) setNext() {
	u.head++
	if u.head == u.col {
		u.head = 0
	}
}
