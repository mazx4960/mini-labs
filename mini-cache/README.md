# mini-cache
[![](https://img.shields.io/badge/Category-Caching-E5A505?style=flat-square)]() [![](https://img.shields.io/badge/Category-Distributed%20Systems-E5A505?style=flat-square)]() [![](https://img.shields.io/badge/Language-Go-E5A505?style=flat-square)]()

Cache system implemented using the O(1) LFU eviction scheme

## Description  

A cache system implemented using Go. See [References](#references) for the research paper behind this algorithm. 

Inspired by [memcached](https://github.com/memcached/memcached) and [redis](https://github.com/redis/redis). 

## Features

- LFU eviction scheme
- Break tie by access time for same frequencies (WIP)

## Usage

- Create a new cache with a capacity of 10 items: `lfu.New(10)`
- Add items to the cache: `cache.Insert("key1", "value1")`
- Get items from the cache: `cache.Get("key1")`
- Update items in the cache: `cache.Update("key1", "new value")`
- Delete items from the cache: `cache.Delete("key1")`

## References

- [An O(1) algorithm for implementing the LFUcache eviction scheme](https://www.researchgate.net/publication/355493987_An_O1_algorithm_for_implementing_the_LFU_cache_eviction_scheme)

## Similar resources

- [Anti-Caching: A New Approach to
Database Management System Architecture](https://www.vldb.org/pvldb/vol6/p1942-debrabant.pdf)
- [A Program Optimization for
Automatic Database Result Caching](http://adam.chlipala.net/papers/SqlcachePOPL17/SqlcachePOPL17.pdf)