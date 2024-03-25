package character

import (
	"encoding/xml"
	"fmt"
	"log"

	"github.com/kmontag42/idle-of-building/types"
)

// load a character from the path-of-building xml
func LoadCharacter(xml_string string) (types.Character, error) {
	var character types.Character
	err := xml.Unmarshal([]byte(xml_string), &character)
	if err != nil {
		return character, err
	}
	character.Experience = 0

        // build the name based on the class and ascendancy
        // as well as the name of the primary six link
        six_links := character.SixLinks()
        log.Println(six_links)
        if len(six_links) > 0 {
          main_skill_gem_name := six_links[0].Gems[0].NameSpec
          character.Name = fmt.Sprintf("%s (%s)", character.Build.AscendClassName, main_skill_gem_name)
        } else {
          character.Name = character.Build.AscendClassName
        }

        character.MapResults = []types.MapResult{}
        
        id := len(types.CharactersMap) + 1
        character.Id = id
        types.CharactersMap[id] = character

	return character, nil
}

