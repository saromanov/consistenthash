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
	members map[string]string
}

// New provides implementation of initializaer
func New(hasher Hasher)*Consistent {
	if hasher == nil {
		hasher = sha1.New()
	}
	return &Consistent {
		h: hasher,
		mu: &sync.RWMutex{},
		members: map[string]string{},
	}
}

// AddMember provides adding of the new member
func (c *Consistent) AddMember(m string) error {
	if m == "" {
		return fmt.Errorf("node name is not defined")
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	
	if _, ok := c.members[m]; ok {
		return fmt.Errorf("member already added")
	}
	return nil

}

func (c *Consistent) Remove(m string) error {
	if m == "" {
		return fmt.Errorf("not name is not defined")
	}
	return nil
}
