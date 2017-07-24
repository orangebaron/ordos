package main

import "./buildings"

//all this setup allows the board to expand infinitely out in all directions
var board [2][2][][]buildings.Building //the first 2 indices are for negative in X and Y
func bget(x, y int) *buildings.Building {
	xIndex, yIndex := 0, 0 //0 means positive, 1 means negative
	if x < 0 {
		xIndex = 1
		x = -x
	}
	if y < 0 {
		yIndex = 1
		y = -y
	}
	return &(board[xIndex][yIndex][x][y])
}
func bset(x, y int, b *buildings.Building) {
	xIndex, yIndex := 0, 0
	if x < 0 {
		xIndex = 1
		x = -x
	}
	if y < 0 {
		yIndex = 1
		y = -y
	}
	board[xIndex][yIndex][x][y] = *b
}
func badd(x, y int, b *buildings.Building) { //bset but the capacity is increased if necessary
	xIndex, yIndex := 0, 0
	if x < 0 {
		xIndex = 1
		x = -x
	}
	if y < 0 {
		yIndex = 1
		y = -y
	}
	if cap(board[xIndex][yIndex]) < x {
		slice2 := make([][]buildings.Building, len(board[xIndex][yIndex]), x+4) //leaves 3 spots extra open in that direction
		copy(slice2, board[xIndex][yIndex])
		board[xIndex][yIndex] = slice2
	}
	if cap(board[xIndex][yIndex][x]) < y {
		slice2 := make([]buildings.Building, len(board[xIndex][yIndex][x]), x+4) //leaves 3 spots extra open in that direction
		copy(slice2, board[xIndex][yIndex][x])
		board[xIndex][yIndex][x] = slice2
	}
	board[xIndex][yIndex][x][y] = *b
}
func bForEach(f func(*buildings.Building, int, int)) {
	for xIndex := 0; xIndex <= 1; xIndex++ {
		for yIndex := 0; yIndex <= 1; yIndex++ {
			for x := xIndex; x < len(board[xIndex][yIndex]); x++ {
				for y := yIndex; y < len(board[xIndex][yIndex][x]); y++ {
					f(&board[xIndex][yIndex][x][y], x, y)
				}
			}
		}
	}
}
