package main

import "fmt"

// APPENDING A SLICE

/*
func main() {
slice1 := []int { 1, 2, 3, 4}
slice2 := append(slice1 ,5 ,6 ,7 ,8 ,9)
fmt.Println(slice1 ,slice2)
slice3 := []string{"John"}
slice4 := append(slice3 , "Stone")
fmt.Println(slice3,slice4)
}
*/

//COPYING A SLICE
/*
func main() {
	slice1 :=[]int{1, 2, 3}
	slice2 :=make([]int ,3)
	copy(slice2,slice1)
	fmt.Println(slice1,slice2)
}
*/

//Maps
/*
func main() {
elements := make (map[string]string)
elements["H"]="Hydrogen"
elements["He"]="helium"
elements["Li"]="Litium"
elements["Be"]="Berilium"
elements["B"]="Boron"
elements["C"]="Carbon"
elements["N"]="Nitrogen"
elements["O"]="Oxygen"
elements["F"]="Fluorine"
elements["Ne"]="Neon"
fmt.Println(elements["Li"])
fmt.Println(elements["K"]=="")
}
*/

func main() {

	var n, smallest, largest int
	x := []int{
		48, 96, 86, 68, 57, 64,
		37, 82, 63, 70, 74, 47,
		74, 67, 678, 98, 989, 64,
	}
	for _, v := range x {
		if v > n {
			fmt.Println(v, ">", n)
			n = v
			largest = n
		}else{
			fmt.Println(v, "<" , n)
		}

	}
	fmt.Println("The Largest number is ", largest)
	for _, v := range x {
		if v > n {
			fmt.Println(v, ">", n)
		}else{
			fmt.Println(v, "<" , n)
			n = v
			smallest = n
		}
	}
	fmt.Println("The smallest number is ", smallest)
}
