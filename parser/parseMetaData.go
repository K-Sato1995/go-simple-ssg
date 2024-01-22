package parser

import (
	"bytes"
	"errors"
	"strings"
)

type MetaData struct {
	Title       string
	Description string
}

func ParseMetadata(content []byte) (MetaData, []byte, error) {
	var meta MetaData
	sections := bytes.SplitN(content, []byte("----"), 3)
	if len(sections) < 3 {
		return meta, nil, errors.New("invalid format: metadata not found")
	}
	metadataContent := sections[1]
	mdContent := sections[2]
	lines := bytes.Split(metadataContent, []byte("\n"))
	for _, line := range lines {
		line = bytes.TrimLeft(line, "- ")
		keyValue := bytes.SplitN(line, []byte(":"), 2)
		if len(keyValue) != 2 {
			continue
		}
		key := string(bytes.TrimSpace(keyValue[0]))
		value := string(bytes.TrimSpace(keyValue[1]))
		value = strings.Trim(value, "\"")
		switch key {
		case "Title":
			meta.Title = value
		case "Description":
			meta.Description = value
		}
	}
	return meta, mdContent, nil
}
