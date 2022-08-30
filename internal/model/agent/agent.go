package agent

// Agent is representing agent in the alive system
type Agent struct {
	ID       int64  `json:"id"`
	Location string `json:"location"`
	GeoHash  string `json:"geohash"`
	ISP      string `json:"ISP"`
}
