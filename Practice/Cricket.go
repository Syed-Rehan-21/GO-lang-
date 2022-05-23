package main

import "fmt"

//use Data base
func userInfo() string {
	var name string
	fmt.Print("Enter Your Name : ")
	fmt.Scan(&name)
	var Email string
	fmt.Print("Enter Yout Email :")
	fmt.Scan(&Email)
	var uid string
	uid = name + " " + Email
	return uid
}

//Tickets sold-out
func exception(i uint) {
	if i == 0 {
		fmt.Print("\t\tSorry 😓!!\nThis Match Tickets are Totally Booked\n")
		fmt.Print("\t\tSorry for Disapointing\n")
	}
	return
}

func intro() {
	var t = [5]uint{50, 50, 50, 50, 50}
	//
	var M = [5]string{
		"1. MI   vs RCB",
		"2. CSK  vs SRH ",
		"3. RR   vs GT ",
		"4. PBKS vs LSG ",
		"5. KKR  vs DC "}
	//
	var i int
	var bt uint
	// fmt.Printf("\nUpcoming  Matches  :\n\t%v\n\t%v\n\t%v\n\t%v\n\t%v \nSelect  match Number  : ", M1, M2, M3, M4, M5)
	fmt.Printf("\nUpcoming  Matches  :\n\t%v\n\t%v\n\t%v\n\t%v\n\t%v\nSelect  match Number  : ", M[0], M[1], M[2], M[3], M[4])
	fmt.Scan(&i)

	fmt.Printf("You Have Selected '%v'\n", M[i])
	exception(t[i])
	fmt.Printf("Only %v Tickets are Remaining Hurry Up \n", t[i+1])
	fmt.Printf("\nHow many Match tickets you want to Book  : ")
	fmt.Scan(&bt)
	t[i] = t[i] - bt

}
func main() {
	// Cricket match ticket Booking CLI application
	fmt.Printf("\nEnjoy The Energy of IPL 😎 Directly from the Stadium\n")
	fmt.Printf("\n\tHurry up !! Book your Tickets\n")
	var booking []string
	var j, tid int
	for {
		fmt.Print("Signin (1)/signup (2) : ")
		fmt.Scan(&j)
		if j == 2 {
			booking = append(booking, userInfo())
			fmt.Printf("Your Ticket Number %v \n", len(booking)-1)
		}
		if j == 1 {
			fmt.Print("Enter Your Ticket Number : ")
			fmt.Scan(&tid)
			// fmt.Print("User Info : \n", booking[tid])
			intro()
		}
		// fmt.Println(booking)
	}
}
