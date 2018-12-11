package singleton

// Singleton
type Singleton struct {
	own interface{}
}

func getInstance() *Singleton {
	return &Singleton{}
}
