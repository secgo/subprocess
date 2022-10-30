package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("please select one file subdomain")
		os.Exit(0)
	}
	f, err := ReadFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
	}
	lensub := Lensubdomain(f)

	for {
		var choose string
		for i := range lensub {
			fmt.Println(i)
		}
		fmt.Printf("\033[32mplease select number subDomain or\033[0m [e/exit]:")
		fmt.Scanln(&choose)
		if choose == "e" || choose == "exit" {
			os.Exit(0)
		}
		cint, _ := strconv.Atoi(choose)
		if _, ok := lensub[cint]; ok {
			for _, v := range lensub[cint] {
				fmt.Println(v)

			}
		} else {
			fmt.Printf("No subdomain length:%d\n", cint)
		}

	}
}

func ReadFile(file string) (map[string]int, error) {
	mk := make(map[string]int)
	f, err := os.Open(file)
	if err != nil {
		// fmt.Println(err)
		return nil, err
	}
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		if _, ok := mk[sc.Text()]; ok {
			mk[sc.Text()]++
		} else {
			mk[sc.Text()] = 1
		}
	}
	return mk, nil
}

func Lensubdomain(sub map[string]int) map[int][]string {
	length := make(map[int][]string)
	for k := range sub {
		s := strings.Count(k, ".") - 1 // how many host -> return int
		if _, ok := length[s]; ok {
			if !ContainSlice(length[s], k) {
				length[s] = append(length[s], k)
			}
		} else {
			length[s] = append(length[s], k)

		}

	}
	return length

}

func ContainSlice(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false

}
