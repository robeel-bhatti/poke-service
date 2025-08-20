// Package models provides structs that represents API requests / responses
// and app DTOs
package models

type PaginatedResponse struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type Pokemon struct {
	Name           string     `json:"name"`
	NationalDexNum string     `json:"nationalDexNum"`
	Species        string     `json:"species"`
	Height         string     `json:"height"`
	Weight         string     `json:"weight"`
	Type           *Types     `json:"type"`
	Ability        *Abilities `json:"ability"`
	BaseStats      *Stats     `json:"baseStats"`
}

type Types struct {
	Primary   string `json:"primary"`
	Secondary string `json:"secondary"`
}

type Abilities struct {
	Primary string `json:"primary"`
	Hidden  string `json:"hidden"`
}

type Stats struct {
	HP             int `json:"hp"`
	Attack         int `json:"attack"`
	Defense        int `json:"defense"`
	SpecialAttack  int `json:"specialAttack"`
	SpecialDefense int `json:"specialDefense"`
	Speed          int `json:"speed"`
}

type NamedAPIResource struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Ability struct {
	Ability  NamedAPIResource `json:"ability"`
	IsHidden bool             `json:"is_hidden"`
	Slot     int              `json:"slot"`
}

type Cry struct {
	Latest string `json:"latest"`
	Legacy string `json:"legacy"`
}

type GameIndex struct {
	GameIndex int              `json:"game_index"`
	Version   NamedAPIResource `json:"version"`
}

type VersionGroupDetail struct {
	LevelLearnedAt  int              `json:"level_learned_at"`
	MoveLearnMethod NamedAPIResource `json:"move_learn_method"`
	Order           int              `json:"order"`
	VersionGroup    NamedAPIResource `json:"version_group"`
}

type Move struct {
	Move                NamedAPIResource     `json:"move"`
	VersionGroupDetails []VersionGroupDetail `json:"version_group_details"`
}

type PastAbilityDetail struct {
	Ability  NamedAPIResource `json:"ability"`
	IsHidden bool             `json:"is_hidden"`
	Slot     int              `json:"slot"`
}

type PastAbility struct {
	Abilities  []PastAbilityDetail `json:"abilities"`
	Generation NamedAPIResource    `json:"generation"`
}

type Sprite struct {
	BackDefault      string `json:"back_default"`
	BackFemale       string `json:"back_female"`
	BackShiny        string `json:"back_shiny"`
	BackShinyFemale  string `json:"back_shiny_female"`
	FrontDefault     string `json:"front_default"`
	FrontFemale      string `json:"front_female"`
	FrontShiny       string `json:"front_shiny"`
	FrontShinyFemale string `json:"front_shiny_female"`
}

type PokemonStat struct {
	BaseStat int              `json:"base_stat"`
	Effort   int              `json:"effort"`
	Stat     NamedAPIResource `json:"stat"`
}

type PokemonType struct {
	Slot int              `json:"slot"`
	Type NamedAPIResource `json:"type"`
}

type HeldItems struct {
	Item NamedAPIResource `json:"item"`
}

type PokemonResponse struct {
	Abilities              []Ability          `json:"abilities"`
	BaseExperience         int                `json:"base_experience"`
	Cries                  Cry                `json:"cries"`
	Forms                  []NamedAPIResource `json:"forms"`
	GameIndices            []GameIndex        `json:"game_indices"`
	Height                 int                `json:"height"`
	HeldItems              []HeldItems        `json:"held_items"`
	Id                     int                `json:"id"`
	IsDefault              bool               `json:"is_default"`
	LocationAreaEncounters string             `json:"location_area_encounters"`
	Moves                  []Move             `json:"moves"`
	Name                   string             `json:"name"`
	Order                  int                `json:"order"`
	PastAbilities          []PastAbility      `json:"past_abilities"`
	PastTypes              []string           `json:"past_types"`
	Species                NamedAPIResource   `json:"species"`
	Sprites                Sprite             `json:"sprites"`
	Stats                  []PokemonStat      `json:"stats"`
	Types                  []PokemonType      `json:"types"`
	Weight                 int                `json:"weight"`
}
