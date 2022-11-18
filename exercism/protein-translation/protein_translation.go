package protein

import "errors"

var protein = map[string]string{"AUG": "Methionine", "UUU": "Phenylalanine", "UUC": "Phenylalanine", "UUA": "Leucine", "UUG": "Leucine", "UCU": "Serine", "UCC": "Serine", "UCA": "Serine", "UCG": "Serine", "UAU": "Tyrosine", "UAC": "Tyrosine", "UGU": "Cysteine", "UGC": "Cysteine", "UGG": "Tryptophan"}
var stops = map[string]string{"UAA": "STOP", "UAG": "STOP", "UGA": "STOP"}

var ErrInvalidBase = errors.New("ErrInvalidBase")
var ErrStop = errors.New("ErrStop")

func Chunks(s string, chunkSize int) []string {
	if len(s) == 0 {
		return nil
	}
	if chunkSize >= len(s) {
		return []string{s}
	}
	var chunks []string = make([]string, 0, (len(s)-1)/chunkSize+1)
	currentLen := 0
	currentStart := 0
	for i := range s {
		if currentLen == chunkSize {
			chunks = append(chunks, s[currentStart:i])
			currentLen = 0
			currentStart = i
		}
		currentLen++
	}
	chunks = append(chunks, s[currentStart:])
	return chunks
}

func FromRNA(rna string) ([]string, error) {
	// if len(rna)%3 != 0 {
	// 	return nil, ErrInvalidBase
	// }
	codons := Chunks(rna, 3)
	var res []string
	for _, v := range codons {
		codon, err := FromCodon(v)

		if err == ErrStop {
			break
		}
		if err != nil {
			return nil, err
		}
		res = append(res, codon)

	}
	return res, nil
}

func FromCodon(codon string) (string, error) {
	if _, ok := stops[codon]; ok {
		return "", ErrStop
	} else if aminoacid, ok := protein[codon]; ok {
		return aminoacid, nil
	}
	return "", ErrInvalidBase

}
