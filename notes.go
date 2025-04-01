package main

import "fmt"

type Note struct {
	Title    string
	Content  string
	Category string
}

var notes []Note

func createnote() {
	var num int
	var title, content, category string
	fmt.Print("How many notes do you want to enter:")
	fmt.Scan(&num)
	for i := 0; i < num; {
		fmt.Print("enter the Title: ")
		fmt.Scan(&title)
		fmt.Print("enter the content: ")
		fmt.Scan(&content)
		fmt.Print("enter the category: ")
		fmt.Scan(&category)

		if title == "" {
			fmt.Println("this cannot be empty please")
			continue
		}
		note := Note{Title: title, Content: content, Category: category}
		notes = append(notes, note)
		i++
		fmt.Println("Notes created successfully")
	}

}

func readnote() {
	if len(notes) == 0 {
		fmt.Println("no note available")
	} else {
		for i, note := range notes {
			fmt.Printf("%d. [%s]\n %s\n %s\n", i+1, note.Category, note.Title, note.Content)
		}
	}
}

func updatenote() {
	var index int
	readnote()
	fmt.Print("enter the number you want to update")
	fmt.Scan(&index)

	if index < 1 || index > len(notes) {
		fmt.Println("invalid note number")
		return
	}
	index--

	var title, content, category string

	fmt.Print("enter the new title:")
	fmt.Scan(&title)
	fmt.Print("enter the new content:")
	fmt.Scan(&content)
	fmt.Print("enter the new category:")
	fmt.Scan(&category)

	if title != "" {
		notes[index].Title = title
	}
	if content != "" {
		notes[index].Content = content
	}
	if category != "" {
		notes[index].Category = category
	}
}

func deletenote() {
	var index int
	readnote()
	fmt.Print("enter the number you want to delete")
	fmt.Scan(&index)

	if index < 1 || index > len(notes) {
		fmt.Println("invalid note number")
		return
	}
	index--

	notes = append(notes[:index], notes[index+1:]...)
	fmt.Println("note deleted succesfully ")

}

func main() {
	var option int

	for {
		fmt.Println("\n\n 1. Create a note")
		fmt.Println("2. Read a note")
		fmt.Println("3. Update a note")
		fmt.Println("4. Delete a note")
		fmt.Println("5. Exit the note app")
		fmt.Print("Enter your number:")
		fmt.Scan(&option)

		if option == 1 {
			createnote()
		} else if option == 2 {
			readnote()
		} else if option == 3 {
			updatenote()
		} else if option == 4 {
			deletenote()
		} else if option == 5 {
			fmt.Println("exiting the note app")
			return
		} else {
			fmt.Println("error")
		}
	}
}
