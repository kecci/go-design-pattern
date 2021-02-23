package usecase

import "fmt"

type Fifo struct {
}

func (l *Fifo) Evict(c *Cache) {
    fmt.Println("Evicting by fifo strtegy")
}