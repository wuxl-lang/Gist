package main

import (
	"strings"
	"os"
	"fmt"
	"io/ioutil"
)

const gistURL = "https://github.com/wuxl-lang/Gist/blob/master"

func visitTopics() []string {
	files, err := ioutil.ReadDir(".")
	
	if err != nil {
		panic(err)
	}

	var topics []string
	for _, file := range files {
		if isTopicFolder(file) {
			topics = append(topics, file.Name())
		}
	}

	return topics 
}

func visitGists(topic string) []string {
	var dir = fmt.Sprintf("./%s/", topic)
	files, err := ioutil.ReadDir(dir)

	if err != nil {
		panic(err)
	}

	var gists []string
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".md") && file.Name() != "README.md" {
			gists = append(gists, file.Name())
		}
	}

	return gists
}

func isTopicFolder(file os.FileInfo) bool {
	if !file.IsDir() {
		return false
	}

	return !strings.HasPrefix(file.Name(), ".")
}

func buildTopicIndex(topic string, gists []string) string {
	var lines []string

	lines = append(lines, fmt.Sprintf("## %s", topic))
	
	var content []string 
	var footer []string

	for _, gist := range gists {
		content = append(content, fmt.Sprintf("* [%s]", gist))
		footer = append(footer, fmt.Sprintf("[%s]:%s/%s/%s", gist, gistURL, topic, gist))
	}

	lines = append(lines, strings.Join(content, "\n"))
	lines = append(lines, strings.Join(footer, "\n"))

	return strings.Join(lines, "\n\n")
}

func writeTopicIndex(topic string, index string) {
	filePath := fmt.Sprintf("./%s/README.md", topic)
	content := []byte(index)

	if err := ioutil.WriteFile(filePath, content, 0644); err != nil {
		panic(err)
	}
}

func main() {
	subDirs := visitTopics()
	for _, subDir := range subDirs {
		topicGists := visitGists(subDir) 
		index := buildTopicIndex(subDir, topicGists)
		writeTopicIndex(subDir, index)
	}
}
