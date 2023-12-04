package models

type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Location     *Location
	Relation     *Relation
}

type Relations struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

type Location struct {
	ID       int      `json:"id"`
	Location []string `json:"locations"`
}

type Relation struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}



func (a *Artist) GetLocationAndData() string {
	res := ""

	if a.Location != nil {
		for _, location := range a.Location.Location {
			res += location + "\n"

			if a.Relation != nil {
				arr, ok := a.Relation.DatesLocations[location]
				if ok {
					for _, date := range arr {
						res += "\t" + date + "\n"
					}
				}
			}
		}
	}

	return res
}
