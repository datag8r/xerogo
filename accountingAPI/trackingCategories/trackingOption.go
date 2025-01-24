package trackingcategories

type TrackingOption struct {
	TrackingOptionID string
	Name             string
	Status           string // ? | probs ACTIVE / ARCHIVED
}

func GetTrackingOptions()    {}
func GetTrackingOption()     {}
func UpdateTrackingOption()  {}
func CreateTrackingOption()  {}
func ArchiveTrackingOption() {}
func DeleteTrackingOption()  {}
