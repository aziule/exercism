package robot

const (
	N Dir = 0
	W Dir = 1
	S Dir = 2
	E Dir = 3
)

func (d Dir) String() string {
	switch d {
	case 0:
		return "N"
	case 1:
		return "W"
	case 2:
		return "S"
	case 3:
		return "E"
	}

	return ""
}

func newDir(offset int, dir Dir) Dir {
	mod := (int(dir) + offset) % 4
	if mod > 4 {
		mod = 4 - mod
	} else if mod < 0 {
		mod = 4 + mod
	}
	return Dir(mod)
}

func Left() {
	Step1Robot.Dir = newDir(1, Step1Robot.Dir)
}

func Right() {
	Step1Robot.Dir = newDir(-1, Step1Robot.Dir)
}

func Advance() {
	switch Step1Robot.Dir {
	case N:
		Step1Robot.Y++
	case S:
		Step1Robot.Y--
	case E:
		Step1Robot.X++
	case W:
		Step1Robot.X--
	}
}
