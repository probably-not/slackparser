package sections

type Sections = []Section

type Section interface {
	Type() string
	ToPlainText() (string, error)
}
