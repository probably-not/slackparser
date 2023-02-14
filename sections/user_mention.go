package sections

var _ Section = (*UserMention)(nil)

type UserMention struct {
	userID string
	label  string
}

func NewUserMention(userID, label string) *UserMention {
	return &UserMention{
		userID: userID,
		label:  label,
	}
}

func (um *UserMention) ToPlainText() (string, error) {
	if um.label != "" {
		return "@" + um.label, nil
	}

	return "@" + um.userID, nil
}

func (*UserMention) Type() string {
	return "user_mention"
}
