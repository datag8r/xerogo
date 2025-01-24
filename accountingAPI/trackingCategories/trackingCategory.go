package trackingcategories

type TrackingCategory struct {
	TrackingCategoryID string
	Name               string
	Status             string // ? | probs ACTIVE / ARCHIVED
	Options            []TrackingOption
}

type Tracking struct {
	Name               string `xero:"create,update"`
	Option             string `xero:"create,update"`
	TrackingCategoryID string
}

func GetTrackingCategories()   {}
func GetTrackingCategory()     {} //includeArchived
func UpdateTrackingCategory()  {}
func CreateTrackingCategory()  {}
func ArchiveTrackingCategory() {}
func DeleteTrackingCategory()  {}
