package strand

func ToRNA(dna string) (transcribed string) {
	for _, c := range dna {
		switch c {
		case 'C':
			c = 'G'
		case 'G':
			c = 'C'
		case 'T':
			c = 'A'
		case 'A':
			c = 'U'
		}
		transcribed += string(c)
	}
	return
}
