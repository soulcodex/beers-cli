package model

type Beer struct {
	Id               int          `json:"id"`
	Name             string       `json:"name"`
	Tagline          string       `json:"tagline"`
	FirstBrewed      string       `json:"first_brewed"`
	Description      string       `json:"description"`
	ImageUrl         string       `json:"image_url"`
	Abv              float32      `json:"abv"`
	Ibu              float32      `json:"ibu"`
	TargetFg         float32      `json:"target_fg"`
	TargetOg         float32      `json:"target_og"`
	Ebc              float32      `json:"ebc"`
	Srm              float32      `json:"srm"`
	Ph               float32      `json:"ph"`
	AttenuationValue int          `json:"attenuation_value"`
	Volume           *BeerVolume  `json:"volume"`
	FoodPairing      *FoodPairing `json:"food_pairing"`
	ContributedBy    string       `json:"contributed_by"`
}

type BeerVolume struct {
	Value float32 `json:"value"`
	Unit  string  `json:"unit"`
}

type FoodPairing []string
