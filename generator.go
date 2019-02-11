package trigram

import "strings"

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

func (g *NaturalLanguageGenerator) Generate() string {
	out := make([]string, 0)
	trigrams := g.store.GetByPrefix(g.store.Seed())
	n := len(trigrams)
	for i := 0; i < n; i++ {
		if i == n-1 {
			out = append(out, trigrams[i].second, trigrams[i].third)
		} else {
			out = append(out, trigrams[i].second)
		}
	}

	return strings.Join(out, " ")
}
