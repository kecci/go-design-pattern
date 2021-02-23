package usecase

type EvictionAlgo interface {
    Evict(c *Cache)
}