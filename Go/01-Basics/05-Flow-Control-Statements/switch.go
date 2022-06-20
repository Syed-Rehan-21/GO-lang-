package main
import (
	"fmt"
"time"
)
func main(){
	var i int=2;
	switch i{
	case 1:
		fmt.Println("Day 1");
	case 2:
		fmt.Println("Day 2");
	case 3:
		fmt.Println("Day 3");
	case 4:
		fmt.Println("Day 4");
	case 5:
		fmt.Println("Day 5");
	case 6:
		fmt.Println("Day 6");
	case 7:
		fmt.Println("Day 7");
	default:
		fmt.Printf("Day %d",i);
		fmt.Println()
	}
	t := time.Now()
	// Second with no condition
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
	
}