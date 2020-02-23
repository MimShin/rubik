package rubik

const size = 3

// Face -- a face of the Rubiks cube
type Face [size][size]byte

// Cube -- a Rubiks cube
type Cube struct {
	faces [6]Face
}

// Init -- initalizes the cube
func (cube *Cube) Init() {

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			cube.FaceByName("top")[i][j] = 'W'
			cube.FaceByName("bottom")[i][j] = 'Y'
			cube.FaceByName("left")[i][j] = 'G'
			cube.FaceByName("right")[i][j] = 'B'
			cube.FaceByName("front")[i][j] = 'R'
			cube.FaceByName("back")[i][j] = 'O'
		}
	}
}

func (cube *Cube) Size() int { return size }

func (cube *Cube) FaceByName(name string) *Face {
	switch name {
	case "top":
		return &cube.faces[0]
	case "bottom":
		return &cube.faces[1]
	case "left":
		return &cube.faces[2]
	case "right":
		return &cube.faces[3]
	case "front":
		return &cube.faces[4]
	case "back":
		return &cube.faces[5]
	}
	return nil
}

func (cube *Cube) Solved() bool {
	for i := 0; i < 6; i++ {
		face := cube.faces[i]
		color := face[0][0]
		// fmt.Printf("\ncolor is: %c: ", color)
		for r := 0; r < size; r++ {
			for c := 0; c < size; c++ {
				// fmt.Printf(" %c ", cube.faces[i][r][c])
				if face[r][c] != color {
					return false
				}
			}
		}
	}
	return true
}
