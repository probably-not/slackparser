package sections

var _ Section = (*UserGroupMention)(nil)

type UserGroupMention struct {
	userGroupID string
	label       string
}

func NewUserGroupMention(userGroupID, label string) *UserGroupMention {
	return &UserGroupMention{
		userGroupID: userGroupID,
		label:       label,
	}
}

func (ugm *UserGroupMention) ToPlainText() (string, error) {
	if ugm.label != "" {
		return "@" + ugm.label, nil
	}

	return "@" + ugm.userGroupID, nil
}

func (*UserGroupMention) Type() string {
	return "user_group_mention"
}
