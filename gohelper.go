package gohelper

func ArrayToInt(input []byte) (int) {
	/* This function converts a byte array to an unsigned integer *
	 * a maximum of 64 bits wide, endianness is preserved
	 */

	var returnint int = 0
	for x := len(input)-1; x >=0; x-- {
		returnint += int(input[x]) << (8*uint(x))
	}
	return returnint
}

func IntToBEArray(input int) ([]byte) {
	/* This function takes an integer and converts it to a *
	 * Big Endian byte array, irrespective of integer size *
	 */
	
	var counter int = 1
	var i int = input

	/* Get the actual required size of the integer in bytes */
	for {
		i = i >> 8
		if i != 0 {
			counter += 1
		} else {
			break
		}
	}

	/* Create the byte array */
	var barray []byte = make([]byte, counter)

	for {
		i = (8*(counter-1))
		barray[counter-1] = byte((input >> uint(i)) & 255)
		counter -= 1
		if counter == 1 {
			barray[0] = byte(input&255)
			break
		}
	}
	return barray
}

