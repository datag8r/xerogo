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

func GetTrackingCategories()   {}
func GetTrackingCategory()     {} //includeArchived
func UpdateTrackingCategory()  {}
func CreateTrackingCategory()  {}
func ArchiveTrackingCategory() {}
func DeleteTrackingCategory()  {}

func GetTrackingOptions()    {}
func GetTrackingOption()     {}
func UpdateTrackingOption()  {}
func CreateTrackingOption()  {}
func ArchiveTrackingOption() {}
func DeleteTrackingOption()  {}
