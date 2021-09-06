package protein

type ProteinError string

func (e ProteinError) Error() string {
	return string(e)
}

const ErrStop ProteinError = "found stop"
const ErrInvalidBase ProteinError = "invalid base"

func FromCodon(codon string) (string, error) {
	switch codon {
	case "AUG":
		return "Methionine", nil
	case "UUU":
		return "Phenylalanine", nil
	case "UUC":
		return "Phenylalanine", nil
	case "UUA":
		return "Leucine", nil
	case "UUG":
		return "Leucine", nil
	case "UCU":
		return "Serine", nil
	case "UCC":
		return "Serine", nil
	case "UCA":
		return "Serine", nil
	case "UCG":
		return "Serine", nil
	case "UAU":
		return "Tyrosine", nil
	case "UAC":
		return "Tyrosine", nil
	case "UGU":
		return "Cysteine", nil
	case "UGC":
		return "Cysteine", nil
	case "UGG":
		return "Tryptophan", nil
	case "UAA":
		return "", ErrStop
	case "UAG":
		return "", ErrStop
	case "UGA":
		return "", ErrStop
	default:
		return "", ErrInvalidBase
	}
}

func FromRNA(rna string) ([]string, error) {
	codons := getCodons(rna)
	proteins := []string{}
	for _, codon := range codons {
		switch protein, err := FromCodon(codon); err {
		case ErrInvalidBase:
			return proteins, err
		case ErrStop:
			return proteins, nil
		default:
			proteins = append(proteins, protein)

		}
	}
	return proteins, nil
}

func getCodons(rna string) []string {
	codons := make([]string, 0, len(rna)/3+1)
	cl := 0
	cs := 0
	for i := range rna {
		if cl == 3 {
			codons = append(codons, rna[cs:i])
			cl = 0
			cs = i
		}
		cl++
	}
	codons = append(codons, rna[cs:])
	return codons
}
