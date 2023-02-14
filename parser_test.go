package slackparser

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/probably-not/slackparser/sections"
)

func TestParse(t *testing.T) {
	testCases := []struct {
		desc     string
		mrkdwn   string
		expected *ParsedMessage
	}{
		{
			desc:   "normal text",
			mrkdwn: "this is some normal text without any formatting",
			expected: &ParsedMessage{
				originalText: "this is some normal text without any formatting",
				sections: sections.Sections{
					sections.NewTextSection("this is some normal text without any formatting"),
				},
			},
		},
		{
			desc:   "channel link",
			mrkdwn: "this is some message with a channel link to the <#C024BE7LR> channel and the <#C024BE7LQ> channel",
			expected: &ParsedMessage{
				originalText: "this is some message with a channel link to the <#C024BE7LR> channel and the <#C024BE7LQ> channel",
				sections: sections.Sections{
					sections.NewTextSection("this is some message with a channel link to the "),
					sections.NewChannelLink("C024BE7LR", ""),
					sections.NewTextSection(" channel and the "),
					sections.NewChannelLink("C024BE7LQ", ""),
					sections.NewTextSection(" channel"),
				},
			},
		},
		{
			desc:   "channel link with label",
			mrkdwn: "this is some message with a channel link to the <#C024BE7LR|random-channel-name> channel and the <#C024BE7LQ> channel",
			expected: &ParsedMessage{
				originalText: "this is some message with a channel link to the <#C024BE7LR|random-channel-name> channel and the <#C024BE7LQ> channel",
				sections: sections.Sections{
					sections.NewTextSection("this is some message with a channel link to the "),
					sections.NewChannelLink("C024BE7LR", "random-channel-name"),
					sections.NewTextSection(" channel and the "),
					sections.NewChannelLink("C024BE7LQ", ""),
					sections.NewTextSection(" channel"),
				},
			},
		},
		{
			desc:   "user mention",
			mrkdwn: "this is some message with a user mention to <@U123123123> and <@W123123123>",
			expected: &ParsedMessage{
				originalText: "this is some message with a user mention to <@U123123123> and <@W123123123>",
				sections: sections.Sections{
					sections.NewTextSection("this is some message with a user mention to "),
					sections.NewUserMention("U123123123", ""),
					sections.NewTextSection(" and "),
					sections.NewUserMention("W123123123", ""),
				},
			},
		},
		{
			desc:   "user mention with label",
			mrkdwn: "this is some message with a user mention to <@U123123123|Coby> and <@W123123123>",
			expected: &ParsedMessage{
				originalText: "this is some message with a user mention to <@U123123123|Coby> and <@W123123123>",
				sections: sections.Sections{
					sections.NewTextSection("this is some message with a user mention to "),
					sections.NewUserMention("U123123123", "Coby"),
					sections.NewTextSection(" and "),
					sections.NewUserMention("W123123123", ""),
				},
			},
		},
		{
			desc:   "user group mention",
			mrkdwn: "this is some message with a user group mention to <!subteam^SAZ94GDB8>",
			expected: &ParsedMessage{
				originalText: "this is some message with a user group mention to <!subteam^SAZ94GDB8>",
				sections: sections.Sections{
					sections.NewTextSection("this is some message with a user group mention to "),
					sections.NewUserGroupMention("SAZ94GDB8", ""),
				},
			},
		},
		{
			desc:   "user group mention with label",
			mrkdwn: "this is some message with a user group mention to <!subteam^SAZ94GDB8|random-group>",
			expected: &ParsedMessage{
				originalText: "this is some message with a user group mention to <!subteam^SAZ94GDB8|random-group>",
				sections: sections.Sections{
					sections.NewTextSection("this is some message with a user group mention to "),
					sections.NewUserGroupMention("SAZ94GDB8", "random-group"),
				},
			},
		},
		{
			desc:   "localized date",
			mrkdwn: "this is some message with a localized date formatting <!date^1392734382^Posted {date_num} {time_secs}|Posted 2014-02-18 6:39:42 AM PST> here",
			expected: &ParsedMessage{
				originalText: "this is some message with a localized date formatting <!date^1392734382^Posted {date_num} {time_secs}|Posted 2014-02-18 6:39:42 AM PST> here",
				sections: sections.Sections{
					sections.NewTextSection("this is some message with a localized date formatting "),
					sections.NewLocalizedDate(1392734382, "Posted {date_num} {time_secs}", "", "Posted 2014-02-18 6:39:42 AM PST"),
					sections.NewTextSection(" here"),
				},
			},
		},
		{
			desc:   "special mention",
			mrkdwn: "<!here> this is some message with a special mention to <!channel>",
			expected: &ParsedMessage{
				originalText: "<!here> this is some message with a special mention to <!channel>",
				sections: sections.Sections{
					sections.NewSpecialMention("here", ""),
					sections.NewTextSection(" this is some message with a special mention to "),
					sections.NewSpecialMention("channel", ""),
				},
			},
		},
		{
			desc:   "special mention with label",
			mrkdwn: "<!here|here> this is some message with a special mention to <!channel>",
			expected: &ParsedMessage{
				originalText: "<!here|here> this is some message with a special mention to <!channel>",
				sections: sections.Sections{
					sections.NewSpecialMention("here", "here"),
					sections.NewTextSection(" this is some message with a special mention to "),
					sections.NewSpecialMention("channel", ""),
				},
			},
		},
		{
			desc:   "url link",
			mrkdwn: "This message contains a URL <http://example.com/>",
			expected: &ParsedMessage{
				originalText: "This message contains a URL <http://example.com/>",
				sections: sections.Sections{
					sections.NewTextSection("This message contains a URL "),
					sections.NewUrlLink("http://example.com/", ""),
				},
			},
		},
		{
			desc:   "url link with label",
			mrkdwn: "This message contains a URL <http://example.com/|wow link>",
			expected: &ParsedMessage{
				originalText: "This message contains a URL <http://example.com/|wow link>",
				sections: sections.Sections{
					sections.NewTextSection("This message contains a URL "),
					sections.NewUrlLink("http://example.com/", "wow link"),
				},
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(subT *testing.T) {
			got, err := Parse(tC.mrkdwn)
			if err != nil {
				subT.Errorf("expected no error, got error: %v", err)
				return
			}

			diff := cmp.Diff(
				tC.expected,
				got,
				cmp.AllowUnexported(
					ParsedMessage{},
					sections.ChannelLink{},
					sections.LocalizedDate{},
					sections.SpecialMention{},
					sections.TextSection{},
					sections.UrlLink{},
					sections.UserGroupMention{},
					sections.UserMention{},
				),
			)
			if diff != "" {
				subT.Errorf("Expected: %+v, Got: %+v; Diff: %s", tC.expected, got, diff)
			}
		})
	}
}
