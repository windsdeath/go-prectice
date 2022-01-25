package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	stdin := bufio.NewReader(os.Stdin) // 표준 입력

	var a int
	var b int

	n, err := fmt.Scanln(&a, &b)
	if err != nil {
		fmt.Println(err)
		stdin.ReadString('\n') //'\n'까지 읽어와라 (버퍼를 비워라)
	} else {
		fmt.Println(n, a, b)
	}

	n, err = fmt.Scanln(&a, &b)
	if err != nil {
		fmt.Println(err)
		stdin.ReadString('\n') //'\n'까지 읽어와라 (버퍼를 비워라)
	} else {
		fmt.Println(n, a, b)
	}

}
