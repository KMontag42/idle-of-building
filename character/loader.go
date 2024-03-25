package character

import (
	"encoding/xml"
)

// a character loaded from the path-of-building xml
type Build struct {
	// example of build element:
	// <Build level="97" targetVersion="3_0" pantheonMajorGod="None" bandit="None" className="Ranger" ascendClassName="Pathfinder" characterLevelAutoMode="false" mainSocketGroup="7" viewMode="IMPORT" pantheonMinorGod="None">
	Level                  int          `xml:"level,attr"`
	TargetVersion          string       `xml:"targetVersion,attr"`
	PantheonMajorGod       string       `xml:"pantheonMajorGod,attr"`
	Bandit                 string       `xml:"bandit,attr"`
	ClassName              string       `xml:"className,attr"`
	AscendClassName        string       `xml:"ascendClassName,attr"`
	CharacterLevelAutoMode bool         `xml:"characterLevelAutoMode,attr"`
	MainSocketGroup        int          `xml:"mainSocketGroup,attr"`
	ViewMode               string       `xml:"viewMode,attr"`
	PantheonMinorGod       string       `xml:"pantheonMinorGod,attr"`
	PlayerStats            []PlayerStat `xml:"PlayerStat"`
}

type Skills struct {
	SortGemsByDPSField  string     `xml:"sortGemsByDPSField,attr"`
	ActiveSkillSet      string     `xml:"activeSkillSet,attr"`
	SortGemsByDPS       string     `xml:"sortGemsByDPS,attr"`
	DefaultGemQuality   string     `xml:"defaultGemQuality,attr"`
	DefaultGemLevel     string     `xml:"defaultGemLevel,attr"`
	ShowSupportGemTypes string     `xml:"showSupportGemTypes,attr"`
	ShowAltQualityGems  string     `xml:"showAltQualityGems,attr"`
	SkillSets           []SkillSet `xml:"SkillSet"`
}

type SkillSet struct {
	ID     string  `xml:"id,attr"`
	Skills []Skill `xml:"Skill"`
}

type Skill struct {
	MainActiveSkillCalcs string `xml:"mainActiveSkillCalcs,attr"`
	IncludeInFullDPS     string `xml:"includeInFullDPS,attr"`
	Label                string `xml:"label,attr"`
	Enabled              string `xml:"enabled,attr"`
	Slot                 string `xml:"slot,attr"`
	MainActiveSkill      string `xml:"mainActiveSkill,attr"`
	Gems                 []Gem  `xml:"Gem"`
}

type Gem struct {
	EnableGlobal2 string `xml:"enableGlobal2,attr"`
	Level         string `xml:"level,attr"`
	GemId         string `xml:"gemId,attr"`
	VariantId     string `xml:"variantId,attr"`
	SkillId       string `xml:"skillId,attr"`
	Quality       string `xml:"quality,attr"`
	QualityId     string `xml:"qualityId,attr"`
	EnableGlobal1 string `xml:"enableGlobal1,attr"`
	Enabled       string `xml:"enabled,attr"`
	Count         string `xml:"count,attr"`
	NameSpec      string `xml:"nameSpec,attr"`
}

type PlayerStat struct {
	Stat  string  `xml:"stat,attr"`
	Value float64 `xml:"value,attr"`
}

type Character struct {
	XMLName    xml.Name `xml:"PathOfBuilding"`
	Build      Build    `xml:"Build"`
	Skills     Skills   `xml:"Skills"`
	Experience float64
}

// load a character from the path-of-building xml
func LoadCharacter(xml_string string) (Character, error) {
	var character Character
	err := xml.Unmarshal([]byte(xml_string), &character)
	if err != nil {
		return character, err
	}
        character.Experience = 0

	return character, nil
}

func (c Character) Dps() float64 {
	for _, stat := range c.Build.PlayerStats {
		if stat.Stat == "CombinedDPS" {
			return stat.Value
		}
	}
	panic("DPS not found")
}

func (c Character) Life() float64 {
	for _, stat := range c.Build.PlayerStats {
		if stat.Stat == "TotalEHP" {
			return stat.Value
		}
	}
	panic("Life not found")
}

func (c Character) SetLife(value float64) {
	for i, stat := range c.Build.PlayerStats {
		if stat.Stat == "TotalEHP" {
			c.Build.PlayerStats[i].Value = value
		}
	}
}

func (c Character) SixLinks() []Skill {
	six_links := []Skill{}
	for _, skillSet := range c.Skills.SkillSets {
		for _, skill := range skillSet.Skills {
			if len(skill.Gems) == 6 {
				six_links = append(six_links, skill)
			}
		}
	}
	return six_links
}
