package service

import (
	"campbe/gateway"
	"campbe/model"
)

type Option struct {
	dispatcher     int
	transportation string
	distance       float64
	duration       int
	price          float64
}

func getDispatchingOptions(from, to string) [3]Option {
	var options [3]Option
	// Recommended: robot route
	options[0].transportation = "robot"
	options[0].distance = 1e9
	for index, dispatcher := range dispatchers {
		distance1, duration1 := gateway.GetRobotRoute(dispatcher, from)
		distance2, _ := gateway.GetRobotRoute(to, dispatcher)
		if distance1+distance2 < options[0].distance {
			options[0].dispatcher = index
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
	for index, dispatcher := range dispatchers {
		distance1, duration1 := gateway.GetDroneRoute(dispatcher, from)
		distance2, _ := gateway.GetDroneRoute(to, dispatcher)
		if distance1+distance2 < options[1].distance {
			options[1].dispatcher = index
			options[1].distance = distance1 + distance2
			options[1].duration = duration1
		}
	}
	distance1, duration1 = gateway.GetDroneRoute(from, to)
	options[1].distance += distance1
	options[1].duration += duration1
	// Cheapest: robot route but share with others
	options[2].dispatcher = options[0].dispatcher
	options[2].transportation = "robot"
	options[2].distance = options[0].distance * 2
	options[2].duration = options[0].duration + 20*60
	// return recommended options
	options[0].price = options[0].distance * model.ROBOT_CHARGE
	options[1].price = options[1].distance * model.DRONE_CHARGE
	options[2].price = options[2].distance * model.ROBOT_CHARGE / 3
	return options
}
