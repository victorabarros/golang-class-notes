package main

import "fmt"

func main() {
	simpleLoop()
	loopInDetails()
	loopInDetailsV2()

	simpleIf()
	ifDetailed()
}

func simpleIf() {
	err := fmt.Errorf("new error")

	if err != nil {
		fmt.Println(err)
	}
}

func ifDetailed() {
	var err error = nil

	preStatement := func() {
		err = fmt.Errorf("new error")
	}

	conditional := func() bool {
		return err != nil
	}

	action := func() {
		fmt.Println(err)
	}

	if preStatement(); conditional() {
		action()
	}
}

func simpleLoop() {
	for ii := 0; ii < 10; ii++ {
		fmt.Println(ii)
	}
}

func loopInDetails() {
	var top int = 5
	var ii int

	preStatement := func() {
		fmt.Print("preStatement ii=0\t")
		ii = 0
	}

	conditional := func() bool {
		ans := ii < top
		fmt.Print("conditional ii < top: ", ans, "\t")
		return ans
	}

	posLoop := func() {
		fmt.Print("posLoop ii++\t")
		ii++
	}

	action := func() {
		fmt.Print("action ii: ", ii, "\n")
	}

	fmt.Print("StartLoop\n")
	for preStatement(); conditional(); posLoop() {
		action()
	}
	fmt.Print("\nFinishLoop\n")
}

func loopInDetailsV2() {
	var top int = 5
	var ii int = 0

	conditional := func() bool {
		ans := ii < top
		fmt.Print("conditional ii < top: ", ans, "\t")
		return ans
	}

	action := func() {
		fmt.Print("action ii: ", ii, "\n")
	}

	fmt.Print("StartLoopV2\n")
	for conditional() {
		action()
		ii++
	}
	fmt.Print("\nFinishLoop\n")
}
