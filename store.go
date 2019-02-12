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
	trigrams map[string][]Trigram
}

func NewMapStore(limit int) *MapStore {
	return &MapStore{
		QueryLimit: limit,
		trigrams:   make(map[string][]Trigram),
	}
}

func (s *MapStore) Add(tg Trigram) {
	s.Lock()
	defer s.Unlock()
	key := tg.First + tg.Second
	if slice, ok := s.trigrams[key]; !ok {
		s.trigrams[key] = []Trigram{tg}
	} else {
		s.trigrams[key] = append(slice, tg)
	}
}

func (s *MapStore) GetByPrefix(prefix [2]string) []Trigram {
	s.RLock()
	defer s.RUnlock()
	out := make([]Trigram, 0)
	key := prefix[0] + prefix[1]
	return s.follow(key, out)
}

func (s MapStore) follow(key string, out []Trigram) []Trigram {
	if s.QueryLimit != -1 && len(out) >= s.QueryLimit {
		return out
	}
	trigrams, ok := s.trigrams[key]
	if !ok {
		return out
	}

	trigram := trigrams[rand.Intn(len(trigrams))]

	out = append(out, trigram)
	key = trigram.Second + trigram.Third
	return s.follow(key, out)
}

func (s *MapStore) Seed() [2]string {
	s.RLock()
	defer s.RUnlock()
	i := rand.Intn(len(s.trigrams))
	var key string
	for key = range s.trigrams {
		if i == 0 {
			break
		}
		i--
	}
	return [2]string{
		s.trigrams[key][0].First,
		s.trigrams[key][0].Second}
}

type Trigram struct {
	First, Second, Third string
}

func NewTrigram(first, second, third string) Trigram {
	return Trigram{first, second, third}
}
