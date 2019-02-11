package trigram

type Store interface {
	Add(Trigram)
	GetByPrefix(prefix [2]string) []Trigram
	Seed() [2]string
}

type MapStore map[string][]Trigram

func (s MapStore) Add(tg Trigram) {
	key := tg.first + tg.second
	if slice, ok := s[key]; !ok {
		s[key] = []Trigram{tg}
	} else {
		s[key] = append(slice, tg)
	}
}

func (s MapStore) GetByPrefix(prefix [2]string) []Trigram {
	out := make([]Trigram, 0)
	key := prefix[0] + prefix[1]
	return s.follow(key, out)
}

func (s MapStore) follow(key string, out []Trigram) []Trigram {
	trigrams, ok := s[key]
	if !ok {
		return out
	}
	out = append(out, trigrams[0])
	key = trigrams[0].second + trigrams[0].third
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
