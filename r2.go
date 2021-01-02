package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var (
	gernerate    []string
	userProcess  string
	userterminal string
	userInput    string
	//force out
	errout bool

	//PredictS has 1 rule.
	PredictS []rule
	//PredictC has 2 rules.
	PredictC []rule
	//PredictA has 2 rules.
	PredictA []rule
	//PredictB has 2 rules.
	PredictB []rule
	//PredictQ has 2 rules.
	PredictQ []rule

	//PredictProg has 1 rule.
	//PredictProg []rule
	//PredictDcls has 2 rules.
	//PredictDcls []rule
	//PredictDcl has 2 rules.
	//PredictDcl []rule
	//PredictStmts has 2 rules.
	//PredictStmts []rule
	//PredictStmt is in case 2 and has 2 rules.
	//PredictStmt []rule
	//PredictExpr is in case 2 and has 3 rules.
	//PredictExpr []rule
	//PredictVal is in case 2 and has 3 rules.
	//PredictVal []rule
)

type rule struct {
	num     uint
	predict []string
	context string
}

func init() {
	PredictS = []rule{
		rule{
			num:     1,
			predict: []string{"a", "b", "q", "c", "$"},
			context: "A C $",
		},
	}
	PredictC = []rule{
		rule{
			num:     2,
			predict: []string{"c"},
			context: "c",
		},
		rule{
			num:     3,
			predict: []string{"d", "$"},
			context: "L",
		},
	}
	PredictA = []rule{
		rule{
			num:     4,
			predict: []string{"a"},
			context: "a B C d",
		},
		rule{
			num:     5,
			predict: []string{"b", "q", "c", "$"},
			context: "B Q",
		},
	}
	PredictB = []rule{
		rule{
			num:     6,
			predict: []string{"b"},
			context: "b B",
		},
		rule{
			num:     7,
			predict: []string{"q", "c", "d", "$"},
			context: "L",
		},
	}
	PredictQ = []rule{
		rule{
			num:     8,
			predict: []string{"q"},
			context: "q",
		},
		rule{
			num:     9,
			predict: []string{"c", "$"},
			context: "L",
		},
	}
}

func main() {
	var (
		way   string
		input string
	)
	for way != "exit" {
		fmt.Println("input,file,exit")
		fmt.Scanln(&way)
		switch {
		case way == "input":
			fmt.Print("input:")
			in := bufio.NewReader(os.Stdin)
			input, _ = in.ReadString('\n')
			restart(input)
		case way == "file":
			fmt.Print("filePath:")
			var path string
			fmt.Scanln(&path)
			input, err := getFileContext(path)
			if err != nil {
				fmt.Println(err)
			}
			restart(input)
		case way == "exit":
			fmt.Println("Exit!")
		default:
			continue
		}
	}

}

func getFileContext(filePath string) (string, error) {
	filePath = strings.Trim(filePath, " \n\t")
	context, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(context), nil
}

func reset() {
	userInput = ""
	userProcess = ""
	errout = false
}

func restart(newUserInput string) {
	reset()
	newUserInput = strings.Trim(newUserInput, " \n\t\r")
	userInput = newUserInput
	userProcess = newUserInput
	gernerate = strings.Split(userInput, " ")
	userterminal = gernerate[0]

	S()
}

func S() {
	switch {
	case contains(PredictS[0].predict, userterminal):
		fmt.Printf("%d ", PredictS[0].num)
		A()
		C()
		match("$")
		if !errout {
			fmt.Println("Accept")
		}
	default:
		fmt.Printf("Error(S vs. %s)\n", userterminal)
		errout = true
	}

}

func C() {
	if errout {
		return
	}
	switch {
	case contains(PredictC[0].predict, userterminal):
		fmt.Printf("%d ", PredictC[0].num)
		match("c")
	case contains(PredictC[1].predict, userterminal):
		fmt.Printf("%d ", PredictC[1].num)
		return
	default:
		fmt.Printf("Error(C vs. %s)\n", userterminal)
		errout = true
	}
}

func A() {
	if errout {
		return
	}
	switch {
	case contains(PredictA[0].predict, userterminal):
		fmt.Printf("%d ", PredictA[0].num)
		match("a")
		B()
		C()
		match("d")
	case contains(PredictA[1].predict, userterminal):
		fmt.Printf("%d ", PredictA[1].num)
		B()
		Q()
	default:
		fmt.Printf("Error(A vs. %s)\n", userterminal)
		errout = true
	}
}

func B() {
	if errout {
		return
	}
	switch {
	case contains(PredictB[0].predict, userterminal):
		fmt.Printf("%d ", PredictB[0].num)
		match("b")
		B()
	case contains(PredictB[1].predict, userterminal):
		fmt.Printf("%d ", PredictB[1].num)
		return
	default:
		fmt.Printf("Error(B vs. %s)\n", userterminal)
		errout = true
	}
}

func Q() {
	if errout {
		fmt.Println("out")
		return
	}
	switch {
	case contains(PredictQ[0].predict, userterminal):
		fmt.Printf("%d ", PredictQ[0].num)
		match("q")
	case contains(PredictQ[1].predict, userterminal):
		fmt.Printf("%d ", PredictQ[1].num)
		return
	default:
		fmt.Printf("Error(Q vs. %s)\n", userterminal)
		errout = true
	}
}

func match(terminal string) {
	if errout {
		return
	}
	if terminal == userterminal {
		if len(gernerate) == 1 {
			return
		}
		gernerate = gernerate[1:]
		userterminal = gernerate[0]
		return
	}
	fmt.Printf("Error(Expected %s)\n", terminal)
	errout = true
	return
}

// IsPredict return search result
// Is peek in predict?
func contains(set []string, peek string) bool {
	for _, terminal := range set {
		if peek == terminal {
			return true
		}
	}
	return false
}
