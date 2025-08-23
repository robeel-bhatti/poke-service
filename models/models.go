// Package models provides structs that represents API requests / responses
// and app DTOs
package models

type PokeBasic struct {
	Name   string `json:"name"`
	Number int    `json:"nationalDexNum"`
	Type   Types  `json:"types"`
	Sprite string `json:"displaySprite"`
}

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
	BackDefault      string        `json:"back_default"`
	BackFemale       string        `json:"back_female"`
	BackShiny        string        `json:"back_shiny"`
	BackShinyFemale  string        `json:"back_shiny_female"`
	FrontDefault     string        `json:"front_default"`
	FrontFemale      string        `json:"front_female"`
	FrontShiny       string        `json:"front_shiny"`
	FrontShinyFemale string        `json:"front_shiny_female"`
	OtherSprite      OtherSprite   `json:"other"`
	VersionSprite    VersionSprite `json:"versions"`
}

type OtherSprite struct {
	DreamWorld struct {
		FrontDefault string `json:"front_default"`
		FrontFemale  any    `json:"front_female"`
	} `json:"dream_world"`
	Home struct {
		FrontDefault     string `json:"front_default"`
		FrontFemale      any    `json:"front_female"`
		FrontShiny       string `json:"front_shiny"`
		FrontShinyFemale any    `json:"front_shiny_female"`
	} `json:"home"`
	OfficialArtwork struct {
		FrontDefault string `json:"front_default"`
		FrontShiny   string `json:"front_shiny"`
	} `json:"official-artwork"`
	Showdown struct {
		BackDefault      string `json:"back_default"`
		BackFemale       any    `json:"back_female"`
		BackShiny        string `json:"back_shiny"`
		BackShinyFemale  any    `json:"back_shiny_female"`
		FrontDefault     string `json:"front_default"`
		FrontFemale      any    `json:"front_female"`
		FrontShiny       string `json:"front_shiny"`
		FrontShinyFemale any    `json:"front_shiny_female"`
	} `json:"showdown"`
}

type VersionSprite struct {
	GenerationI struct {
		RedBlue struct {
			BackDefault      string `json:"back_default"`
			BackGray         string `json:"back_gray"`
			BackTransparent  string `json:"back_transparent"`
			FrontDefault     string `json:"front_default"`
			FrontGray        string `json:"front_gray"`
			FrontTransparent string `json:"front_transparent"`
		} `json:"red-blue"`
		Yellow struct {
			BackDefault      string `json:"back_default"`
			BackGray         string `json:"back_gray"`
			BackTransparent  string `json:"back_transparent"`
			FrontDefault     string `json:"front_default"`
			FrontGray        string `json:"front_gray"`
			FrontTransparent string `json:"front_transparent"`
		} `json:"yellow"`
	} `json:"generation-i"`
	GenerationIi struct {
		Crystal struct {
			BackDefault           string `json:"back_default"`
			BackShiny             string `json:"back_shiny"`
			BackShinyTransparent  string `json:"back_shiny_transparent"`
			BackTransparent       string `json:"back_transparent"`
			FrontDefault          string `json:"front_default"`
			FrontShiny            string `json:"front_shiny"`
			FrontShinyTransparent string `json:"front_shiny_transparent"`
			FrontTransparent      string `json:"front_transparent"`
		} `json:"crystal"`
		Gold struct {
			BackDefault      string `json:"back_default"`
			BackShiny        string `json:"back_shiny"`
			FrontDefault     string `json:"front_default"`
			FrontShiny       string `json:"front_shiny"`
			FrontTransparent string `json:"front_transparent"`
		} `json:"gold"`
		Silver struct {
			BackDefault      string `json:"back_default"`
			BackShiny        string `json:"back_shiny"`
			FrontDefault     string `json:"front_default"`
			FrontShiny       string `json:"front_shiny"`
			FrontTransparent string `json:"front_transparent"`
		} `json:"silver"`
	} `json:"generation-ii"`
	GenerationIii struct {
		Emerald struct {
			FrontDefault string `json:"front_default"`
			FrontShiny   string `json:"front_shiny"`
		} `json:"emerald"`
		FireredLeafgreen struct {
			BackDefault  string `json:"back_default"`
			BackShiny    string `json:"back_shiny"`
			FrontDefault string `json:"front_default"`
			FrontShiny   string `json:"front_shiny"`
		} `json:"firered-leafgreen"`
		RubySapphire struct {
			BackDefault  string `json:"back_default"`
			BackShiny    string `json:"back_shiny"`
			FrontDefault string `json:"front_default"`
			FrontShiny   string `json:"front_shiny"`
		} `json:"ruby-sapphire"`
	} `json:"generation-iii"`
	GenerationIv struct {
		DiamondPearl struct {
			BackDefault      string `json:"back_default"`
			BackFemale       any    `json:"back_female"`
			BackShiny        string `json:"back_shiny"`
			BackShinyFemale  any    `json:"back_shiny_female"`
			FrontDefault     string `json:"front_default"`
			FrontFemale      any    `json:"front_female"`
			FrontShiny       string `json:"front_shiny"`
			FrontShinyFemale any    `json:"front_shiny_female"`
		} `json:"diamond-pearl"`
		HeartgoldSoulsilver struct {
			BackDefault      string `json:"back_default"`
			BackFemale       any    `json:"back_female"`
			BackShiny        string `json:"back_shiny"`
			BackShinyFemale  any    `json:"back_shiny_female"`
			FrontDefault     string `json:"front_default"`
			FrontFemale      any    `json:"front_female"`
			FrontShiny       string `json:"front_shiny"`
			FrontShinyFemale any    `json:"front_shiny_female"`
		} `json:"heartgold-soulsilver"`
		Platinum struct {
			BackDefault      string `json:"back_default"`
			BackFemale       any    `json:"back_female"`
			BackShiny        string `json:"back_shiny"`
			BackShinyFemale  any    `json:"back_shiny_female"`
			FrontDefault     string `json:"front_default"`
			FrontFemale      any    `json:"front_female"`
			FrontShiny       string `json:"front_shiny"`
			FrontShinyFemale any    `json:"front_shiny_female"`
		} `json:"platinum"`
	} `json:"generation-iv"`
	GenerationV struct {
		BlackWhite struct {
			Animated struct {
				BackDefault      string `json:"back_default"`
				BackFemale       any    `json:"back_female"`
				BackShiny        string `json:"back_shiny"`
				BackShinyFemale  any    `json:"back_shiny_female"`
				FrontDefault     string `json:"front_default"`
				FrontFemale      any    `json:"front_female"`
				FrontShiny       string `json:"front_shiny"`
				FrontShinyFemale any    `json:"front_shiny_female"`
			} `json:"animated"`
			BackDefault      string `json:"back_default"`
			BackFemale       any    `json:"back_female"`
			BackShiny        string `json:"back_shiny"`
			BackShinyFemale  any    `json:"back_shiny_female"`
			FrontDefault     string `json:"front_default"`
			FrontFemale      any    `json:"front_female"`
			FrontShiny       string `json:"front_shiny"`
			FrontShinyFemale any    `json:"front_shiny_female"`
		} `json:"black-white"`
	} `json:"generation-v"`
	GenerationVi struct {
		OmegarubyAlphasapphire struct {
			FrontDefault     string `json:"front_default"`
			FrontFemale      any    `json:"front_female"`
			FrontShiny       string `json:"front_shiny"`
			FrontShinyFemale any    `json:"front_shiny_female"`
		} `json:"omegaruby-alphasapphire"`
		XY struct {
			FrontDefault     string `json:"front_default"`
			FrontFemale      any    `json:"front_female"`
			FrontShiny       string `json:"front_shiny"`
			FrontShinyFemale any    `json:"front_shiny_female"`
		} `json:"x-y"`
	} `json:"generation-vi"`
	GenerationVii struct {
		Icons struct {
			FrontDefault string `json:"front_default"`
			FrontFemale  any    `json:"front_female"`
		} `json:"icons"`
		UltraSunUltraMoon struct {
			FrontDefault     string `json:"front_default"`
			FrontFemale      any    `json:"front_female"`
			FrontShiny       string `json:"front_shiny"`
			FrontShinyFemale any    `json:"front_shiny_female"`
		} `json:"ultra-sun-ultra-moon"`
	} `json:"generation-vii"`
	GenerationViii struct {
		Icons struct {
			FrontDefault string `json:"front_default"`
			FrontFemale  any    `json:"front_female"`
		} `json:"icons"`
	} `json:"generation-viii"`
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
