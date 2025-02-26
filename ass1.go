package main

import "fmt"

// now creating a structure named person
type person struct {
	Index  int
	Name   string
	Age    int
	Gender string
	Email  string
	Ph_No  int64
}

var persons []person // A slice to store multiple persons

func addingNewPerson(name string, Age int, Gender string, Email string, Ph_No int64) {
	//Conditions to handle edge cases where name or age or Phone no is invalid
	if name == "" {
		fmt.Println("Name Can Not be Empty\n")
	}
	if Age <= 0 {
		fmt.Println("Age Cannot be less than Zero")
	}
	if Ph_No <= 0 {
		fmt.Println("Enter a Valid Phone No")
	}
	ind := len(persons) + 1 //to start with index 1
	newPerson := person{
		Index:  ind,
		Name:   name,
		Age:    Age,
		Gender: Gender,
		Email:  Email,
		Ph_No:  Ph_No,
	}
	persons = append(persons, newPerson)
	fmt.Println("Person added")
	//return nil
}

func introduce() {
	if len(persons) == 0 { // in case there is no person data
		fmt.Println("No Data To Show")
		return
	}
	for i, person := range persons {
		fmt.Printf(
			"Index: %d || Name: %s || Age: %d || Gender: %s || Email: %s || Phone Number: %d\n",
			i+1,           // Index
			person.Name,   // fied stores the name data
			person.Age,    // storing the age
			person.Gender, // storing gender
			person.Email,  // stores email addresses
			person.Ph_No,  // stores phone numbers
		)
	}
}

func updatename() {
	var nm string  //to take the name from the user , whose name to be changed
	var nn string  //to store new name and update it in the slice
	found := false //flag is person not found
	fmt.Println("\nEnter the Name you want to update")
	fmt.Scanln(&nm)
	fmt.Println("\nEnter the new Name")
	fmt.Scanln(&nn)
	for i, person := range persons {
		if person.Name == nm {
			fmt.Println("Person found and name will be updated")
			persons[i].Name = nn
			found = true
		}
	}
	if !found {
		fmt.Println("\n404 Not Found")
	}
}

func updateage() {
	var nn string  //to take the name from the user , whose age you want to change
	var newAge int //to store the new age to update it in the slice
	found := false //flag is person not found
	fmt.Println("\nEnter the person's name for who you want to update the Age")
	fmt.Scanln(&nn)
	fmt.Println("\nEnter the person's name new Name")
	fmt.Scanln(&newAge)
	for i, person := range persons {
		if person.Name == nn {
			persons[i].Age = newAge
			found = true
		}
	}
	if !found {
		fmt.Println("\n404 Not Found")
	}
}
func checkVote() {
	var nn string  //to input the name for which you want to check if age>18
	found := false //flag is person not found
	fmt.Println("\nEnter the person's name for who you want to check if he is eligible")
	fmt.Scanln(&nn)
	for _, person := range persons {
		if person.Name == nn {
			found = true
			if person.Age >= 18 {
				fmt.Println("\nEligible for Voting")
			} else {
				fmt.Println("\nNot Eligible")
			}
			break // Exit the loop after checking
		}
	}
	if !found {
		fmt.Println("\n404 Name Not Found")
	}
}

func main() {
	var (
		Name   string
		Age    int
		Gender string
		Email  string
		Ph_No  int64
		key    int
	)

	for {
		fmt.Println("*****************************************************************Welcome*****************************************************************")
		fmt.Println("=======> 1. Add a Person")
		fmt.Println("==========> 2. Introduce/Display")
		fmt.Println("==============> 3. Update Name")
		fmt.Println("=================> 4. Update Age")
		fmt.Println("====================> 5. Check Voting Eligibility")
		fmt.Println("=======================> 6. Exit")
		fmt.Println("************************************************************Enter your choice index number*****************************************************************")
		fmt.Scanln(&key) //for user input

		switch key {
		case 1:
			fmt.Println("\nEnter the Name")
			fmt.Scanln(&Name)
			fmt.Println("\nEnter the Age")
			fmt.Scanln(&Age)
			fmt.Println("\nEnter the Gender as MALE OR FEMALE")
			fmt.Scanln(&Gender)
			fmt.Println("\nEnter the Email")
			fmt.Scanln(&Email)
			fmt.Println("\nEnter the Phone Number")
			fmt.Scanln(&Ph_No)
			addingNewPerson(Name, Age, Gender, Email, Ph_No)
		case 2:
			introduce()
		case 3:
			updatename()
		case 4:
			updateage()
		case 5:
			checkVote()
		case 6:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice. Please enter a valid option (1-6).")
		}

		fmt.Println("\nPress Enter to continue...")
		fmt.Scanln()
	}
}
