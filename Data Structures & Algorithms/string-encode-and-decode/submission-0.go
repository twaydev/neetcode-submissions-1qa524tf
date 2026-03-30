
const defaultSeparator = "#"

type Solution struct {
	separator string
}

func (s *Solution) ensureSeparator() {
	if len(s.separator) == 0 {
		s.separator = defaultSeparator
	}
}
func (s *Solution) Encode(strs []string) string {
	s.ensureSeparator()

	stringBuilder := strings.Builder{}

	// count string
	for _, str := range strs {
		stringBuilder.WriteString(strconv.Itoa(len(str)))
		stringBuilder.WriteString(s.separator)
		stringBuilder.WriteString(str)
	}

	return stringBuilder.String()
}

func (s *Solution) Decode(encoded string) []string {
	s.ensureSeparator()

	result := []string{}
	i := 0
	for i < len(encoded) {
		currentPointer := i

		// found idx of separator
		for currentPointer < len(encoded) && encoded[currentPointer] != s.separator[0] {
			currentPointer++
		}

		//// get length of str
		length, _ := strconv.Atoi(encoded[i:currentPointer])

		// fmt.Printf("Decode: i: %d, currentPointer: %d, foundString: %s, length: %d\n", i, currentPointer, encoded[i:currentPointer], length)

		// move to next element (ignore separator)
		currentPointer++

		// append str to result
		result = append(result, encoded[currentPointer:currentPointer+length])

		// increase i
		i = currentPointer + length
	}

	return result
}