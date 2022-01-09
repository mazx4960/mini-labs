# mini-cache
[![](https://img.shields.io/badge/Language-Go-E5A505?style=flat-square)]()

Cache system implemented using the O(1) LFU eviction scheme

## Description  

A cache system implemented using Go. See [References](#references) for the research paper behind this algorithm. 

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

- https://www.researchgate.net/publication/355493987_An_O1_algorithm_for_implementing_the_LFU_cache_eviction_scheme