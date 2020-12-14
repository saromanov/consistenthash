package consistenthash

import (
	"fmt"
	"sync"
	"crypto/sha1"
)
type Hasher interface {
	Sum([]byte) []byte
}

type Consistent struct {
	h Hasher
	mu *sync.RWMutex
}

func New(hasher Hasher)*Consistent {
	if hasher == nil {
		hasher = sha1.New()
	}
	return &Consistent {
		h: hasher,
		mu: &sync.RWMutex{},
	}
}

// Add provides adding of the new member
func (c *Consistent) Add(m string) error {
	if m == "" {
		return fmt.Errorf("node name is not defined")
	}

	return nil

}