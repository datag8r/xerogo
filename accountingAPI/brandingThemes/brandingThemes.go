package brandingthemes

type BrandingTheme struct {
	BrandingThemeID string
	Name            string
	LogoURL         string
	Type            string // doc type
	SortOrder       int
	CreatedDateUTC  string
}
