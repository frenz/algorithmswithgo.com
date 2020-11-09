package module01

// Reverse will return the provided word in reverse
// order. Eg:
//
//   Reverse("cat") => "tac"
//   Reverse("alphabet") => "tebahpla"
//

func Reverse(word string) string {
	if len(word) <= 1 {
		return word
	}
	return Reverse(string(word[1:])) + string(word[0])
}
