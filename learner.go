package trigram

type Learner interface {
	Learn([]string)
}
