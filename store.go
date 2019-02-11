package trigram

import "math/rand"

type Store interface {
	Add(Trigram)
	GetByPrefix(prefix [2]string) []Trigram
	Seed() [2]string
}

type MapStore struct {
	MaxResultLength int
	trigrams        map[string][]Trigram
}

func NewMapStore(max int) MapStore {
	return MapStore{
		MaxResultLength: max,
		trigrams:        make(map[string][]Trigram),
	}
}

func (s MapStore) Add(tg Trigram) {
	key := tg.first + tg.second
	if slice, ok := s.trigrams[key]; !ok {
		s.trigrams[key] = []Trigram{tg}
	} else {
		s.trigrams[key] = append(slice, tg)
	}
}

func (s MapStore) GetByPrefix(prefix [2]string) []Trigram {
	out := make([]Trigram, 0)
	key := prefix[0] + prefix[1]
	return s.follow(key, out)
}

func (s MapStore) follow(key string, out []Trigram) []Trigram {
	if s.MaxResultLength != -1 && len(out) >= s.MaxResultLength {
		return out
	}
	trigrams, ok := s.trigrams[key]
	if !ok {
		return out
	}

	trigram := trigrams[rand.Intn(len(trigrams))]

	out = append(out, trigram)
	key = trigram.second + trigram.third
	return s.follow(key, out)
}

func (s MapStore) Seed() [2]string {
	return [2]string{}
}

type Trigram struct {
	first, second, third string
}

func NewTrigram(first, second, third string) Trigram {
	return Trigram{first, second, third}
}
