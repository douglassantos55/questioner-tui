package main

import (
	"os"
	"path"
)

type Loader interface {
	GetTopics(path string) []Topic
}

type LocalLoader struct {
	Parser Parser
}

func NewLocalLoader(parser Parser) LocalLoader {
	return LocalLoader{
		Parser: parser,
	}
}

func (l LocalLoader) GetTopics(source string) []Topic {
	topics := make([]Topic, 0)
	for _, path := range l.GetFiles(source) {
		if contents, err := os.ReadFile(path); err == nil {
			topics = append(topics, l.Parser.Parse(contents))
		}
	}
	return topics
}

func (l LocalLoader) GetFiles(directory string) []string {
	files := []string{}
	entries, err := os.ReadDir(directory)

	if err != nil {
		return []string{}
	}

	for _, entry := range entries {
		if entry.IsDir() {
			sub := l.GetFiles(path.Join(directory, entry.Name()))
			files = append(files, sub...)
		} else {
			files = append(files, path.Join(directory, entry.Name()))
		}
	}

	return files
}
