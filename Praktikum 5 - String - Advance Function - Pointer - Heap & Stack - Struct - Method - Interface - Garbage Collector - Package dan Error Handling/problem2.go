package main

func main() {

	/*-----Problem 2-----*/
	println(problem2(1, "abc"))
	println(problem2(1, "alta"))
	println(problem2(1, "alterraacademy"))
	println(problem2(1, "abcdefghijklmnopqrstuvwxyz"))
	println(problem2(1, "abcdefghijklmnopqrstuvwxyz"))
}

func problem2(offset int, input string) string {

	var result = ""
	for _, v := range input {

		if offset+int(v) > 122 {
			var differ = (int(v) + offset) % 122
			result += string(96 + differ)
		} else {
			result += string(int(v) + offset)
		}

	}

	return result
}
