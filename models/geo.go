package models

type Country struct {
	Short   string
	Long    string
	Regions []Region
}

type Region struct {
	Name string
}

// type City struct {
// }
