package slackparser

import "github.com/probably-not/slackparser/sections"

type ParsedMessage struct {
	originalText string
	sections     sections.Sections
}
