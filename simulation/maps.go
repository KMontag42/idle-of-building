package main

import (
  "sync"
  "time"
  "log"
)

type Entity struct {
  Name string
  HP   int
  DPS  int
}

type Character struct {
  Entity
}

type Enemy struct {
  Entity
}

func Battle(hero *Character, enemy Enemy) bool {
  for hero.HP > 0 && enemy.HP > 0 {
    enemy.HP -= hero.DPS
    hero.HP -= enemy.DPS
    log.Printf("%s HP: %d\n", hero.Name, hero.HP)
    log.Printf("%s HP: %d\n", enemy.Name, enemy.HP)
    time.Sleep(1 * time.Second)
  }

  if hero.HP <= 0 {
    log.Printf("%s has been defeated\n", hero.Name)
    return false
  } else {
    log.Printf("%s has been defeated\n", enemy.Name)
    return true
  }
}

func RunMap(hero *Character, enemies []Enemy, wg *sync.WaitGroup) bool {
  defer wg.Done()
  log.Printf("%s has entered the map\n", hero.Name)
  for _, enemy := range enemies {
    log.Printf("%s has encountered %s\n", hero.Name, enemy.Name)
    if Battle(hero, enemy) {
      log.Printf("%s has won the battle\n", hero.Name)
    } else {
      log.Printf("%s has lost the battle\n", hero.Name)
      return false
    }
  }
  log.Printf("%s has cleared the map\n", hero.Name)
  return true
}

// TODO: rename this to something else after testing
func main() {
  var wg sync.WaitGroup

  hero := &Character{
    Entity: Entity{
      Name: "Hero",
      HP:   100,
      DPS:  10,
    },
  }
  hero2 := &Character{
    Entity: Entity{
      Name: "Hero 2",
      HP:   100,
      DPS:  10,
    },
  }

  enemies := []Enemy{
    {
      Entity: Entity{
        Name: "Enemy 1",
        HP:   50,
        DPS:  5,
      },
    },
    {
      Entity: Entity{
        Name: "Enemy 2",
        HP:   50,
        DPS:  5,
      },
    },
  }

  maps := map[Character][]Enemy{
    *hero: enemies,
    *hero2: enemies,
  }

  for hero, enemies := range maps {
    wg.Add(1)
    go RunMap(&hero, enemies, &wg)
  }

  wg.Wait()
}
