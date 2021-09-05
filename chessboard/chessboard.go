package chessboard

// Rank stores if a square is occupied by a piece
type Rank []bool

// Chessboard contains eight Ranks, accessed with values from 'A' to 'H'
type Chessboard map[byte]Rank

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank
func (cb Chessboard) CountInRank(rank byte) (ret int) {
	r := cb[rank]
	for _, square := range r {
		if square {
			ret++
		}
	}
	return
}

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file
func (cb Chessboard) CountInFile(file int) (ret int) {
	for _, rank := range cb {
		if file <= len(rank) && rank[file-1] {
			ret++
		}
	}
	return
}

// CountAll should count how many squares are present in the chessboard
func (cb Chessboard) CountAll() (ret int) {
	for _, r := range cb {
		ret += len(r)
	}
	return

}

// CountOccupied returns how many squares are occupied in the chessboard
func (cb Chessboard) CountOccupied() (ret int) {
	for r := range cb {
		ret += cb.CountInRank(r)
	}
	return
}
