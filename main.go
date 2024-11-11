package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	type TODOs struct {
		Id          int    `json:"id"`
		Description string `json:"description"`
		Status      bool   `json:"status"`
	}
	counter := 0
	data := make([]TODOs, 0)
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("CMDs: show create <description> remove <id> done <id>")

	for {
		fmt.Print(">")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		split := strings.Fields(line)

		if len(split) == 0 {
			continue
		}
		switch split[0] {
		case "create":
			if len(split) < 2 {
				fmt.Println("error: Description is required to create a TODO item.")
				continue
			}
			data = append(data, TODOs{
				Id:          counter,
				Description: strings.Join(split[1:], " "),
				Status:      false,
			})
			counter++
			fmt.Println("Created TODO item")
		case "remove":
			if len(split) != 2 {
				fmt.Println("Error: You need to provide an ID to remove a TODO item")
				continue
			}
			var id int
			_, err := fmt.Sscanf(split[1], "%d", &id)
			if err != nil || id < 0 || id >= len(data) {
				fmt.Println("Error: Invalid TODO ID.")
				continue
			}
			data = append(data[:id], data[id+1:]...)
			fmt.Printf("Removed TODO item with ID %d.\n", id)

		case "show":
			if len(data) == 0 {
				fmt.Println("No TODO items.")
			} else {
				for _, todo := range data {
					status := "Not Done"
					if todo.Status {
						status = "Done"
					}
					fmt.Printf("ID: %d, Description: %s, Status: %s\n", todo.Id, todo.Description, status)
				}
			}
		case "done":
			if len(split) != 2 {
				fmt.Println("Error: You need to provide an ID to mark the TODO as done.")
				continue
			}
			var id int
			_, err := fmt.Sscanf(split[1], "%d", &id)
			if err != nil || id < 0 || id >= len(data) {
				fmt.Println("Error: Invalid TODO ID")
				continue
			}
			data[id].Status = true
			fmt.Printf("Marked TODO item with ID %d as done.\n", id)
		default:
			fmt.Println("Unknown Command. Valid commands are: show, create, remove, done.")
		}
	}
}
