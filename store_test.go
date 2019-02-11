package trigram_test

import (
	"testing"

	"github.com/cheekybits/is"

	"github.com/burrbd/trigram"
)

func TestMapStoreGetByPrefix(t *testing.T) {
	is := is.New(t)
	store := trigram.NewMapStore(-1)
	exp := trigram.NewTrigram("a", "b", "c")
	store.Add(exp)

	result := store.GetByPrefix([2]string{"a", "b"})

	is.Equal(1, len(result))
	act := result[0]
	is.Equal(exp, act)
}

func TestMapStoreGetByPrefixMultiple(t *testing.T) {
	is := is.New(t)

	store := trigram.NewMapStore(-1)
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

func TestMapStoreGetByPrefixCircular(t *testing.T) {
	is := is.New(t)

	store := trigram.NewMapStore(4)
	store.Add(trigram.NewTrigram("a", "b", "c"))
	store.Add(trigram.NewTrigram("b", "c", "a"))
	store.Add(trigram.NewTrigram("c", "a", "b"))

	act := store.GetByPrefix([2]string{"a", "b"})
	exp := []trigram.Trigram{
		trigram.NewTrigram("a", "b", "c"),
		trigram.NewTrigram("b", "c", "a"),
		trigram.NewTrigram("c", "a", "b"),
		trigram.NewTrigram("a", "b", "c"),
	}

	is.Equal(exp, act)
}

func TestMapStoreGetByPrefixWithWeightedOptions(t *testing.T) {
	is := is.New(t)

	store := trigram.NewMapStore(-1)
	store.Add(trigram.NewTrigram("a", "b", "c"))
	store.Add(trigram.NewTrigram("b", "c", "foo"))
	store.Add(trigram.NewTrigram("b", "c", "bar"))

	found := 0
	for {
		result := store.GetByPrefix([2]string{"a", "b"})
		if result[1] == trigram.NewTrigram("b", "c", "foo") {
			found++
			break
		}
	}
	for {
		result := store.GetByPrefix([2]string{"a", "b"})
		if result[1] == trigram.NewTrigram("b", "c", "bar") {
			found++
			break
		}
	}

	is.Equal(2, found)
}

func TestMapStoreGetByPrefixWithProportionateRandomness(t *testing.T) {
	is := is.New(t)

	store := trigram.NewMapStore(-1)
	store.Add(trigram.NewTrigram("a", "b", "c"))
	store.Add(trigram.NewTrigram("b", "c", "foo"))
	for i := 0; i <= 3; i++ {
		store.Add(trigram.NewTrigram("b", "c", "bar"))
	}
	contest := map[string]int{"foo": 0, "bar": 0}

	for i := 0; i <= 10000; i++ {
		result := store.GetByPrefix([2]string{"a", "b"})[1]
		contest[result.Third]++
	}

	is.Equal(3, contest["bar"]/contest["foo"])
}

func TestMapStoreSeed(t *testing.T) {
	is := is.New(t)

	store := trigram.NewMapStore(-1)
	store.Add(trigram.NewTrigram("a", "b", "c"))
	seed := store.Seed()

	is.Equal([2]string{"a", "b"}, seed)
}
