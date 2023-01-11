package strand

var dtr = map[rune]rune{'G': 'C', 'C': 'G', 'T': 'A', 'A': 'U'}

// G -> C
// C -> G
// T -> A
// A -> U

func ToRNA(dna string) string {
	var rna string
	for _, v := range dna {
		rna += string(dtr[v])
	}
	return rna
}
