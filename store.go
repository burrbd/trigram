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
	key := tg.First + tg.Second
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
	key = trigram.Second + trigram.Third
	return s.follow(key, out)
}

func (s MapStore) Seed() [2]string {
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
