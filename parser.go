package slackparser

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/probably-not/slackparser/sections"
)

var advFmtPattern = regexp.MustCompile(`<(.*?)>`)

func Parse(mrkdwn string) (*ParsedMessage, error) {
	advFmts := advFmtPattern.FindAllStringSubmatchIndex(mrkdwn, -1)
	if len(advFmts) == 0 {
		// No advanced formatting sections (not including formatting markdown shit yet)
		parsed := ParsedMessage{
			originalText: mrkdwn,
			sections: sections.Sections{
				sections.NewTextSection(mrkdwn),
			},
		}
		return &parsed, nil
	}

	messageSections := make(sections.Sections, 0)
	startIndex := 0

	for _, advMatch := range advFmts {
		fullMatchIdxs := advMatch[0:2]
		subMatchIdxs := advMatch[2:4]

		// If the fullMatch index is greater than 0,
		// then we need to create a simple text section
		// between the startIndex (where the last match ended)
		// and the full match index (where the current match starts).
		if fullMatchIdxs[0] > 0 {
			messageSections = append(messageSections, sections.NewTextSection(mrkdwn[startIndex:fullMatchIdxs[0]]))
		}

		section, err := parseAdvancedFormattingSection(mrkdwn[subMatchIdxs[0]:subMatchIdxs[1]])
		if err != nil {
			return nil, err
		}
		messageSections = append(messageSections, section)
		startIndex = fullMatchIdxs[1]
	}

	if startIndex < len(mrkdwn) {
		messageSections = append(messageSections, sections.NewTextSection(mrkdwn[startIndex:]))
	}

	parsed := ParsedMessage{
		originalText: mrkdwn,
		sections:     messageSections,
	}
	return &parsed, nil
}

func parseAdvancedFormattingSection(extractedValue string) (sections.Section, error) {
	if strings.HasPrefix(extractedValue, "#C") {
		splitLabel := strings.Split(extractedValue, "|")
		channelID := splitLabel[0][1:]
		var label string
		if len(splitLabel) == 2 {
			label = splitLabel[1]
		}
		return sections.NewChannelLink(channelID, label), nil
	}

	if strings.HasPrefix(extractedValue, "@U") || strings.HasPrefix(extractedValue, "@W") {
		splitLabel := strings.Split(extractedValue, "|")
		userID := splitLabel[0][1:]
		var label string
		if len(splitLabel) == 2 {
			label = splitLabel[1]
		}
		return sections.NewUserMention(userID, label), nil
	}

	if strings.HasPrefix(extractedValue, "!subteam^") {
		splitLabel := strings.Split(extractedValue, "|")
		userGroupID := splitLabel[0][9:]
		var label string
		if len(splitLabel) == 2 {
			label = splitLabel[1]
		}
		return sections.NewUserGroupMention(userGroupID, label), nil
	}

	if strings.HasPrefix(extractedValue, "!date^") {
		splitFallback := strings.Split(extractedValue, "|")
		dateFmt := splitFallback[0][6:]

		splitDateFmt := strings.Split(dateFmt, "^")
		unix, err := strconv.ParseInt(splitDateFmt[0], 10, 64)
		if err != nil {
			return nil, err
		}
		format := splitDateFmt[1]
		var optionalLink string
		if len(splitDateFmt) == 3 {
			optionalLink = splitDateFmt[2]
		}

		var fallback string
		if len(splitFallback) == 2 {
			fallback = splitFallback[1]
		}
		return sections.NewLocalizedDate(unix, format, optionalLink, fallback), nil
	}

	if strings.HasPrefix(extractedValue, "!") {
		splitLabel := strings.Split(extractedValue, "|")
		mention := splitLabel[0][1:]
		var label string
		if len(splitLabel) == 2 {
			label = splitLabel[1]
		}
		return sections.NewSpecialMention(mention, label), nil
	}

	splitLabel := strings.Split(extractedValue, "|")
	url := splitLabel[0]
	var label string
	if len(splitLabel) == 2 {
		label = splitLabel[1]
	}
	return sections.NewUrlLink(url, label), nil
}
