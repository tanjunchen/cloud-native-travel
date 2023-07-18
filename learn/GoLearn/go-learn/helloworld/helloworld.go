package helloworld

import "fmt"

func HelloWorld() (result int) {
	ss, err := fmt.Println("a")
	if err != nil {
		fmt.Println(ss)
	}
	fmt.Println(ss)
	return ss
}

func Sum(set []int) int {
	var result int
	for _, num := range set {
		result += num
	}
	return result
}

const englishHelloPrefix = "Hello,"

func Hello(name string) string {
	return englishHelloPrefix + name
}

func HelloRefactor(name string) string {
	return englishHelloPrefix + name
}

func HelloParameter(name string, language string) string {
	if name == "" {
		name = "World"
	}
	if language == "Spanish" {
		return "Hola, " + name
	}

	return englishHelloPrefix + name
}
