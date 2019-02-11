package trigram_test

import (
	"testing"

	"github.com/cheekybits/is"

	"github.com/burrbd/trigram"
)

func TestMapStoreGetByPrefix(t *testing.T) {
	is := is.New(t)
	store := trigram.MapStore{}
	exp := trigram.NewTrigram("a", "b", "c")
	store.Add(exp)

	result := store.GetByPrefix([2]string{"a", "b"})

	is.Equal(1, len(result))
	act := result[0]
	is.Equal(exp, act)
}

func TestMapStoreGetByPrefixMultiple(t *testing.T) {
	is := is.New(t)

	store := trigram.MapStore{}
	store.Add(trigram.NewTrigram("a", "b", "c"))
	store.Add(trigram.NewTrigram("b", "c", "d"))
	store.Add(trigram.NewTrigram("c", "d", "e"))

	act := store.GetByPrefix([2]string{"a", "b"})
	exp := []trigram.Trigram{
		trigram.NewTrigram("a", "b", "c"),
		trigram.NewTrigram("b", "c", "d"),
		trigram.NewTrigram("c", "d", "e"),
	}

	is.Equal(exp, act)
}
