package models

//From 2 days before
type FeatureCollection struct {
	MetaD    Metadata  `json:"metadata, omitempty"`
	Features []Feature `json:"features, omitempty"`
}

//from ID
type SpecificID struct {
	Propertie Properties `json:"properties, omitempty"`
	Geo       Geometry   `json:"geometry, omitempty"`
	ID        string     `json:"id, omitempty"`
}

type Metadata struct {
	Status int `json:"status, omitempty"`
}

type Feature struct {
	Propertie Properties `json:"properties, omitempty"`
	Geo       Geometry   `json:"geometry, omitempty"`
	ID        string     `json:"id, omitempty"`
}

type Properties struct {
	Magnitud float64 `json:"mag, omitempty"`
	OTime    int64   `json:"time, omitempty"`
	UTime    int64   `json:"updated, omitempty"`
	Type     string  `json:"type, omitempty"`
	Title    string  `json:"title, omitempty"`
}

type Geometry struct {
	Coordinates []float64 `json:"coordinates, omitempty"`
}

type Response struct {
	ID       string
	Title    string
	Magnitud string
	Type     string
	Time     string
	UpdatedT string
	Depth    string
	Lat      string
	Lon      string
}

type ErrorResponse struct {
	Code int
	Msg  string
}

type PeriodTime struct {
	Start string `json:"start, omitempty"`
	End   string `json:"end, omitempty"`
}
