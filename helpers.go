package sevensegment

// Funcs that help convert to set up the correct bytes used for commands or
// referencing HT16K33 memory/digits

// Translate human pos on display (0 - 4) to corresponding rows
// in HT16K33 mem that reference a digit or colons, pos 5 is colons.
func (ss *SevenSegment) HwPos(human_pos int) int {
	hw_positions := [5]int{9, 7, 3, 1, 5}
	return hw_positions[human_pos]
}

// Return 2 byte value representing which segments are on for supplied array of 'on' segments
func (ss *SevenSegment) ArrayToSegments(segments_on [7]bool) uint16 {
	tot := uint16(0)
	for index, on := range segments_on {
		if on {
			tot += 1 << index
		}
	}
	return tot
}

// Return 2 byte value which segments should be on for a supplied letter
func (ss *SevenSegment) AlphaToSegments(letter byte) uint16 {
	var alphamap map[string]uint16
	alphamap = make(map[string]uint16)
	alphamap["A"] = ss.ArrayToSegments([7]bool{true, true, true, false, true, true, true})
	alphamap["b"] = ss.ArrayToSegments([7]bool{false, false, true, true, true, true, true})
	alphamap["c"] = ss.ArrayToSegments([7]bool{false, false, false, true, true, false, true})
	alphamap["d"] = ss.ArrayToSegments([7]bool{false, true, true, true, true, false, true})
	alphamap["e"] = ss.ArrayToSegments([7]bool{true, true, false, true, true, true, true})
	alphamap["f"] = ss.ArrayToSegments([7]bool{true, false, false, false, true, true, true})
	alphamap["g"] = ss.ArrayToSegments([7]bool{true, true, true, true, false, true, true})
	alphamap["h"] = ss.ArrayToSegments([7]bool{false, false, true, false, true, true, true})
	alphamap["I"] = ss.ArrayToSegments([7]bool{false, false, false, false, true, true, false})
	alphamap["o"] = ss.ArrayToSegments([7]bool{false, false, true, true, true, false, true})
	alphamap["u"] = ss.ArrayToSegments([7]bool{false, false, true, true, true, false, false})
	alphamap["S"] = ss.ArrayToSegments([7]bool{true, false, true, true, false, true, true})

	return alphamap[string(letter)]
}

// Return array (reverse order) of digits for a given 1-4 digit number
func NumToArray(number int) []int {
	var num_a []int
	if number == 0 {
		num_a = append(num_a, 0)
	} else {
		num_tmp := number
		for num_tmp != 0 {
			// This would put numbers in correct order as an array
			//num_a = append([]int{num_tmp % 10}, num_a...)
			// Instead, we put them in reverse to assist
			// calling the SetDigit function
			num_a = append(num_a, num_tmp%10)
			num_tmp = num_tmp / 10
		}
	}
	for len(num_a) < 4 {
		num_a = append(num_a, -10000)
	}
	return num_a
}

// Returns 2 byt val representing the data required to turn segments on for supplied 'digit' similar to numbertable at:
// https://github.com/adafruit/Adafruit_LED_Backpack/blob/master/Adafruit_LEDBackpack.cpp
func (ss *SevenSegment) IntToSs(digit int) uint16 {
	ss_vals := []uint16{63, 6, 91, 79, 102, 109, 125, 7, 127, 103}
	if digit == -10000 {
		return 0
	} else if digit < 0 {
		return 64
	}
	return ss_vals[digit]
}
