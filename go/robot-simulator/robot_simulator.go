package robot

import (
	"fmt"
	"strconv"
)

// See defs.go for other definitions

// Step 1
const (
	N Dir = 0
	E Dir = 1
	S Dir = 2
	W Dir = 3
)

func Right() {
	Step1Robot.Dir = (Step1Robot.Dir + 1) % 4
}

func Left() {
	Step1Robot.Dir = (4 + (Step1Robot.Dir - 1)) % 4
}

func Advance() {
	switch Step1Robot.Dir {
	case N:
		Step1Robot.Y++
	case E:
		Step1Robot.X++

	case S:
		Step1Robot.Y--

	case W:
		Step1Robot.X--
	default:
		panic("unknown directions")
	}

}

func (d Dir) String() string {
	return strconv.Itoa(int(d))
}

// Step 2
// Define Action type here.

type Action string

const (
	Moved             Action = "Moved"
	ClockWiseTurn     Action = "Turned Right"
	AntiClockWiseTurn Action = "Turned Left"
	Created           Action = "Created"
)

func StartRobot(commands chan Command, action chan Action) {

	go func() {
		for command := range commands {
			switch command {
			case 'A':
				action <- Moved
			case 'R':
				action <- ClockWiseTurn
			case 'L':
				action <- AntiClockWiseTurn
			default:
				panic("something went wrong")
			}
		}
		close(action)
	}()

}

func Move(robot *Step2Robot) {
	switch robot.Dir {
	case N:
		robot.Pos.Northing++
	case E:
		robot.Pos.Easting++

	case S:
		robot.Pos.Northing--

	case W:
		robot.Pos.Easting--
	default:
		panic("unknown directions")
	}

}

func Room(extent Rect, robot Step2Robot, actions chan Action, report chan Step2Robot) {
	go func() {
		for action := range actions {
			switch action {
			case Moved:
				if movementWithinRoom(extent, robot) {
					Move(&robot)
				}
			case AntiClockWiseTurn:
				TurnLeft(&robot)

			case ClockWiseTurn:
				TurnRight(&robot)
			}
		}
		report <- robot
	}()

}
func withinRoom(extent Rect, robot Step2Robot) bool {
	return (robot.Northing <= extent.Max.Northing) && (robot.Easting <= extent.Max.Easting) &&
		(robot.Northing >= extent.Min.Northing) && (robot.Easting >= extent.Min.Easting)
}

func movementWithinRoom(extent Rect, robot Step2Robot) bool {
	switch robot.Dir {
	case N:
		robot.Northing++
	case E:
		robot.Easting++
	case S:
		robot.Northing--
	case W:
		robot.Easting--
	}
	return withinRoom(extent, robot)
}

func TurnRight(robot *Step2Robot) {
	robot.Dir = (robot.Dir + 1) % 4
}

func TurnLeft(robot *Step2Robot) {
	robot.Dir = (4 + (robot.Dir - 1)) % 4
}

// Step 3
// Define Action3 type here.

type Action3 struct {
	Action
	name string
}

func StartRobot3(name, script string, action chan Action3, log chan string) {
	go func() {
		defer close(action)

		action <- Action3{
			Created,
			name,
		}
		for _, command := range script {
			fmt.Println(command)
			switch command {
			case 'A':
				action <- Action3{
					Moved,
					name,
				}
			case 'R':
				action <- Action3{
					ClockWiseTurn,
					name,
				}
			case 'L':
				action <- Action3{
					AntiClockWiseTurn,
					name,
				}
			default:
				log <- "Bad command"
			}
		}
	}()
}

func Room3(extent Rect, robots []Step3Robot, actions chan Action3, rep chan []Step3Robot, log chan string) {
	var robotsInRoom = map[string]int{}
	var currentPosition = map[Pos]any{}

	fmt.Println("Starting room")
	go func() {
		for action := range actions {
			robot_name := action.name
			robot_id := robotsInRoom[robot_name]
			switch action.Action {

			case Created:
				fmt.Println("Creating robot", robot_name)
				if !withinRoom(extent, robots[robot_id].Step2Robot) {
					log <- "Robot outside the area"
				}
				if robot_name == "" {
					log <- "no name robot craeted"
				}
				for i, robot := range robots {
					if _, ok := robotsInRoom[robot.Name]; ok {
						log <- "Robot with the name already exists."
					}
					robotsInRoom[robot.Name] = i
					currentPosition[robot.Pos] = true
				}

			case Moved:
				if !movementWithinRoom(extent, robots[robot_id].Step2Robot) {
					log <- "wall bump"
					continue
				}

				delete(currentPosition, robots[robot_id].Step2Robot.Pos)
				Move(&robots[robot_id].Step2Robot)
				if _, ok := currentPosition[robots[robot_id].Step2Robot.Pos]; ok {
					log <- "Collisions"
					continue
				}
				currentPosition[robots[robot_id].Step2Robot.Pos] = true

			case AntiClockWiseTurn:
				TurnLeft(&robots[robot_id].Step2Robot)

			case ClockWiseTurn:
				TurnRight(&robots[robot_id].Step2Robot)
			}
		}
		rep <- robots
	}()
}
