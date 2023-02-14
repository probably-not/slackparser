package sections

var _ Section = (*SpecialMention)(nil)

type SpecialMention struct {
	mention string
	label   string
}

func NewSpecialMention(mention, label string) *SpecialMention {
	return &SpecialMention{
		mention: mention,
		label:   label,
	}
}

func (sm *SpecialMention) ToPlainText() (string, error) {
	if sm.label != "" {
		return "@" + sm.label, nil
	}

	return "@" + sm.mention, nil
}

func (*SpecialMention) Type() string {
	return "special_mention"
}
