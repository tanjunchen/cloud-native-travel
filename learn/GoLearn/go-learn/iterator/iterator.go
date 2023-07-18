package iterator

func Repeat(character string) string {
	var repeated string
	for i := 0; i < 5; i++ {
		repeated = repeated + character
	}
	return repeated
}

const repeatCount = 5

func Repeat2(character string) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += character
	}
	return repeated
}

func Repeat3(character string,repeatCount2 int) string {
	var repeated string
	for i := 0; i < repeatCount2; i++ {
		repeated += character
	}
	return repeated
}

