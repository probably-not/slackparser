package sections

import "time"

var _ Section = (*LocalizedDate)(nil)

type LocalizedDate struct {
	unix         int64
	format       string
	optionalLink string
	fallback     string
}

func NewLocalizedDate(unix int64, format, optionalLink, fallback string) *LocalizedDate {
	return &LocalizedDate{
		unix:         unix,
		format:       format,
		optionalLink: optionalLink,
		fallback:     fallback,
	}
}

func (ld *LocalizedDate) ToPlainText() (string, error) {
	return time.Unix(ld.unix, 0).Format(time.RFC1123), nil
}

func (*LocalizedDate) Type() string {
	return "localized_date"
}
