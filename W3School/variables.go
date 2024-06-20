package main;

import("fmt")

func main() {
	//Single Variable Decleration 
	 
	var item1 int  =10; // standard declearation of a variable 
	item2 := "Anjanay Raina"; // inferred declaration of the type of the variable 
	fmt.Printf("Item1 type : %T and Item2 type : %T \n" , item1 , item2);

	// Multiple Variable Decleration 

	var a , b , c = 10 , "Hello" , 12.34; 
	fmt.Println(a , b , c);
}