package Config

var DefinedWeapon definedWeapon

type definedWeapon struct {
	WeaponCategory struct {
		Type struct {
			Name        string `json:"Name,omitempty"`
			Description string `json:"Description,omitempty"`
			Weapons     struct {
				Position        string `json:"Position,omitempty"`
				Damage          int    `json:"Damage,omitempty"`
				DamageRange     int    `json:"DamageRange,omitempty"`
				DamagePrecision int    `json:"DamagePrecision,omitempty"`
				ExtraAttribute  struct {
					Strength     int `json:"Strength,omitempty"`
					Wisdom       int `json:"Wisdom,omitempty"`
					Intelligent  int `json:"Intelligent,omitempty"`
					Constitution int `json:"Constitution,omitempty"`
					Dexterity    int `json:"Dexterity,omitempty"`
				} `json:"ExtraAttribute,omitempty"`
				DestroyWhenZeroQuota bool     `json:"DestroyWhenZeroQuota,omitempty"`
				AllowExecuteTimes    int      `json:"AllowExecuteTimes,omitempty"`
				Name                 string   `json:"Name,omitempty"`
				SystemCode           string   `json:"SystemCode,omitempty"`
				ShortName            string   `json:"ShortName,omitempty"`
				Description          string   `json:"Description,omitempty"`
				DescriptionDetail    string   `json:"DescriptionDetail,omitempty"`
				Unit                 string   `json:"Unit,omitempty"`
				Weight               int      `json:"Weight,omitempty"`
				Pricing              int      `json:"Pricing,omitempty"`
				Decoration           []string `json:"Decoration,omitempty"`
			} `json:"Weapons,omitempty"`
		} `json:"Type,omitempty"`
	} `json:"WeaponCategory,omitempty"`
}

func init() {
	if err := convertFromFile("DefinedWeapon.json", &DefinedWeapon); err != nil {
		panic(err)
	}
}
