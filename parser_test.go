package main

import (
	"reflect"
	"testing"
)

func TestParser(t *testing.T) {
	t.Run("MarkdownParser", func(t *testing.T) {
		parser := MarkdownParser{}
		markdown := `# Topic title
## Question 1
Answer 1
**very important text**

## Question 2
Answer 2
- some item
- some other item

## Question 3
### Answer 3`

		contents := []byte(markdown)
		topic := parser.Parse(contents)

		if topic.Title != "Topic title" {
			t.Errorf("Expected %v, got %v", "Topic title", topic.Title)
		}
		if len(topic.Questions) != 3 {
			t.Errorf("Expected %v questions, got %v", 3, len(topic.Questions))
		}
		questions := []Question{
			{Statement: "Question 1", Answer: "Answer 1\n**very important text**"},
			{Statement: "Question 2", Answer: "Answer 2\n- some item\n- some other item"},
			{Statement: "Question 3", Answer: "### Answer 3"},
		}
		if !reflect.DeepEqual(topic.Questions, questions) {
			t.Errorf("Expected %v questions, got %v", questions, topic.Questions)
		}
	})
}
