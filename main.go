package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// func showText(message string) {
// 	fmt.Println(message)
// }

// func multiTextHandler(names []string, functionHandler func(string)) {
// 	for _, v := range names {
// 		functionHandler(v)
// 	}
// }

// func updateName(name *string) {
// 	*name = "new value"
// }

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)

	// fmt.Print("Create new bill : ")
	// name, _ := reader.ReadString('\n')
	// fmt.Println(name)
	// name = strings.TrimSpace(name)

	name, _ := getInput("create new bill : ", reader)

	b := newBill(name)
	fmt.Println("your bill is :", b.name)

	return b

}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)
	opt, _ := getInput("choose option (a - add item,s - save bill,t - add tip) : ", reader)
	switch opt {
	case "a":
		name, _ := getInput("enter name : ", reader)
		price, _ := getInput("enter price : ", reader)
		p, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("the price must be number")
			promptOptions(b)
		}
		b.addItemToBill(name, p)
		fmt.Println("item added")
		promptOptions(b)
	case "t":
		tip, _ := getInput("enter tip : ", reader)
		t, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("the tip must be number")
			promptOptions(b)
		}
		b.updateTip(t)
		fmt.Println("tip updated successfuly")
		promptOptions(b)
	case "s":
		b.saveBill()
		fmt.Println("you saved file",b.name)
	default:
		fmt.Print("please select valid option \n")
		promptOptions(b)
	}
}

func main() {

	myBill := createBill()
	promptOptions(myBill)
	// myNewBill := newBill("mahdi")

	// myNewBill.addItemToBill("test",2893)
	// myNewBill.updateTip(1000)
	// fmt.Println(myNewBill.format())

	// name := "test"
	// fmt.Println(name)
	// updateName(&name)
	// fmt.Println(name)
	// multiTextHandler([]string{"test", "this is sample text", "teststtts"}, showText)
	// var test1 string = "test1"
	// var test2 = "test2"
	// var test3 string
	// var ages = [3]int{1, 2, 3}
	// fmt.Println(ages,len(ages))

	// names := []string{"mahdi", "amir", "masooud"}
	// names[2] = "amir mohammad"
	// names = append(names, "test")
	// fmt.Println(names)

	// content := "hello this is test"

	// fmt.Println(strings.Contains(content,"this"))

	// // mytesttext := strings.Split(content," ")

	// myints := []int{46,56464,3245345,345,435,45345,2,1,453,545,5,66,45,35345,34534535345,35345343,333,33,224}

	// sort.Ints(myints)

	// fmt.Println(myints)
	// x := 0
	// for x < 5 {

	// 	fmt.Println("x value is : ", x)
	// 	x++
	// }

	// for i := 0; i < 1000000; i++ {
	// 	fmt.Println("i value is : ", i)
	// }

	// names := []string{"mahdi", "amir", "masooud"}

	// for index, value := range names{
	// 	fmt.Printf("this is the index and value : %v , %v \n",index,value)
	// }

	// for index, value := range names {
	// 	if index == 1 {
	// 		fmt.Printf("index continue , %v \n", index)
	// 		continue
	// 	}
	// 	fmt.Printf("index is : %v and value is %v \n", index, value)
	// }

}
