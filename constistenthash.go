package consistenthash

import (
	"fmt"
	"sync"
	"crypto/sha1"
)
type Hasher interface {
	Sum([]byte) []byte
	Write([]byte) (int, error)
}

type Consistent struct {
	h Hasher
	mu *sync.RWMutex
}

// New provides implementation of initializaer
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
	
	if _, err := c.h.Write([]byte(m)); err != nil {
		return err
	}
	fmt.Println(c.h.Sum(nil))
	return nil

}

func (c *Consistent) Remove(m string) error {
	if m == "" {
		return fmt.Errorf("not name is not defined")
	}
	return nil
}
