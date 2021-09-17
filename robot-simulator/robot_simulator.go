package robot

import "fmt"

const (
	N Dir = iota
	E
	S
	W
)

func (dir Dir) String() string {
	switch dir {
	case N:
		return "North"
	case E:
		return "East"
	case S:
		return "South"
	case W:
		return "West"
	default:
		return ""
	}

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

	}
}

func Advance2(robot Step2Robot) Step2Robot {
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
	return robot
}

func Right() {
	switch Step1Robot.Dir {
	case N:
		Step1Robot.Dir = E
	case E:
		Step1Robot.Dir = S
	case S:
		Step1Robot.Dir = W
	case W:
		Step1Robot.Dir = N

	}
}

func Right2(robot Step2Robot) Step2Robot {
	switch robot.Dir {
	case N:
		robot.Dir = E
	case E:
		robot.Dir = S
	case S:
		robot.Dir = W
	case W:
		robot.Dir = N
	}
	return robot
}

func Left() {
	switch Step1Robot.Dir {
	case N:
		Step1Robot.Dir = W
	case E:
		Step1Robot.Dir = N
	case S:
		Step1Robot.Dir = E
	case W:
		Step1Robot.Dir = S
	}
}

func Left2(robot Step2Robot) Step2Robot {
	switch robot.Dir {
	case N:
		robot.Dir = W
	case E:
		robot.Dir = N
	case S:
		robot.Dir = E
	case W:
		robot.Dir = S
	}
	return robot
}

type Action func(Step2Robot) Step2Robot

func StartRobot(commands chan Command, actions chan Action) {
	for cmd := range commands {
		switch cmd {
		case 'A':
			actions <- Advance2
		case 'L':
			actions <- Left2
		case 'R':
			actions <- Right2
		}
	}
	close(actions)
}

func (r Step2Robot) String() string {
	return fmt.Sprintf("robot {Dir=%s, pos=%v}", r.Dir, r.Pos)
}

func Room(limits Rect, robot Step2Robot, actions chan Action, report chan Step2Robot) {
	robot.Pos = limits.Min
	robot.Dir = N
	for action := range actions {
		fmt.Printf("action %v\n", action)
		fmt.Printf("robot in = %+v\n", robot)
		robot = checkPos(limits, action(robot))
		fmt.Printf("robot out = %+v\n\n", robot)
	}
	report <- robot
}

func checkPos(limits Rect, robot Step2Robot) Step2Robot {
	if robot.Pos.Northing < limits.Min.Northing {
		robot.Pos.Northing = limits.Min.Northing
	} else if robot.Pos.Northing > limits.Max.Northing {
		robot.Pos.Northing = limits.Max.Northing
	}
	if robot.Pos.Easting < limits.Min.Easting {
		robot.Pos.Easting = limits.Min.Easting
	} else if robot.Pos.Easting > limits.Max.Easting {
		robot.Pos.Easting = limits.Max.Easting
	}
	return robot
}
