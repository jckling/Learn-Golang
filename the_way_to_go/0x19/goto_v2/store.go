package main

import (
	"encoding/gob"
	"io"
	"log"
	"os"
	"sync"
)

type URLStore struct {
	urls map[string]string
	mu   sync.RWMutex
	file *os.File // 持久化存储
}

type record struct {
	Key, URL string // 短网址（键）：长网址
}

func NewURLStore(filename string) *URLStore {
	s := &URLStore{urls: make(map[string]string)} // 打开文件
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal("Error opening URLStore:", err)
	}
	s.file = f
	if err := s.load(); err != nil { // 读取数据
		log.Println("Error loading URLStore:", err)
	}
	return s
}

func (s *URLStore) Get(key string) string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.urls[key]
}

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
		key := genKey(s.Count())
		if ok := s.Set(key, url); ok {
			if err := s.save(key, url); err != nil { // 保存数据
				log.Println("Error saving to URLStore:", err)
			}
			return key
		}
	}
	panic("shouldn't get here")
}

// 加载数据
func (s *URLStore) load() error {
	if _, err := s.file.Seek(0, 0); err != nil {
		return err
	}
	d := gob.NewDecoder(s.file)
	var err error
	for err == nil {
		var r record
		if err = d.Decode(&r); err == nil { // 解码
			s.Set(r.Key, r.URL)
		}
	}
	if err == io.EOF {
		return nil
	}
	return err
}

// 保存数据
func (s *URLStore) save(key, url string) error {
	e := gob.NewEncoder(s.file)
	return e.Encode(record{key, url}) // 编码
}
