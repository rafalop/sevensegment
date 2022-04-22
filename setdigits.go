package sevensegment

//Funcs for setting what is displayed

import "fmt"

// Set specific segments on for digit at pos 0 to 3 from right to left
// To determine which segments, we label the segments 0-6 starting from
// topmost horizontal, incrementing clockwise to 5, and 6 is the middle
// horizontal and specify which are on in an array
func (ss *SevenSegment) SetSegments(human_pos int, on [7]bool) {
	ss.Buffer[ss.HwPos(human_pos)] = ss.ArrayToSegments(on)
}

// Set digit at position from 0 to 3 from right to left
func (ss *SevenSegment) SetDigit(pos int, digit int) {
	ss.Buffer[ss.HwPos(pos)] = ss.IntToSs(digit)
}

// Set letter at position from 0 to 3 from right to left
func (ss *SevenSegment) SetAlpha(pos int, letter byte) {
	ss.Buffer[ss.HwPos(pos)] = ss.AlphaToSegments(letter)
}

// Set a number on the display up to 9999. A negative number will result
// in --- being displayed.
func (ss *SevenSegment) SetNum(number int) (err error) {
	if number > 9999 {
		err = fmt.Errorf("Error, number must be between 0 and 9999: %d", number)
		return err
	} else if number < 0 {
		for i := 0; i <= 2; i++ {
			ss.SetDigit(i, -1)
		}
	} else {
		num_a := NumToArray(number)
		for i, j := range num_a {
			// We call setdigit using array holding digits in reverse
			ss.SetDigit(i, j)
		}
	}
	return
}
