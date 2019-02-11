package trigram_test

import (
	"testing"

	"github.com/cheekybits/is"

	"github.com/burrbd/trigram"
)

type mockStore struct {
	AddFunc func(trigram.Trigram)
}

func (s mockStore) Add(tg trigram.Trigram) {
	s.AddFunc(tg)
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

	trigram.NewNaturalLanguageGenerator(store).
		Learn([]string{"to", "be", "or", "not", "to"})

	is.Equal(3, count)
}
