package main

import (
	"fmt"
	"time"
	"./my_strings"
)

func main() {
	fmt.Printf("This is the main function begin:\n");

	fmt.Printf("================================\n");
	fmt.Printf("Run the RotatingTest \n");
	my_strings.Rotating();
	fmt.Printf("================================\n");

	fmt.Printf("Run the IsPalindrome \n");
	fmt.Printf("%t \n", my_strings.IsPalindrome("ABABABA"));
	fmt.Printf("================================\n");

	fmt.Printf("Run the Manacher \n");
	start_time := time.Now();
	fmt.Printf("%d \n",my_strings.Manacher("ABABABA"));
	end_time := time.Now();
	fmt.Println(end_time.Sub(start_time));
	fmt.Printf("================================\n");
}
