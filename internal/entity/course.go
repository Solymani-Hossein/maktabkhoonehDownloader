package entity

type Course struct {
	Name    string
	Seasons []Season
	About   About
}

type Season struct {
	Title string
	Links []Link
}

type Link struct {
	Name string
	Href string
}
type About struct {
	// TODO info about course
}

type Quality int

const (
	Low Quality = iota + 1
	Medium
	High
)

type Args struct {
	Url     string
	Quality Quality
}
