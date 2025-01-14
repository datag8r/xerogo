package brandingthemes

type BrandingTheme struct {
	BrandingThemeID string `json:",omitempty"`
	Name            string
	LogoURL         string
	Type            string // doc type
	SortOrder       int
	CreatedDateUTC  string
}

func (b BrandingTheme) IsZero() bool {
	return b.BrandingThemeID == ""
}
