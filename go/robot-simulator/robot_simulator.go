package robot

import (
	"strconv"
	"sync"
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
	Stop              Action = "Stopped"
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
	newPosition := newPositionOnMoving(*robot)
	robot.Pos = newPosition
}

func newPositionOnMoving(robot Step2Robot) Pos {
	switch robot.Dir {
	case N:
		robot.Pos.Northing++
	case E:
		robot.Pos.Easting++
	case S:
		robot.Pos.Northing--

	case W:
		robot.Pos.Easting--
	}
	return robot.Pos
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
		action <- Action3{
			Created,
			name,
		}
	forloop:
		for _, command := range script {
			// fmt.Println(command)
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
				break forloop
			}
		}

		// fmt.Println("Sending closure for ", name)
		action <- Action3{
			Stop,
			name,
		}
	}()
}

var lock sync.Map

func Room3(extent Rect, robots []Step3Robot, actions chan Action3, rep chan []Step3Robot, log chan string) {
	var robotsInRoom = map[string]int{}
	var currentPosition sync.Map
	numberOfRobotsLeftToStop := len(robots)

	// defer close(actions)
	// fmt.Println("Starting room")

	go func() {
	loop:
		for {
			select {
			case action := <-actions:
				robot_name := action.name
				robot_id, found := robotsInRoom[robot_name]
				switch action.Action {

				case Created:
					if found {
						log <- "Robot with same name already exists."
						continue
					}
					// fmt.Println("Creating robot", robot_name)
					if robot_name == "" {
						log <- "no name robot created"
						continue
					}
					for i, robot := range robots {
						if robot.Name == robot_name {
							robot_id = i
							if _, ok := currentPosition.Load(robot.Pos); ok {
								log <- "Robot exists at same position"
								continue loop
							}
							currentPosition.Store(robot.Pos, true)
						}
					}
					if !withinRoom(extent, robots[robot_id].Step2Robot) {
						log <- "Robot outside the area"
						continue
					}
					robotsInRoom[robot_name] = robot_id

				case Moved:
					if !movementWithinRoom(extent, robots[robot_id].Step2Robot) {
						log <- "wall bump"
						continue loop
					}

					newPosition := newPositionOnMoving(robots[robot_id].Step2Robot)
					if _, ok := currentPosition.Load(newPosition); ok {
						// fmt.Printf("%v", currentPosition)
						log <- "Collisions"
						continue
					}
					currentPosition.Delete(robots[robot_id].Step2Robot.Pos)
					Move(&robots[robot_id].Step2Robot)
					// fmt.Printf("%s moved to %v\n", robot_name, robots[robot_id].Step2Robot.Pos)
					currentPosition.Store(newPosition, true)

				case AntiClockWiseTurn:
					TurnLeft(&robots[robot_id].Step2Robot)

				case ClockWiseTurn:
					TurnRight(&robots[robot_id].Step2Robot)

				case Stop:
					// fmt.Println("Received Stop signal for ", action.name)
					numberOfRobotsLeftToStop--
					if numberOfRobotsLeftToStop == 0 {
						break loop
					}
					if numberOfRobotsLeftToStop < 0 {
						panic("something is wrong")
					}
				}
			}
		}
		rep <- robots
	}()
}
