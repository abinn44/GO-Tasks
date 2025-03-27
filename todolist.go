package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
	"os"
)

var students []string
var scanner = bufio.NewScanner(os.Stdin)

func addStudent() {
	fmt.Println("Enter the student name")
	scanner.Scan()
	name := scanner.Text()
	students = append(students, name)
	fmt.Println("Student added Successfully")
}

func listStudents() {
	fmt.Println("Students List:")
	for i, students := range students {
		fmt.Printf("| ID: %d | Name: %s |\n", i+1, students)
	}
}

func deleteStudent() {
	fmt.Println("Enter student ID to delete")
	scanner.Scan()
	i, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))
	students = append(students[:i-1], students[i:]...)
	fmt.Println("Student deleted successfully")
}

func updateStudent() {
	fmt.Println("Enter student ID to Update")
	scanner.Scan()
	i, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))

	fmt.Println("Enter new name")
	scanner.Scan()

	newName := scanner.Text()
	students[i-1] = newName
	fmt.Println("Student updated successfully")
}

func searchStudent() {
	fmt.Println("Enter student name to search")
	scanner.Scan()
	searchName := strings.ToLower(scanner.Text())

	found := false

	for i, student := range students {
		if strings.ToLower(student) == searchName {
			fmt.Printf("Student Found. ID: %d | Name: %s\n", i+1, student)

			found = true
			break
		}
	}
	if !found {
		fmt.Println("Student Not found")
	}
}

func main() {
	for {
		fmt.Println("Enter your choice")
		fmt.Println("1 Add student")
		fmt.Println("2 List student")
		fmt.Println("3 Update student")
		fmt.Println("4 Delete student")
		fmt.Println("5 Search student")
		fmt.Println("6 Exit")

		scanner.Scan()
		choice, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))

		switch choice {
		case 1:
			addStudent()
		case 2:
			listStudents()
		case 3:
			updateStudent()
		case 4:
			deleteStudent()
		case 5:
			searchStudent()
		case 6:
			fmt.Println("Exiting..")
			return
		}
	}
}