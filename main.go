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
func exception() {
	fmt.Print("\t\tSorry 😓!!\nThis Match Tickets are Totally Booked\n")
	fmt.Print("\t\tSorry for Disapointing\n")
	return
}

func intro() {
	var M1 = "1. MI   vs RCB"
	var M2 = "2. CSK  vs SRH "
	var M3 = "3. RR   vs GT "
	var M4 = "4. PBKS vs LSG "
	var M5 = "5. KKR  vs DC"
	var t = [5]uint{50, 50, 50, 50, 50}

	var i int
	var bt uint
	fmt.Printf("\nUpcoming  Matches  :\n\t%v\n\t%v\n\t%v\n\t%v\n\t%v \nSelect  match Number  : ", M1, M2, M3, M4, M5)
	fmt.Scan(&i)
	switch i {
	case 1:
		fmt.Printf("You Have Selected '%v'\n", M1)
		if t[0] == 0 {
			exception()
		}
		fmt.Printf("Only %v Tickets are Remaining Hurry Up \n", t[0])
		fmt.Printf("\nHow many Match tickets you want to Book  : ")
		fmt.Scan(&bt)
		t[0] = t[0] - bt

	case 2:
		fmt.Printf("You Have Selected '%v'\n", M2)
		if t[1] == 0 {
			exception()
		}
		fmt.Printf("Only %v Tickets are Remaining Hurry Up \n", t[1])
		fmt.Printf("\nHow many Match tickets you want to Book  : ")
		fmt.Scan(&bt)
		t[1] = t[1] - bt

	case 3:
		fmt.Printf("You Have Selected '%v'\n", M3)
		if t[2] == 0 {
			exception()
		}
		fmt.Printf("Only %v Tickets are Remaining Hurry Up \n", t[2])
		fmt.Printf("\nHow many Match tickets you want to Book  : ")
		fmt.Scan(&bt)
		t[2] = t[2] - bt
	case 4:
		fmt.Printf("You Have Selected '%v'\n", M4)
		if t[3] == 0 {
			exception()
		}
		fmt.Printf("Only %v Tickets are Remaining Hurry Up \n", t[3])
		fmt.Printf("\nHow many Match tickets you want to Book  : ")
		fmt.Scan(&bt)
		t[3] = t[3] - bt

	case 5:
		fmt.Printf("You Have Selected '%v'\n", M5)
		if t[4] == 0 {
			exception()
		}
		fmt.Printf("Only %v  Tickets are Remaining Hurry Up \n", t[4])
		fmt.Printf("\nHow many Match tickets you want to Book  : ")
		fmt.Scan(&bt)
		t[4] = t[4] - bt
	}
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
			fmt.Printf("Your TicketID %v \n", len(booking)-1)
		}
		if j == 1 {
			fmt.Print("Enter Your TicketID : ")
			fmt.Scan(&tid)
			fmt.Print("User Info : \n", booking[tid])
			intro()
		}
		// fmt.Println(booking)
	}
}
