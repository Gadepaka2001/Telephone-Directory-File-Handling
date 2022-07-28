package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func ReadFile(F string) string {
	path := "./DATA/" + F
	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic("Error Ocuured while reading file")
	}
	return string(data)
}

func AppendIntoFile(F string, data string) {
	path := "./DATA/" + F
	S := ReadFile(F) + data + "\n"
	err := ioutil.WriteFile(path, []byte(S), 0644)
	if err != nil {
		panic("Error while opening a file")
	}
}

func Search(F string, name string) (string, bool) {
	S := ReadFile(F)
	Lst := strings.Split(S, "\n")
	for i := 0; i < len(Lst)-1; i++ {
		list := strings.Split(Lst[i], ":")
		if list[0] == name {
			return "Contact Name: " + list[0] + "\nMobile Number: " + list[1], true
		}
	}
	return " ", false
}

func WriteIntoFile(F string, data string, B bool) {
	path := "./DATA/" + F
	if B == true {
		err := ioutil.WriteFile(path, []byte(data+"\n"), 0644)
		if err != nil {
			panic("Error while opening a file")
		}
	} else {
		err := ioutil.WriteFile(path, []byte(data), 0644)
		if err != nil {
			panic("Error while opening a file")
		}
	}
}

func Delete(F string, C string) bool {
	S := ReadFile(F)
	Lst := strings.Split(S, "\n")
	for i := 0; i < len(Lst)-1; i++ {
		list := strings.Split(Lst[i], ":")
		if list[0] == C {
			Lst = append(Lst[:i], Lst[i+1:]...)
			S = strings.Join(Lst, "\n")
			WriteIntoFile(F, S, false)
			return true
		}
	}
	return false
}

func CreateFile(F string) {
	path := "./DATA/" + F
	_, err := os.Create(path)
	if err != nil {
		panic("Error while creating a file")
	}
}

func Check(F string, name string) bool {
	S := ReadFile(F)
	Lst := strings.Split(S, "\n")
	for i := 0; i < len(Lst)-1; i++ {
		lst := strings.Split(Lst[i], ":")
		if name == lst[0] {
			return true
		}
	}
	return false
}

func Command(user string) {
	for {
		var C string
		fmt.Printf("\n[%v] Enter Command ('a' for add) || ('m' for modify) || ('s' for search) || ('d' for delete) || ('all' for all contacts) || ('e' for login): ", user)
		fmt.Scanln(&C)
		if C == "a" || C == "A" {
			var name string
			var phone string
			fmt.Printf("Contact Name: ")
			fmt.Scanln(&name)
			fmt.Printf("Contact Number: ")
			fmt.Scanln(&phone)
			File := user + ".txt"
			S := name + ":" + phone
			b := Check(File, name)
			if b == true {
				fmt.Printf("Cannot add %v\n", name)
			} else {
				AppendIntoFile(File, S)
				println("---------------------------Added Successfully-----------------------------")
			}
		} else if C == "m" || C == "M" {
			File := user + ".txt"
			L := strings.Split(ReadFile(File), "\n")
			var temp bool = false
			for {
				var S string
				fmt.Printf("Enter contact name to modify: ")
				fmt.Scanln(&S)
				for i := 0; i < len(L)-1; i++ {
					Lst := strings.Split(L[i], ":")
					name := Lst[0]
					if name == S {
						temp = true
						for {
							var Mc string
							fmt.Printf("what You want to modify Contact name or phoneno (c/p): ")
							fmt.Scanln(&Mc)
							if Mc == "c" {
								var Name string
								fmt.Printf("Enter New UserName: ")
								fmt.Scanln(&Name)
								L[i] = Name + ":" + Lst[1]
								WriteIntoFile(File, strings.Join(L, "\n"), false)
								fmt.Printf("\n-----------------------Contact Name Changed from %v to %v-----------------------\n\n", name, Name)
								break
							} else if Mc == "p" {
								var Phone string
								fmt.Printf("Enter New Mobile Number: ")
								fmt.Scanln(&Phone)
								L[i] = Lst[0] + ":" + Phone
								WriteIntoFile(File, strings.Join(L, "\n"), false)
								fmt.Printf("\n-----------------------Contact Number Changed to %v of contact '%v'-----------------------\n\n", Phone, Lst[0])
								break
							} else {
								continue
							}
						}
					}
				}
				if temp == true {
					break
				} else {
					println("-------------------------Contact name is Not Found-----------------------------")
				}
			}
		} else if C == "s" || C == "S" {
			var S1 string
			for {
				fmt.Printf("Enter a Contact name: ")
				fmt.Scanln(&S1)
				S, found := Search(user+".txt", S1)
				if found == true {
					println("=========================================================================")
					println(S)
					println("=========================================================================")
					break
				} else {
					fmt.Printf("\n-----------------------------Contact name %v is Not Found-----------------------------\n", S1)
				}
			}
		} else if C == "d" || C == "D" {
			var S1 string
			for {
				fmt.Printf("Enter Contact Name to delete: ")
				fmt.Scanln(&S1)
				ok := Delete(user+".txt", S1)
				if ok == true {
					fmt.Println("Cona")
					break
				} else {
					fmt.Printf("\n-----------------------------Contact name %v is Not Found-----------------------------\n", S1)
				}
			}
		} else if C == "e" || C == "E" {
			break
		} else if C == "all" || C == "All" {
			S := ReadFile(user + ".txt")
			Lst := strings.Split(S, "\n")
			println("\n--------------------------------------------------------------\n")
			println("All Contacts List of User ", user)
			println("\n=============================\n")
			for i := 0; i < len(Lst)-1; i++ {
				list := strings.Split(Lst[i], ":")
				println("Contact Name: " + list[0] + "\nMobile Number: " + list[1] + "\n")
				println("=============================\n")
			}
			println("Number of Contacts = ", len(Lst)-1)
			println("---------------------------\n")
		}
	}
}

func NewUser() {
	var name, password, rePassword string
	var temp int = 0
	for {
		fmt.Printf("Enter UserName: ")
		fmt.Scanln(&name)
		S := ReadFile("contactInfo.txt")
		Lst := strings.Split(S, "\n")
		for i := 0; i < len(Lst)-1; i++ {
			lst := strings.Split(Lst[i], ":")
			if lst[0] == name {
				temp = 1
				println("---------------------------Your User name is already taken------------------")
				break
			}
		}
		if temp == 0 {
			break
		}
		temp = 0
	}
	for {
		fmt.Printf("Enter Password: ")
		fmt.Scanln(&password)
		fmt.Printf("Re-Enter Password: ")
		fmt.Scanln(&rePassword)
		if password == rePassword {
			CreateFile(name + ".txt")
			AppendIntoFile("contactInfo.txt", name+":"+password)
			println("------------------------Your Account Created Successfully-----------------------------")
			break
		}
		println("----------Error Password----------")
	}
	Command(name)
}

func OldUser() {
	var name, password string
	for {
		temp := true
		fmt.Printf("Enter Your Username: ")
		fmt.Scanln(&name)
		fmt.Printf("Enter password: ")
		fmt.Scanln(&password)
		userList := strings.Split(ReadFile("contactInfo.txt"), "\n")
		for i := 0; i < len(userList); i++ {
			USER := strings.Split(userList[i], ":")
			if USER[0] == name && USER[1] == password {
				Command(USER[0])
				temp = false
				break
			}
		}
		if temp == false {
			break
		} else {
			println("------------------------Username or password is incorrect---------------------------")
		}
	}
}

func Login() {
	for {
		var S string
		println("------------------------------WELCOME TO CLI(Command Line Interface) of TELEPHONE DIRECTORY---------------------------")
		fmt.Printf("Are you new user:(y/n)")
		fmt.Scanf("%v", &S)
		if S == "y" || S == "yes" {
			NewUser()
		} else if S == "n" || S == "no" {
			OldUser()
		} else {
			continue
		}
	}
}

func main() {
	Login()
}
