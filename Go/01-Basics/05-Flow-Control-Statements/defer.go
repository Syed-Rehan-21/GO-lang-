package main
import ("fmt")
func main(){
	CountDown(10);
}
func CountDown(x int){
	for i := 1;i <=x ; i++ {
		defer fmt.Println(i)// This defer keyword helps in executing the code
		// When the program is done
	}
}