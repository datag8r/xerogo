package trackingcategories

type TrackingCategory struct {
	TrackingCategoryID string
	Name               string
	Status             string // ? | probs ACTIVE / ARCHIVED
	Options            []TrackingOption
}

type TrackingOption struct {
	TrackingOptionID string
	Name             string
	Status           string // ? | probs ACTIVE / ARCHIVED
}
