package trigram

import (
	"math/rand"
	"sync"
)

type Store interface {
	Add(Trigram)
	GetByPrefix(prefix [2]string) []Trigram
	Seed() [2]string
}

type MapStore struct {
	QueryLimit int
	sync.RWMutex
	trigrams map[[2]string][]string
}

func NewMapStore(limit int) *MapStore {
	return &MapStore{
		QueryLimit: limit,
		trigrams:   make(map[[2]string][]string),
	}
}

func (s *MapStore) Add(tg Trigram) {
	s.Lock()
	defer s.Unlock()
	key := [2]string{tg.First, tg.Second}
	if slice, ok := s.trigrams[key]; !ok {
		s.trigrams[key] = []string{tg.Third}
	} else {
		s.trigrams[key] = append(slice, tg.Third)
	}
}

func (s *MapStore) GetByPrefix(prefix [2]string) []Trigram {
	s.RLock()
	defer s.RUnlock()
	out := make([]Trigram, 0)
	return s.follow(prefix, out)
}

func (s MapStore) follow(prefix [2]string, out []Trigram) []Trigram {
	if s.QueryLimit != -1 && len(out) >= s.QueryLimit {
		return out
	}
	words, ok := s.trigrams[prefix]
	if !ok {
		return out
	}

	word := words[rand.Intn(len(words))]
	trigram := Trigram{First: prefix[0], Second: prefix[1], Third: word}
	out = append(out, trigram)
	prefix = [2]string{trigram.Second, trigram.Third}
	return s.follow(prefix, out)
}

func (s *MapStore) Seed() [2]string {
	s.RLock()
	defer s.RUnlock()
	var key [2]string
	for key = range s.trigrams {
		break
	}
	return key
}

type Trigram struct {
	First, Second, Third string
}

func NewTrigram(first, second, third string) Trigram {
	return Trigram{first, second, third}
}
