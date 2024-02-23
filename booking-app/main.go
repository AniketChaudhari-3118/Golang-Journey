package main

//import "fmt"
import (
	"fmt"
	"sync"
	"time"
	//"strconv"
)

/*var conferenceName = "Go Conference"
const conferenceTickets = 50
var remainingTickets = 50
*/
//var conferenceName string = "Go Conference"
//conferenceName := "Go Conference"
var conferenceName = "Go Conference"

const conferenceTickets int = 50

var remainingTickets uint = 50

// var bookings = []string{}  (created a empty slice of list of strings)
// var bookings = make([]map[string]string, 0) //(created a slice of list of maps) we have to give initial size because it is not just a map it is list of maps
var bookings = make([]UserData, 0) // As in maps we cannot use heterogenous data so that is why we use structure
// var bookings [50]string (Array)
//bookings := []string{}

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{} // synchronization

func main() {

	//greetUsers(conferenceName, conferenceTickets, remainingTickets) it can also be declared this way as the variables are declared globally
	//we dont need to pass the parameters to the function
	greetUsers()
	/*var userName string
	var userTickets int
	fmt.Println(remainingTickets)
	fmt.Println(&remainingTickets)
	*/
	//remainingTickets > 0 && len(bookings) < 50 (it can be the condition for the for loop)
	//for {
	firstName, lastName, email, userTickets := getUserInput()
	isValidName, isValidEmail, isValidTicketNumber := ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

	if isValidName && isValidEmail && isValidTicketNumber {

		bookticket(userTickets, firstName, lastName, email)

		wg.Add(1)                                              // the number of threads the application should wait for after the execution of main thread or other thread
		go sendTicket(userTickets, firstName, lastName, email) // go keyword is used because sendTicket method will take some time to run
		//and the next code needs to be executed continueously therefore to create a seperate thread for it gois used which is also
		//known as multithreading in java

		firstNames := getFirstNames()
		fmt.Printf("These are all our bookings: %v\n", firstNames)

		if remainingTickets == 0 {
			//end program
			fmt.Println("Our conference is booked out, come back next year.")
			//break (because no for loop)
		}
	} else {
		//fmt.Printf("We only have %v tickets remaining, so you cant book %v tickets\n", remainingTickets, userTickets)
		//fmt.Println("Your input data is invalid, try again")
		if !isValidName {
			fmt.Println("First name or Last name you entered is too short")
		}
		if !isValidEmail {
			fmt.Println("Email address you entered doesn't contain @ sign")
		}
		if !isValidTicketNumber {
			fmt.Println("Number of tickets you entered is invalid")
		}
	}
	wg.Wait() // this function tells the main thread to wait before termination for the other thread to get executed

	//}
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

}

func getFirstNames() []string {
	firstNames := []string{}
	//for index, booking := range bookings {
	for _, booking := range bookings { //_ is used when we a variable is not used further(in this case index is the variable)
		//var names = strings.Fields(booking) //Split the string (this slice is just a collections of strings)
		//var firstName = names[0]
		//firstNames = append(firstNames, booking["firstName"]) // as each element in booking slice has became a map so it is more easy to extract the first name from the map
		firstNames = append(firstNames, booking.firstName) // to access the firstName from the userdefined structure
	}
	//fmt.Printf("These are all our bookings: %v\n", bookings)

	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)

	fmt.Printf("Enter your last name:")
	fmt.Scan(&lastName)

	fmt.Printf("Enter your email:")
	fmt.Scan(&email)

	fmt.Printf("Enter number of tickets:")
	fmt.Scan(&userTickets)
	/*userName = "Tom"
	userTickets = 2
	fmt.Printf("User %v booked %v tickets.\n", userName, userTickets)
	*/

	return firstName, lastName, email, userTickets
}

func bookticket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - uint(userTickets)
	//bookings[0] = firstName + " " + lastName

	//create a map for user
	//var userData = make(map[string]string)
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}
	/*
		userData["firstName"] = firstName
		userData["lastName"] = lastName
		userData["email"] = email
		userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)
	*/
	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(20 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for  %v %v", userTickets, firstName, lastName)
	fmt.Println("##################")
	fmt.Printf("Sending ticket:\n %v \nto email addresss %v\n", ticket, email)
	fmt.Println("##################")
	wg.Done() //after the execution of the waiting thread (this method decreases the counter of waiting thread which is in Add function)
}

/*city := "London"

switch city {
    case "New York":
	    //execute code for booking new york conference tickets

    case "Singapore", "Hong Kong":
	    //execute code for booking singapore and hong kong conference tickets

    case "Mexico  City":
	    //execute code for booking Mexico city conference tickets
    default:
	    fmt.Println("no valid city selected")
}
*/

/*
	fmt.Printf("The whole Slice: %v\n", bookings)
		fmt.Printf("The first value: %v\n", bookings[0])
		fmt.Printf("Slice type: %T\n", bookings)
		fmt.Printf("Slice length: %v\n", len(bookings))
*/

//note: when we remove the for loop the main thread doesnt waits for other thread that is why the "go sendTicket" didnt get executed
