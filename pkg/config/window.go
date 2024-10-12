package config

// Window stores the last position and title of the window
type Window struct {
	X      int    `json:"x"`
	Y      int    `json:"y"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Title  string `json:"title"`
}
