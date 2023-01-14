package flights

type Airport string

type Route struct {
	StartAirport Airport
	EndAirport   Airport
}

type Record struct {
	FullRoute Route   `json:"full_route"`
	FullPath  []Route `json:"full_path"`
}
