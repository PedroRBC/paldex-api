package pkg

// Pal represents a Pal creature in the game
type Pal struct {
	ID          int    `json:"id"`
	Key         string `json:"key"`
	Image       string `json:"image"`
	Name        string `json:"name"`
	Wiki        string `json:"wiki"`
	Types       []Type `json:"types"`
	ImageWiki   string `json:"imageWiki"`
	Suitability []Suitability `json:"suitability"`
	Drops       []string `json:"drops"`
	Aura        Aura `json:"aura"`
	Description string `json:"description"`
	Skills      []Skill `json:"skills"`
	Stats       Stats `json:"stats"`
	Asset       string `json:"asset"`
	Genus       string `json:"genus"`
	Rarity      int    `json:"rarity"`
	Price       int    `json:"price"`
	Size        string `json:"size"`
	Maps        Maps `json:"maps"`
	Breeding    Breeding `json:"breeding"`
}

type Type struct {
	Name  string `json:"name"`
	Image string `json:"image"`
}

type Suitability struct {
	Type  string `json:"type"`
	Image string `json:"image"`
	Level int    `json:"level"`
}

type Aura struct {
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Tech        interface{} `json:"tech"`
}

type Skill struct {
	Level       int    `json:"level"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	Cooldown    int    `json:"cooldown"`
	Power       int    `json:"power"`
	Description string `json:"description"`
}

type Stats struct {
	HP      int `json:"hp"`
	Attack  Attack `json:"attack"`
	Defense int `json:"defense"`
	Speed   Speed `json:"speed"`
	Stamina int `json:"stamina"`
	Support int `json:"support"`
	Food    int `json:"food"`
}

type Attack struct {
	Melee  int `json:"melee"`
	Ranged int `json:"ranged"`
}

type Speed struct {
	Ride int `json:"ride"`
	Run  int `json:"run"`
	Walk int `json:"walk"`
}

type Maps struct {
	Day   string `json:"day"`
	Night string `json:"night"`
}

type Breeding struct {
	Rank            int     `json:"rank"`
	Order           int     `json:"order"`
	ChildEligble    bool    `json:"child_eligble"`
	MaleProbability float64 `json:"male_probability"`
} 