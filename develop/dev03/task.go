package main
 

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func getFile(path string, uniqueRequired bool) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	set := make(map[string]struct{})
	scanner := bufio.NewScanner(file)
	data := make([]string, 0)

	for scanner.Scan() {
		text := scanner.Text()
		if uniqueRequired {
			if _, ok := set[text]; ok {
				continue
			}
			set[text] = struct{}{}
		}
		data = append(data, text)
	}

	return data, nil
}

func sortByNumeric(data []string) {
	sort.Slice(data, func(i, j int) bool {
		vi, _ := strconv.Atoi(data[i])
		vj, _ := strconv.Atoi(data[j])
		return vi < vj
	})
}

func main() {
	var (
		filename     string
		sortNumeric  bool
		sortReverse  bool
		unique       bool
	)

	args := os.Args[1:]
	for i := 0; i < len(args); i++ {
		arg := args[i]
		switch arg {
		case "-f":
			if i+1 < len(args) {
				filename = args[i+1]
				i++
			} else {
				fmt.Println("No file in  -f")
				return
			}
		case "-n":
			sortNumeric = true
		case "-r":
			sortReverse = true
		case "-u":
			unique = true
		default:
			fmt.Printf("Incorect flag %s\n", arg)
			return
		}
	}

	if filename == "" {
		fmt.Println("No file in (-f)")
		return
	}

	file, err := getFile(filename, unique)
	if err != nil {
		panic(err)
	}

	if sortNumeric {
		sortByNumeric(file)
	} else {
		sort.Strings(file)
	}

	if sortReverse {
		for i := len(file) - 1; i >= 0; i-- {
			fmt.Println(file[i])
		}
	} else {
		for _, value := range file {
			fmt.Println(value)
		}
	}
} //TODO push another vers.