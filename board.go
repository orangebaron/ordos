package main

//all this setup allows the board to expand infinitely out in all directions
var board [2][2][][]building //the first 2 indices are for negative in X and Y
func bget(x, y int) *building {
	xIndex, yIndex := 1, 1
	if x < 0 {
		xIndex = 0
		x = -x
	}
	if y < 0 {
		yIndex = 0
		y = -y
	}
	return &(board[xIndex][yIndex][x][y])
}
func bset(x, y int, b *building) {
	xIndex, yIndex := 1, 1
	if x < 0 {
		xIndex = 0
		x = -x
	}
	if y < 0 {
		yIndex = 0
		y = -y
	}
	board[xIndex][yIndex][x][y] = *b
}
func badd(x, y int, b *building) { //bset but the capacity is increased if necessary
	xIndex, yIndex := 1, 1
	if x < 0 {
		xIndex = 0
		x = -x
	}
	if y < 0 {
		yIndex = 0
		y = -y
	}
	if cap(board[xIndex][yIndex]) < x {
		slice2 := make([][]building, len(board[xIndex][yIndex]), x+4) //leaves 3 spots extra open in that direction
		copy(slice2, board[xIndex][yIndex])
		board[xIndex][yIndex] = slice2
	}
	if cap(board[xIndex][yIndex][x]) < y {
		slice2 := make([]building, len(board[xIndex][yIndex][x]), x+4) //leaves 3 spots extra open in that direction
		copy(slice2, board[xIndex][yIndex][x])
		board[xIndex][yIndex][x] = slice2
	}
	board[xIndex][yIndex][x][y] = *b
}
