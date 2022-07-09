package main

import "strings"

type Parser interface {
	Parse(contents []byte) Topic
}

type MarkdownParser struct{}

func (MarkdownParser) Parse(contents []byte) Topic {
	topic := NewTopic()

	statement := ""
	answerLines := make([]string, 0)

	lines := strings.Split(string(contents), "\n")
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		count := strings.Count(line, "#")
		if count == 2 {
			if len(answerLines) > 0 {
				topic.AddQuestion(Question{
					Statement: statement,
					Answer:    strings.Join(answerLines, "\n"),
				})
				answerLines = make([]string, 0)
			}
			statement = strings.Trim(strings.Replace(line, "##", "", -1), " ")
		} else if count == 1 {
			topic.Title = strings.Trim(strings.Replace(line, "#", "", -1), " ")
		} else {
			answerLines = append(answerLines, line)
		}
	}

	if len(answerLines) > 0 {
		topic.AddQuestion(Question{
			Statement: statement,
			Answer:    strings.Join(answerLines, "\n"),
		})
		answerLines = make([]string, 0)
	}
	return topic
}

func NewMarkdownParser() MarkdownParser {
	return MarkdownParser{}
}
