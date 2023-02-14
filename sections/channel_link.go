package sections

var _ Section = (*ChannelLink)(nil)

type ChannelLink struct {
	channelID string
	label     string
}

func NewChannelLink(channelID, label string) *ChannelLink {
	return &ChannelLink{
		channelID: channelID,
		label:     label,
	}
}

func (cl *ChannelLink) ToPlainText() (string, error) {
	if cl.label != "" {
		return cl.label, nil
	}

	return cl.channelID, nil
}

func (*ChannelLink) Type() string {
	return "channel_link"
}
