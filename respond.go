package main

import (
	"fmt"
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// func respondToUser() {
// 	switch {
// 	case rl.IsKeyReleased(rl.KeyQ):
// 		os.Exit(0)
// 	case rl.IsKeyPressed(rl.KeyOne) && rl.IsKeyPressed(rl.KeyLeftShift):
// 		randomHands = 1;
// 		hands = randomRectangularHands(randomHands)
// 	case rl.IsKeyPressed(rl.KeyNine) && rl.IsKeyPressed(rl.KeyLeftShift):
// 		randomHands = 9;
// 		hands = randomRectangularHands(randomHands)
// 	case rl.IsKeyReleased(rl.KeyH):
// 		hands = randomRectangularHands(randomHands)
// 	case rl.IsKeyReleased(rl.KeySpace):
// 		mutField[0].randomize()
// 		mutField[1].randomize()
// 	case rl.IsKeyReleased(rl.KeyRight):
// 		adjustSpeed(5)
// 	case rl.IsKeyReleased(rl.KeyLeft):
// 		adjustSpeed(-5)
// 	case rl.IsKeyReleased(rl.KeyW):
// 		writeHandsToFile()
// 	case rl.IsKeyReleased(rl.KeyOne):
// 		changeHand(0)
// 	case rl.IsKeyReleased(rl.KeyTwo):
// 		changeHand(1)
// 	case rl.IsKeyReleased(rl.KeyThree):
// 		changeHand(2)
// 	case rl.IsKeyReleased(rl.KeyFour):
// 		changeHand(3)
// 	case rl.IsKeyReleased(rl.KeyFive):
// 		changeHand(4)
// 	case rl.IsKeyReleased(rl.KeySix):
// 		changeHand(5)
// 	case rl.IsKeyReleased(rl.KeySeven):
// 		changeHand(6)
// 	case rl.IsKeyReleased(rl.KeyEight):
// 		changeHand(7)
// 	case rl.IsKeyReleased(rl.KeyNine):
// 		changeHand(8)
// 	case rl.IsKeyReleased(rl.KeyZero):
// 		changeHand(9)
// 	case rl.IsKeyPressed(rl.KeyNine) && rl.IsKeyPressed(rl.KeyLeftShift):
// 		randomHands = 0;
// 		hands = randomRectangularHands(randomHands)
// 	}
// }

func respondToUser() {
	switch {
	case rl.IsKeyReleased(rl.KeyQ):
		os.Exit(0)
	case rl.IsKeyReleased(rl.KeyW):
		writeHandsToFile()
	case rl.IsKeyReleased(rl.KeyS):
		mutField[0].randomize()
		mutField[1].randomize()
	case rl.IsKeyReleased(rl.KeyF):
		hands = randomRectangularHands(randomHands)
	case rl.IsKeyReleased(rl.KeyRight):
		adjustSpeed(5)
	case rl.IsKeyReleased(rl.KeyLeft):
		adjustSpeed(-5)
	case rl.IsKeyDown(rl.KeyRightShift):
		switch {
		case rl.IsKeyPressed(rl.KeyOne):
			randomHands = 1;
			hands = randomRectangularHands(randomHands)
		case rl.IsKeyPressed(rl.KeyTwo):
			randomHands = 2;
			hands = randomRectangularHands(randomHands)
		case rl.IsKeyPressed(rl.KeyThree):
			randomHands = 3;
			hands = randomRectangularHands(randomHands)
		case rl.IsKeyPressed(rl.KeyFour):
			randomHands = 4;
			hands = randomRectangularHands(randomHands)
		case rl.IsKeyPressed(rl.KeyFive):
			randomHands = 5;
			hands = randomRectangularHands(randomHands)
		case rl.IsKeyPressed(rl.KeySix):
			randomHands = 6;
			hands = randomRectangularHands(randomHands)
		case rl.IsKeyPressed(rl.KeySeven):
			randomHands = 7;
			hands = randomRectangularHands(randomHands)
		case rl.IsKeyPressed(rl.KeyEight):
			randomHands = 8;
			hands = randomRectangularHands(randomHands)
		case rl.IsKeyPressed(rl.KeyNine):
			randomHands = 9;
			hands = randomRectangularHands(randomHands)
		}
	case rl.IsKeyDown(rl.KeyLeftShift):
		switch {
		case rl.IsKeyReleased(rl.KeyRight):
			adjustSpeed(5)
		case rl.IsKeyReleased(rl.KeyLeft):
			adjustSpeed(-5)
		case rl.IsKeyReleased(rl.KeyW):
			writeHandsToFile()
		case rl.IsKeyReleased(rl.KeyOne):
			changeHand(0)
		case rl.IsKeyReleased(rl.KeyTwo):
			changeHand(1)
		case rl.IsKeyReleased(rl.KeyThree):
			changeHand(2)
		case rl.IsKeyReleased(rl.KeyFour):
			changeHand(3)
		case rl.IsKeyReleased(rl.KeyFive):
			changeHand(4)
		case rl.IsKeyReleased(rl.KeySix):
			changeHand(5)
		case rl.IsKeyReleased(rl.KeySeven):
			changeHand(6)
		case rl.IsKeyReleased(rl.KeyEight):
			changeHand(7)
		case rl.IsKeyReleased(rl.KeyNine):
			changeHand(8)
		case rl.IsKeyReleased(rl.KeyZero):
			changeHand(9)
		}
	case rl.IsKeyDown(rl.KeyLeftControl):
		switch {
		case rl.IsKeyPressed(rl.KeyOne):
			rows = 25
			cols = 50
			pixelHeight = 40
			pixelWidth = 20
			canvas  = grid[Color]()
			mutField = [2]Field{grid[Real](), grid[Real]()}
			mutField[0].randomize()
			mutField[1].randomize()
		case rl.IsKeyPressed(rl.KeyTwo):
			rows = 50
			cols = 100
			pixelHeight = 20
			pixelWidth = 10
			canvas  = grid[Color]()
			mutField = [2]Field{grid[Real](), grid[Real]()}
			mutField[0].randomize()
			mutField[1].randomize()
		case rl.IsKeyPressed(rl.KeyThree):
			rows = 100
			cols = 200
			pixelHeight = 10
			pixelWidth = 5
			canvas  = grid[Color]()
			mutField = [2]Field{grid[Real](), grid[Real]()}
			mutField[0].randomize()
			mutField[1].randomize()
		case rl.IsKeyPressed(rl.KeyFour):
			rows = 250
			cols = 250
			pixelHeight = 4
			pixelWidth = 4
			canvas  = grid[Color]()
			mutField = [2]Field{grid[Real](), grid[Real]()}
			mutField[0].randomize()
			mutField[1].randomize()
		}
		// case rl.IsKeyPressed(rl.KeyTwo):
		// 	randomHands = 2;
		// 	hands = randomRectangularHands(randomHands)
		// case rl.IsKeyPressed(rl.KeyThree):
		// 	randomHands = 3;
		// 	hands = randomRectangularHands(randomHands)
		// case rl.IsKeyPressed(rl.KeyFour):
		// 	randomHands = 4;
		// 	hands = randomRectangularHands(randomHands)
		// case rl.IsKeyPressed(rl.KeyFive):
		// 	randomHands = 5;
		// 	hands = randomRectangularHands(randomHands)
		// case rl.IsKeyPressed(rl.KeySix):
		// 	randomHands = 6;
		// 	hands = randomRectangularHands(randomHands)
		// case rl.IsKeyPressed(rl.KeySeven):
		// 	randomHands = 7;
		// 	hands = randomRectangularHands(randomHands)
		// case rl.IsKeyPressed(rl.KeyEight):
		// 	randomHands = 8;
		// 	hands = randomRectangularHands(randomHands)
		// case rl.IsKeyPressed(rl.KeyNine):
		// 	randomHands = 9;
		// 	hands = randomRectangularHands(randomHands)
		
	}	
}

func adjustSpeed(i int) {
	fps += i
	if fps <= 0 {
		fps = 1
	}
	rl.SetTargetFPS(int32(fps))
	fmt.Print("fps = ", fps)
}
