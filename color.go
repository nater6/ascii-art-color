package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	//Create the map and put in all values
	cMap := make(map[string]string)
	cMap["reset"] = "\033[0m"
	cMap["red"] = "\033[31m"
	cMap["green"] = "\033[32m"
	cMap["yellow"] = "\033[33m"
	cMap["blue"] = "\033[34m"
	cMap["purple"] = "\033[35m"
	cMap["cyan"] = "\033[36m"
	cMap["gray"] = "\033[37m"
	cMap["white"] = "\033[97m"
	cMap["orange"] = "\033[38;5;208m"

	//Take the collor from the terminal
	if len(os.Args)< 3 {
		fmt.Println("Usage: go run . [STRING] [OPTION]")
		fmt.Println("EX: go run . something --color=<color>")
		return
	}
	color := os.Args[2]
	if len(color) < 8 {
		fmt.Println("Usage: go run . [STRING] [OPTION]")
		fmt.Print("EX: go run . something --color=<color>")
		return
	}
	color = color[8:]

	//Create error msg for incorrect format
	if color != "red" && color != "green" && color != "yellow" && color != "blue" && color != "purple" && color != "cyan" && color != "gray" && color != "white" && color != "orange" {
		fmt.Println("Usage: go run . [STRING] [OPTION]")
		fmt.Print("EX: go run . something --color=<color>")
		return
	}

	file, err := os.Open("letters.txt")
	if err != nil {
		log.Fatalf("ERROR: %s", err)
	}

	// file to slice of string by line

	var lttrlines []string
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		lttrlines = append(lttrlines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	// find a string in a scanned file

	// Splits on newlines by default.
	ln1 := []string{}
	ln2 := []string{}
	ln3 := []string{}
	ln4 := []string{}
	ln5 := []string{}
	ln6 := []string{}
	ln7 := []string{}
	ln8 := []string{}

	line := 0
	Output := os.Args[1]
	SlcOutput := []rune(Output)
	dash := strings.Index(Output, `\n`)

	//If every letter in being colored
	if len(os.Args) == 3 {

		for i := 0; i < len(SlcOutput); i++ {
			if i == dash && dash >= 0 {
				if len(ln1) != 0 && len(ln2) != 0 && len(ln3) != 0 && len(ln4) != 0 && len(ln5) != 0 && len(ln6) != 0 && len(ln7) != 0 && len(ln8) != 0 {

					Printer(ln1, cMap, color)
					Printer(ln2, cMap, color)
					Printer(ln3, cMap, color)
					Printer(ln4, cMap, color)
					Printer(ln5, cMap, color)
					Printer(ln6, cMap, color)
					Printer(ln7, cMap, color)
					Printer(ln8, cMap, color)

				} else {
					fmt.Println()
				}

				ln1 = []string{}
				ln2 = []string{}
				ln3 = []string{}
				ln4 = []string{}
				ln5 = []string{}
				ln6 = []string{}
				ln7 = []string{}
				ln8 = []string{}
				dash = dash + strings.Index(Output[dash+1:], `\n`) + 1
				i++
			} else {
				for j, s := range lttrlines {
					line = j
					if len(s) > 0 && s == string(SlcOutput[i]) {
						break
					}

				}

				for k := line + 1; k <= line+8; k++ {
					if k == line+1 {
						ln1 = append(ln1, lttrlines[k])
					}
					if k == line+2 {
						ln2 = append(ln2, lttrlines[k])
					}
					if k == line+3 {
						ln3 = append(ln3, lttrlines[k])
					}
					if k == line+4 {
						ln4 = append(ln4, lttrlines[k])
					}
					if k == line+5 {
						ln5 = append(ln5, lttrlines[k])
					}
					if k == line+6 {
						ln6 = append(ln6, lttrlines[k])
					}
					if k == line+7 {
						ln7 = append(ln7, lttrlines[k])
					}
					if k == line+8 {
						ln8 = append(ln8, lttrlines[k])
					}

				}
				if i == len(Output)-1 {
					Printer(ln1, cMap, color)
					Printer(ln2, cMap, color)
					Printer(ln3, cMap, color)
					Printer(ln4, cMap, color)
					Printer(ln5, cMap, color)
					Printer(ln6, cMap, color)
					Printer(ln7, cMap, color)
					Printer(ln8, cMap, color)
				}
			}
		}
	} else if len(os.Args) == 4 {
		num, errx := strconv.Atoi(os.Args[3])
		if errx != nil {
			log.Fatal(errx.Error())
		}
		num = num - 1

		for i := 0; i < len(SlcOutput); i++ {

			if i == dash && dash >= 0 {
				if len(ln1) != 0 && len(ln2) != 0 && len(ln3) != 0 && len(ln4) != 0 && len(ln5) != 0 && len(ln6) != 0 && len(ln7) != 0 && len(ln8) != 0 {

					Printer1(num, ln1, cMap, color)
					Printer1(num, ln2, cMap, color)
					Printer1(num, ln3, cMap, color)
					Printer1(num, ln4, cMap, color)
					Printer1(num, ln5, cMap, color)
					Printer1(num, ln6, cMap, color)
					Printer1(num, ln7, cMap, color)
					Printer1(num, ln8, cMap, color)

				} else {
					fmt.Println()
				}

				ln1 = []string{}
				ln2 = []string{}
				ln3 = []string{}
				ln4 = []string{}
				ln5 = []string{}
				ln6 = []string{}
				ln7 = []string{}
				ln8 = []string{}
				dash = dash + strings.Index(Output[dash+1:], `\n`) + 1
				i++
				num = num - i - 2
			} else {
				for j, s := range lttrlines {
					line = j
					if len(s) > 0 && s == string(SlcOutput[i]) {
						break
					}

				}

				for k := line + 1; k <= line+8; k++ {
					if k == line+1 {
						ln1 = append(ln1, lttrlines[k])
					}
					if k == line+2 {
						ln2 = append(ln2, lttrlines[k])
					}
					if k == line+3 {
						ln3 = append(ln3, lttrlines[k])
					}
					if k == line+4 {
						ln4 = append(ln4, lttrlines[k])
					}
					if k == line+5 {
						ln5 = append(ln5, lttrlines[k])
					}
					if k == line+6 {
						ln6 = append(ln6, lttrlines[k])
					}
					if k == line+7 {
						ln7 = append(ln7, lttrlines[k])
					}
					if k == line+8 {
						ln8 = append(ln8, lttrlines[k])
					}

				}
				if i == len(Output)-1 {

					Printer1(num, ln1, cMap, color)
					Printer1(num, ln2, cMap, color)
					Printer1(num, ln3, cMap, color)
					Printer1(num, ln4, cMap, color)
					Printer1(num, ln5, cMap, color)
					Printer1(num, ln6, cMap, color)
					Printer1(num, ln7, cMap, color)
					Printer1(num, ln8, cMap, color)

				}
			}

		}
	} else if len(os.Args) == 5 {
		num1, err1 := strconv.Atoi(os.Args[3])
		if err1 != nil {
			log.Fatal(err1.Error())
		}
		num2, err2 := strconv.Atoi(os.Args[4])
		if err2 != nil {
			log.Fatal(err2.Error())
		}
		num1 -= 1
		num2 -= 1
		for i := 0; i < len(SlcOutput); i++ {

			if i == dash && dash >= 0 {
				if len(ln1) != 0 && len(ln2) != 0 && len(ln3) != 0 && len(ln4) != 0 && len(ln5) != 0 && len(ln6) != 0 && len(ln7) != 0 && len(ln8) != 0 {

					Printer2(num1, num2, ln1, cMap, color)
					Printer2(num1, num2, ln2, cMap, color)
					Printer2(num1, num2, ln3, cMap, color)
					Printer2(num1, num2, ln4, cMap, color)
					Printer2(num1, num2, ln5, cMap, color)
					Printer2(num1, num2, ln6, cMap, color)
					Printer2(num1, num2, ln7, cMap, color)
					Printer2(num1, num2, ln8, cMap, color)

				} else {
					fmt.Println()
				}

				ln1 = []string{}
				ln2 = []string{}
				ln3 = []string{}
				ln4 = []string{}
				ln5 = []string{}
				ln6 = []string{}
				ln7 = []string{}
				ln8 = []string{}
				dash = dash + strings.Index(Output[dash+1:], `\n`) + 1
				i++
				num1 = num1 - i - 2
				num2 = num2 - i - 2
			} else {
				for j, s := range lttrlines {
					line = j
					if len(s) > 0 && s == string(SlcOutput[i]) {
						break
					}

				}

				for k := line + 1; k <= line+8; k++ {
					if k == line+1 {
						ln1 = append(ln1, lttrlines[k])
					}
					if k == line+2 {
						ln2 = append(ln2, lttrlines[k])
					}
					if k == line+3 {
						ln3 = append(ln3, lttrlines[k])
					}
					if k == line+4 {
						ln4 = append(ln4, lttrlines[k])
					}
					if k == line+5 {
						ln5 = append(ln5, lttrlines[k])
					}
					if k == line+6 {
						ln6 = append(ln6, lttrlines[k])
					}
					if k == line+7 {
						ln7 = append(ln7, lttrlines[k])
					}
					if k == line+8 {
						ln8 = append(ln8, lttrlines[k])
					}

				}
				if i == len(Output)-1 {

					//Print each loop
					// 	fmt.Println(ln1)
					Printer2(num1, num2, ln1, cMap, color)
					Printer2(num1, num2, ln2, cMap, color)
					Printer2(num1, num2, ln3, cMap, color)
					Printer2(num1, num2, ln4, cMap, color)
					Printer2(num1, num2, ln5, cMap, color)
					Printer2(num1, num2, ln6, cMap, color)
					Printer2(num1, num2, ln7, cMap, color)
					Printer2(num1, num2, ln8, cMap, color)

				}
			}

		}

	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		// Handle the error
	}
	file.Close()

}

func Printer2(n1 int, n2 int, ln []string, cMap map[string]string, color string) {

	for j := 0; j < len(ln); j++ {
		if (j >= n1 && j <= n2) && j == len(ln)-1 {
			fmt.Println(string(cMap[color] + ln[j] + string(cMap["reset"])))
			continue
		}

		if j >= n1 && j <= n2 {
			fmt.Print(string(cMap[color] + ln[j] + string(cMap["reset"])))
			continue
		}

		if j == len(ln)-1 {
			fmt.Println(ln[j])
			continue
		}
		fmt.Print(ln[j])
	}

}

func Printer1(n int, ln []string, cMap map[string]string, color string) {
	for j := 0; j < len(ln); j++ {
		if (j == n) && j == len(ln)-1 {
			fmt.Println(string(cMap[color] + ln[j] + string(cMap["reset"])))
			continue
		}

		if j == n {
			fmt.Print(string(cMap[color] + ln[j] + string(cMap["reset"])))
			continue
		}

		if j == len(ln)-1 {
			fmt.Println(ln[j])
			continue
		}
		fmt.Print(ln[j])
	}
}

func Printer(ln []string, cMap map[string]string, color string) {
	for j := 0; j < len(ln); j++ {
		if j == len(ln)-1 {
			fmt.Println(string(cMap[color]) + ln[j] + cMap["reset"])
			continue
		}
		fmt.Print(string(cMap[color]) + ln[j] + cMap["reset"])
	}
}
