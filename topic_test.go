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
		topic := NewTopic("A topic", []Question{
			{Statement: "Q1", Answer: "A1"},
			{Statement: "Q2", Answer: "A2"},
			{Statement: "Q3", Answer: "A3"},
			{Statement: "Q4", Answer: "A4"},
		})

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
}
