package main

import "github.com/kecci/design-pattern-go/behavioral_patterns/strategy/usecase"

func main() {
	lfu := &usecase.Lfu{}
	cache := usecase.InitCache(lfu)

	cache.Add("a", "1")
	cache.Add("b", "2")

	cache.Add("c", "3")

	lru := &usecase.Lru{}
	cache.SetEvictionAlgo(lru)

	cache.Add("d", "4")

	fifo := &usecase.Fifo{}
	cache.SetEvictionAlgo(fifo)

	cache.Add("e", "5")

}
