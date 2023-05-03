package model

import (
	"sync"
)

type cityLocation struct {
	Abscissa int
	Ordinate int
}

type city struct {
	Id           int              `json:"id"`
	Name         string           `json:"name"`
	Location     *cityLocation    `json:"location"`
	Image        string           `json:"image"`
	Level        *Level           `json:"level"`
	OwnerId      int              `json:"owner_id"`
	Neighbors    []int            `json:"neighbors"`
	neighborsMap map[int]struct{} `json:"-"`

	lock *sync.Mutex `json:"-"`
}

func (c *city) addNeighbors(cities []*city) {
	if c.neighborsMap == nil {
		c.neighborsMap = make(map[int]struct{})
	}
	if c.Neighbors == nil {
		c.Neighbors = make([]int, 0)
	}
	for _, cc := range cities {
		if cc == nil {
			continue
		}
		c.Neighbors = append(c.Neighbors, cc.Id)
		c.neighborsMap[cc.Id] = struct{}{}
	}
}

func (c *city) LevelUp() {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.Level.Incr()
}

func (c *city) GoToNewCity(newCity *city) bool {
	if newCity == nil {
		return false
	}
	_, ok := c.neighborsMap[newCity.Id]
	return ok
}

type CityMap struct {
	Cities []*city
}

func NewTestCities() []*city {
	city1 := &city{
		Id:   1,
		Name: "白银港",
		Location: &cityLocation{
			Abscissa: 120,
			Ordinate: 110,
		},
		Image:   "",
		Level:   ZeroLevel(),
		OwnerId: 12,
	}

	city2 := &city{
		Id:   2,
		Name: "沙朵镇",
		Location: &cityLocation{
			Abscissa: 120,
			Ordinate: 120,
		},
		Image: "",
		Level: NewLevel(1),
	}

	city3 := &city{
		Id:   3,
		Name: "海歌城",
		Location: &cityLocation{
			Abscissa: 120,
			Ordinate: 130,
		},
		Image: "",
		Level: NewLevel(2),
	}

	city4 := &city{
		Id:   4,
		Name: "陨星城",
		Location: &cityLocation{
			Abscissa: 120,
			Ordinate: 150,
		},
		Image: "",
		Level: NewLevel(3),
	}

	city1.addNeighbors([]*city{city2, city3})
	city2.addNeighbors([]*city{city1, city4})
	city3.addNeighbors([]*city{city1, city4})
	city4.addNeighbors([]*city{city2, city3})

	return []*city{city1, city2, city3, city4}
}
