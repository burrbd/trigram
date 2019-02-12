package trigram_test

import (
	"testing"

	"github.com/cheekybits/is"

	"github.com/burrbd/trigram"
)

type mockStore struct {
	AddFunc         func(trigram.Trigram)
	GetByPrefixFunc func([2]string) []trigram.Trigram
	SeedFunc        func() [2]string
}

func (s mockStore) Add(tg trigram.Trigram) {
	s.AddFunc(tg)
}

func (s mockStore) GetByPrefix(prefix [2]string) []trigram.Trigram {
	return s.GetByPrefixFunc(prefix)
}

func (s mockStore) Seed() [2]string {
	return s.SeedFunc()
}

func TestNauturalLanguageGeneratorLearn(t *testing.T) {
	is := is.New(t)

	count := 0
	store := mockStore{AddFunc: func(tg trigram.Trigram) {
		switch count {
		case 0:
			is.Equal(trigram.NewTrigram("to", "be", "or"), tg)
		case 1:
			is.Equal(trigram.NewTrigram("be", "or", "not"), tg)
		case 2:
			is.Equal(trigram.NewTrigram("or", "not", "to"), tg)
		default:
			is.Failf("unexpected trigram: %+v", tg)
		}
		count++
	}}

	trigram.NewLanguageGenerator(store).
		Learn([]string{"to", "be", "or", "not", "to"})

	is.Equal(3, count)
}

func TestNaturalLanguageGeneratorGenerate(t *testing.T) {
	is := is.New(t)

	getByPrefixCalled, seedCalled := false, false
	store := mockStore{
		GetByPrefixFunc: func(_ [2]string) []trigram.Trigram {
			getByPrefixCalled = true
			return []trigram.Trigram{
				trigram.NewTrigram("two", "three", "four"),
				trigram.NewTrigram("three", "four", "five"),
				trigram.NewTrigram("four", "five", "six"),
			}
		},
		SeedFunc: func() [2]string {
			seedCalled = true
			return [2]string{"two", "three"}
		}}

	act := trigram.NewLanguageGenerator(store).Generate()
	exp := "three four five six"

	is.Equal(exp, act)
	is.True(getByPrefixCalled)
	is.True(seedCalled)
}
