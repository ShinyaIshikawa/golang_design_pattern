package singleton

import (
	"sync"
)

// singleton instance.
var singleton *Singleton

//  mutex.
var mux sync.Mutex

// Singleton is an instance.
type Singleton struct {
	s string
}

// GetInstance return singleton incetance.
// use Lcok and Unlock function for mutual exclusion.
func GetInstance() *Singleton {
	// Lock so only one goroutine at a time can access the singleton.
	mux.Lock()
	if singleton == nil {
		//		time.Sleep(2 * time.Second)
		singleton = newSinleton()
	}
	defer mux.Unlock()
	//fmt.Println(&singleton)
	return singleton
}

func newSinleton() *Singleton {
	//fmt.Println("create instance.")
	return &Singleton{"something"}
}
