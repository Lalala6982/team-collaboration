package service

import (
	"campbe/database"
	"campbe/gateway"
	"campbe/model"
	"fmt"
	"log"
)

type Option struct {
	baseId         int
	transportation string
	distance       float64
	duration       int
	price          float64
}

func GetDispatchingOptions(from, to string) [3]Option {
	var options [3]Option
	// Prepare SQL query
	query := "SELECT id, base_address, num_of_robots, num_of_drones FROM bases"
	results, err := database.ReadFromDB(query)
	if err != nil {
		log.Fatal(err)
	}
	// Iterate over all the results
	var bases []model.Base
	for results.Next() {
		var base model.Base
		if err := results.Scan(&base.Id, &base.BaseAddress, &base.NumOfRobots, &base.NumOfDrones); err != nil {
			log.Fatal(err)
		}
		bases = append(bases, base)
	}
	fmt.Println(bases)
	// Recommended: robot route
	options[0].transportation = "robot"
	options[0].distance = 1e9
	for index, base := range bases {
		distance1, duration1 := gateway.GetRobotRoute(base.BaseAddress, from)
		distance2, _ := gateway.GetRobotRoute(to, base.BaseAddress)
		if distance1+distance2 < options[0].distance {
			options[0].baseId = index
			options[0].distance = distance1 + distance2
			options[0].duration = duration1
		}
	}
	distance1, duration1 := gateway.GetRobotRoute(from, to)
	options[0].distance += distance1
	options[0].duration += duration1
	// Fastest: drone route
	options[1].transportation = "drone"
	options[1].distance = 1e9
	for index, base := range bases {
		distance1, duration1 := gateway.GetDroneRoute(base.BaseAddress, from)
		distance2, _ := gateway.GetDroneRoute(to, base.BaseAddress)
		if distance1+distance2 < options[1].distance {
			options[1].baseId = index
			options[1].distance = distance1 + distance2
			options[1].duration = duration1
		}
	}
	distance1, duration1 = gateway.GetDroneRoute(from, to)
	options[1].distance += distance1
	options[1].duration += duration1
	// Cheapest: robot route but share with others
	options[2].baseId = options[0].baseId
	options[2].transportation = "robot"
	options[2].distance = options[0].distance * 2
	options[2].duration = options[0].duration + 20*60
	// return recommended options
	options[0].price = options[0].distance * model.ROBOT_CHARGE
	options[1].price = options[1].distance * model.DRONE_CHARGE
	options[2].price = options[2].distance * model.ROBOT_CHARGE / 3
	fmt.Println(options)
	return options
}
