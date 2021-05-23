package main

import (
	"sync"
)

// 数据结构
type URLStore struct {
	urls map[string]string // 映射短网址
	mu   sync.RWMutex      // 读写锁，确保线程安全
}

// 工厂函数
func NewURLStore() *URLStore {
	return &URLStore{urls: make(map[string]string)}
}

// 读
func (s *URLStore) Get(key string) string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.urls[key]
}

// 写
func (s *URLStore) Set(key, url string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, present := s.urls[key]; present {
		return false
	}
	s.urls[key] = url
	return true
}

func (s *URLStore) Count() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.urls)
}

func (s *URLStore) Put(url string) string {
	for {
		key := genKey(s.Count()) // generate the short URL
		if ok := s.Set(key, url); ok {
			return key
		}
	}
	// shouldn't get here
	return ""
}
