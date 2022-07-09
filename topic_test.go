package main

import "testing"

func Contains(needle Question, haystack []Question) bool {
	for _, question := range haystack {
		if needle == question {
			return true
		}
	}
	return false
}

func TestTopic(t *testing.T) {
	t.Run("next question", func(t *testing.T) {
		topic := NewTopic()
		topic.Title = "A topic"
		topic.Questions = []Question{
			{Statement: "Q1", Answer: "A1"},
			{Statement: "Q2", Answer: "A2"},
			{Statement: "Q3", Answer: "A3"},
			{Statement: "Q4", Answer: "A4"},
		}

		seen := []Question{}
		for len(seen) < len(topic.Questions) {
			if next := topic.NextQuestion(); next != nil {
				if Contains(*next, seen) {
					t.Error("Should not get the same question twice")
				}
				seen = append(seen, *next)
			} else {
				t.Error("Expected question, got nil")
			}
		}
	})

	t.Run("previous question", func(t *testing.T) {
		topic := NewTopic()
		topic.Title = "A topic"
		topic.Questions = []Question{
			{Statement: "Q1", Answer: "A1"},
			{Statement: "Q2", Answer: "A2"},
			{Statement: "Q3", Answer: "A3"},
			{Statement: "Q4", Answer: "A4"},
		}

		if prev := topic.PrevQuestion(); prev != nil {
			t.Errorf("Expected nil, got %v", prev)
		}

		questions := []*Question{
			topic.NextQuestion(),
			topic.NextQuestion(),
			topic.NextQuestion(),
			topic.NextQuestion(),
		}

		for i := len(topic.Questions) - 2; i >= 0; i-- {
			prev := topic.PrevQuestion()
			if prev != questions[i] {
				t.Errorf("Expected %v, got %v", questions[i], prev)
			}
		}

		for i := 1; i < len(topic.Questions); i++ {
			next := topic.NextQuestion()
			if next != questions[i] {
				t.Errorf("Expected %v, got %v", questions[i], next)
			}
		}
	})
}
