package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	todos := []string{}
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("ToDoアプリへようこそ！")
	fmt.Println("コマンド: add, list, delete, exit")

	for {
		fmt.Print("> ")
		scanner.Scan()
		input := scanner.Text()
		args := strings.Fields(input)

		if len(args) == 0 {
			continue
		}

		command := args[0]

		switch command {
		case "add":
			if len(args) < 2 {
				fmt.Println("タスクの内容を入力してください")
				continue
			}
			task := strings.Join(args[1:], " ")
			todos = append(todos, task)
			fmt.Printf("タスクを追加しました: %s\n", task)

		case "list":
			if len(todos) == 0 {
				fmt.Println("タスクはありません")
				continue
			}
			fmt.Println("現在のタスク:")
			for i, todo := range todos {
				fmt.Printf("%d. %s\n", i+1, todo)
			}

		case "delete":
			if len(args) < 2 {
				fmt.Println("削除するタスクの番号を指定してください")
				continue
			}
			index := 0
			fmt.Sscanf(args[1], "%d", &index)
			if index < 1 || index > len(todos) {
				fmt.Println("無効な番号です")
				continue
			}
			todos = append(todos[:index-1], todos[index:]...)
			fmt.Println("タスクを削除しました")

		case "exit":
			fmt.Println("アプリケーションを終了します")
			return

		default:
			fmt.Println("無効なコマンドです")
		}
	}
} 