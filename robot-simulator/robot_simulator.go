package robot

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

func Room(limits Rect, robot Step2Robot, actions chan Action, report chan Step2Robot) {
	robot.Pos = limits.Min
	robot.Dir = N
	for action := range actions {
		robot = action(robot)
		robot.Pos, _ = checkPos(limits, robot.Pos)
	}
	report <- robot
}

type Action3 struct {
	Stop  bool
	Robot string
	Action
}

func StartRobot3(robotName string, script string, actions chan Action3, log chan string) {
	for _, cmd := range script {
		switch cmd {
		case 'A':
			actions <- Action3{Robot: robotName, Action: Advance2}
		case 'L':
			actions <- Action3{Robot: robotName, Action: Left2}
		case 'R':
			actions <- Action3{Robot: robotName, Action: Right2}
		default:
			log <- "An undefined command in a script"
			actions <- Action3{Stop: true, Robot: robotName}
			return
		}
	}
	actions <- Action3{Stop: true, Robot: robotName}
}

func Room3(limits Rect, robots []Step3Robot, actions chan Action3, report chan []Step3Robot, log chan string) {
	robotMap := map[string]Step2Robot{}
	for _, robot := range robots {
		if robot.Name == "" {
			log <- "Robot without name"
			report <- robots
			return
		}
		if _, bump := checkPos(limits, robot.Step2Robot.Pos); bump {
			log <- "A robot placed outside of the room"
			report <- robots
			return
		}
		if _, ok := robotMap[robot.Name]; ok {
			log <- "Duplicated robot name"
			report <- robots
			return
		}
		if checkCollide(robot.Name, robot.Step2Robot, robotMap) {
			log <- "Robots placed at the same place"
			report <- robots
			return
		}
		robotMap[robot.Name] = robot.Step2Robot
	}

	robotCount := len(robotMap)
	robots = []Step3Robot{}
	for len(robots) < robotCount {
		action := <-actions
		robot, ok := robotMap[action.Robot]
		if !ok {
			log <- "An action from an unknown robot"
			break
		}
		if action.Stop {
			robots = append(robots, Step3Robot{Name: action.Robot, Step2Robot: robot})
			continue
		}
		bump := false
		pos := robot.Pos
		robot = action.Action(robot)
		if checkCollide(action.Robot, robot, robotMap) {
			log <- "A robot attempting to advance into another robot"
		} else {
			pos, bump = checkPos(limits, robot.Pos)
			if bump {
				log <- "A robot attempting to advance into a wall"
			}
		}
		robot.Pos = pos
		robotMap[action.Robot] = robot
	}
	report <- robots
}

func checkCollide(robotName string, robot Step2Robot, robots map[string]Step2Robot) bool {
	for name, otherRobot := range robots {
		if name != robotName && robot.Pos == otherRobot.Pos {
			return true
		}
	}
	return false
}

func checkPos(limits Rect, pos Pos) (Pos, bool) {
	bump := false
	if pos.Northing < limits.Min.Northing {
		pos.Northing = limits.Min.Northing
		bump = true
	} else if pos.Northing > limits.Max.Northing {
		pos.Northing = limits.Max.Northing
		bump = true
	}
	if pos.Easting < limits.Min.Easting {
		pos.Easting = limits.Min.Easting
		bump = true
	} else if pos.Easting > limits.Max.Easting {
		pos.Easting = limits.Max.Easting
		bump = true
	}
	return pos, bump
}
