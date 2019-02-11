package trigram

type Store interface {
	Add(Trigram)
}

type Learner interface {
	Learn([]string)
}

type LanguageGenerator interface {
	Generate() string
}

type NaturalLanguageGenerator struct {
	store Store
}

func NewNaturalLanguageGenerator(store Store) *NaturalLanguageGenerator {
	return &NaturalLanguageGenerator{store}
}

func (g *NaturalLanguageGenerator) Learn(words []string) {
	n := len(words)
	for k := range words {
		if k+3 > n {
			break
		}
		g.store.Add(NewTrigram(words[k], words[k+1], words[k+2]))
	}
}

type Trigram struct {
	first, second, third string
}

func NewTrigram(first, second, third string) Trigram {
	return Trigram{first, second, third}
}
