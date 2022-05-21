package chessboard

// Declare a type named Rank which stores if a square is occupied by a piece - this will be a slice of bools
type Rank = []bool

// Declare a type named Chessboard which contains a map of eight Ranks, accessed with keys from "A" to "H"
type Chessboard = map[string]Rank

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank
func CountInRank(cb Chessboard, rank string) int {
	occupied := 0
	for _, square := range cb[rank] {
		if square {
			occupied++
		}
	}
	return occupied
}

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file
func CountInFile(cb Chessboard, file int) int {
	if file < 0 || file > 8 {
		return 0
	}

	file -= 1
	occupied := 0
	for _, rank := range cb {
		if rank[file] {
			occupied++
		}
	}
	return occupied
}

// CountAll should count how many squares are present in the chessboard
func CountAll(cb Chessboard) int {
	return len(cb) * len(cb["A"])
}

// CountOccupied returns how many squares are occupied in the chessboard
func CountOccupied(cb Chessboard) int {
	totalOccupied := 0
	for rank := range cb {
		totalOccupied += CountInRank(cb, rank)
	}

	return totalOccupied
}
