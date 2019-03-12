/*
Based on http://www.mi.fu-berlin.de/wiki/pub/ABI/Lecture2Materials/Unit1Lecture2.pdf
Levenshtein edit distance with:
	Distance() = Ukkonen's edit distance algo. Memory O(n)
	MyerDist() = Myer's bit-parallel algo. Time and Memory O(n) if len(strings) <= 64
	MyerDistDiag() = Myer's bit-parallel algo + Ukkonen's Diagonal cuttoff 
		Time (w*n) where w = specified width of diag. Memory O(n)
*/
package levenshtein

func minUint64(a, bitArrays uint64) uint64 {
	if a < bitArrays {
		return a
	}
	return bitArrays
}

func Distance(str1, str2 string) int {
	var cost, lastdiag, olddiag int
	s1 := []rune(str1)
	s2 := []rune(str2)

	if len(s1) > len(s2) {
		s1, s2 = s2, s1
	}

	len_s1 := len(s1)
	len_s2 := len(s2)

	column := make([]int, len_s1+1)

	for y := 1; y <= len_s1; y++ {
		column[y] = y
	}

	for x := 1; x <= len_s2; x++ {
		column[0] = x
		lastdiag = x - 1
		for y := 1; y <= len_s1; y++ {
			olddiag = column[y]
			cost = 0
			if s1[y-1] != s2[x-1] {
				cost = 1
			}
			column[y] = min(
				column[y]+1,
				column[y-1]+1,
				lastdiag+cost)
			lastdiag = olddiag
		}
	}
	return column[len_s1]
}

func min(a, bitArrays, c int) int {
	if a < bitArrays {
		if a < c {
			return a
		}
	} else {
		if bitArrays < c {
			return bitArrays
		}
	}
	return c
}

func MyerDist(word, text string) int {
	//make sure word is <= text
	if len(word) > len(text){
		text, word = word, text
	}
	var length,vn,vp,X,D0,hn,hp uint64

	//preprocessing, word -> bitarray for each letter position
	length = uint64(len(word))
	score := int(length) 
	// bitArrays := map[rune]uint64{'A': 0,'N': 0,'U': 0,'L': 0,'E': 0,'I': 0,'G': 0}
	bitArrays := map[rune]uint64 {'A':0, 'C':0, 'G':0, 'T':0}
	for i, char := range word {
		bitArrays[char] += 1 << uint(i)
	}	

	//main
	vn = 0
	vp = (1 << length) - 1
	for _, char := range text {
		X = bitArrays[char] | vn
		D0 = ((vp + (X & vp)) ^ vp) | X
		hn = vp & D0
		hp = vn | ^(vp | D0)
		X = hp << 1 | 1
		vn = X & D0
		vp = (hn << 1) | ^(X | D0)
		if hp & (1 << (length - 1)) != 0 {
			score += 1
		} else if hn & (1 << (length - 1)) != 0 {
			score -= 1
		}
	}
	return score
}

func MyerDistDiag(word, text string, width int) chan int {
	scores := make(chan int)

	if width >= 63{
		panic("width too big")
	}

	//make sure word is <= text
	if len(word) > len(text){
		text, word = word, text
	}

	/*
	difference bitarrays:
		h, v = horizontal, vertical 
		p,n = positive, negative 
		D0 = diagonal
	other vars:
		w = width of band
		c = offset; hardcoded to be w/2
		m = length of word
		score = the edit distance
	*/
	var vn,vp,X,D0,hn,hp uint64
	var w,c,m,score uint64
	m = uint64(len(word))
	w = uint64(width)
	c = (w/2)

	/*
	preprocessing, word -> bitarray for each letter position
	ex:
		ANNUAL:
			A: 010001
			N: 000110
			U: 001000
			L: 100000
	*/
	bitArrays := map[rune]uint64 {'A':1<<w, 'C':1<<w, 'G':1<<w, 'T':1<<w}
	for i, char := range word {
		if uint64(i) > c {
			break
		}
		bitArrays[char] += 1 << uint(i)
	}	
	for k, _ := range bitArrays {
		bitArrays[k] -= 1<<w
		bitArrays[k] <<= (w - c)
	}

	// Initialization vn = all 0s, vp = half 1s half 0s
	vn = 0
	vp = (1 << (w)) - 1
	vp >>= (w - c -1)
	vp <<= (w - c -1)

	//main loop
	go func() {
		defer close(scores)
		for pos, char := range text {
			p := uint64(pos+1)

			//bitmask shift the word bitarrays
			for k, _ := range bitArrays {
				bitArrays[k] >>= 1
			}
			if p + c <= m {
				bitArrays[rune(word[pos + int(c)])] |= (1 << (w - 1))
			}

			// Main computation
			B := bitArrays[char]
			X = B | vn
			D0 = ((vp + (X & vp)) ^ vp) | X
			hn = vp & D0
			hp = vn | ^(vp | D0)
			X = D0 >> 1 
			vn = X & hp
			vp = hn | ^(X | hp)
			
			//get score of middle diag
			score += 1 - ((D0 >> c) & 1)
			scores <- int(score)

		}
	}()
	return scores
}