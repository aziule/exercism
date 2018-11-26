package robot

import "fmt"

type Action int

const (
	GoLeft     Action = 0
	GoRight    Action = 1
	GoStraight Action = 2
)

func StartRobot(com <-chan Command, act chan<- Action) {
	for c := range com {
		switch string(c) {
		case "L":
			act <- GoLeft
		case "R":
			act <- GoRight
		case "A":
			act <- GoStraight
		default:
			fmt.Println(c)
		}
	}

	close(act)
}

func Room(extent Rect, robot Step2Robot, act chan Action, rep chan Step2Robot) {
	for a := range act {
		switch a {
		case GoLeft:
			robot.Dir = newDir(1, robot.Dir)
		case GoRight:
			robot.Dir = newDir(-1, robot.Dir)
		case GoStraight:
			switch robot.Dir {
			case N:
				if robot.Pos.Northing == extent.Max.Northing {
					break
				}
				robot.Pos.Northing++
			case W:
				if robot.Pos.Easting == extent.Min.Easting {
					break
				}
				robot.Pos.Easting--
			case S:
				if robot.Pos.Northing == extent.Min.Northing {
					break
				}
				robot.Pos.Northing--
			case E:
				if robot.Pos.Easting == extent.Max.Easting {
					break
				}
				robot.Pos.Easting++
			}
		}
	}

	rep <- robot
}
