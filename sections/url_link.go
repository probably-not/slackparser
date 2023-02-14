package sections

var _ Section = (*UrlLink)(nil)

type UrlLink struct {
	url   string
	label string
}

func NewUrlLink(url, label string) *UrlLink {
	return &UrlLink{
		url:   url,
		label: label,
	}
}

func (ul *UrlLink) ToPlainText() (string, error) {
	if ul.label != "" {
		return ul.label, nil
	}

	return ul.url, nil
}

func (*UrlLink) Type() string {
	return "url_link"
}
