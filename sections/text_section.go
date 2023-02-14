package sections

var _ Section = (*TextSection)(nil)

type TextSection struct {
	text string
}

func NewTextSection(text string) *TextSection {
	return &TextSection{text: text}
}

func (ts *TextSection) ToPlainText() (string, error) {
	return ts.text, nil
}

func (*TextSection) Type() string {
	return "text"
}
