package queenattack

import (
	"fmt"
	"math"
	"strconv"
)

var columnMap = map[string]int{
	"a": 1,
	"b": 2,
	"c": 3,
	"d": 4,
	"e": 5,
	"f": 6,
	"g": 7,
	"h": 8,
}

func CanQueenAttack(whitePosition, blackPosition string) (bool, error) {

	if whitePosition == blackPosition {
		return false, fmt.Errorf("multiple pieces in same position")
	}

	blackPositionRow, ok1 := rowNumberFromPosition(blackPosition)
	blackPositionColumn, ok2 := columnNumberFromPosition(blackPosition)

	whitePositionRow, ok3 := rowNumberFromPosition(whitePosition)
	whitePositionColumn, ok4 := columnNumberFromPosition(whitePosition)

	if ok1 != nil || ok2 != nil || ok3 != nil || ok4 != nil {
		return false, fmt.Errorf("incorrect position string")
	}

	return _canQueenAttack(blackPositionRow, blackPositionColumn, whitePositionRow, whitePositionColumn), nil
}

func rowNumberFromPosition(position string) (int, error) {
	if len(position) != 2 {
		return 0, fmt.Errorf("incorrect position string")
	}
	val, err := strconv.Atoi(position[1:])

	if err != nil {
		return 0, fmt.Errorf("incorrect position string")
	}
	if val > 8 || val < 1 {
		return 0, fmt.Errorf("Incorrect row number")
	}
	return val, nil
}

func columnNumberFromPosition(position string) (int, error) {
	if len(position) != 2 {
		return 0, fmt.Errorf("incorrect position string")
	}

	result, ok := columnMap[position[:1]]
	if !ok {
		return 0, fmt.Errorf("unknown position string")
	}

	return result, nil
}

func _canQueenAttack(queen1Row, queen1Column, queen2Row, queen2Column int) bool {

	if (queen1Column == queen2Column) || (queen1Row == queen2Row) {
		return true
	}

	if math.Abs(float64(queen1Row)-float64(queen2Row)) == math.Abs(float64(queen1Column)-float64(queen2Column)) {
		return true
	}

	return false

}
