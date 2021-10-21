package main

func main() {
BREAK:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 3 {
				break BREAK
			}
		}

	}
}
