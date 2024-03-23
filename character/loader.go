package character

import (
	"encoding/xml"
	"fmt"
)

// a character loaded from the path-of-building xml
type Build struct {
	// example of build element:
	// <Build level="97" targetVersion="3_0" pantheonMajorGod="None" bandit="None" className="Ranger" ascendClassName="Pathfinder" characterLevelAutoMode="false" mainSocketGroup="7" viewMode="IMPORT" pantheonMinorGod="None">
	Level                  int      `xml:"level,attr"`
	TargetVersion          string   `xml:"targetVersion,attr"`
	PantheonMajorGod       string   `xml:"pantheonMajorGod,attr"`
	Bandit                 string   `xml:"bandit,attr"`
	ClassName              string   `xml:"className,attr"`
	AscendClassName        string   `xml:"ascendClassName,attr"`
	CharacterLevelAutoMode bool     `xml:"characterLevelAutoMode,attr"`
	MainSocketGroup        int      `xml:"mainSocketGroup,attr"`
	ViewMode               string   `xml:"viewMode,attr"`
	PantheonMinorGod       string   `xml:"pantheonMinorGod,attr"`
}

type Character struct {
	XMLName xml.Name `xml:"PathOfBuilding"`
        Build   Build    `xml:"Build"`
}

// load a character from the path-of-building xml
func LoadCharacter(xml_string string) (Character, error) {
	var character Character
	err := xml.Unmarshal([]byte(xml_string), &character)
	if err != nil {
		panic(err)
	}
        fmt.Println(character)

	return character, nil
}
