package main

type Topic struct {
	Title     string
	Questions []Question
}

type Question struct {
	Statement string
	Answer    string
}

func (t *Topic) AddQuestion(question Question) {
	t.Questions = append(t.Questions, question)
}
