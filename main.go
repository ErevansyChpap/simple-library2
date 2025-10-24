package main

import "fmt"

func main() {
	fmt.Println("Запуск системы управления библиотекой...")

	myLibrary := &Library{}

	fmt.Println("\n--- Наполняю библиотеку ---")

	reader, err := myLibrary.AddReader("Владимир", "Асеев")
	if err != nil{
		fmt.Println(err)
	} else {
		fmt.Println("Зарегестрирован новый читатель:", reader)
	}
	reader, err = myLibrary.AddReader("Алан", "Кусраев")
	if err != nil{
		fmt.Println(err)
	} else {
		fmt.Println("Зарегестрирован новый читатель:", reader)
	}

	b, err := myLibrary.AddBook(1867, "Война и мир", "Лев Толстой")
	if err != nil{
		fmt.Println(err)
	} else {
		fmt.Printf("Книга (%v) успешно добавлена.", b)
	}
	b, err = myLibrary.AddBook(1835, "Мертвые души", "Николай Гоголь")
	if err != nil{
		fmt.Println(err)
	} else {
		fmt.Printf("\nКнига (%v) успешно добавлена.", b)
	}
	b, err = myLibrary.AddBook(1925, "Собачье сердце", "Михаил Булгаков")
	if err != nil{
		fmt.Println(err)
	} else {
		fmt.Printf("\nКнига (%v) успешно добавлена.\n", b)
	}
	
	fmt.Println("----Книга успешно выдана----")
	err = myLibrary.IssueBookToReader(1, 1)
	if err != nil{
		fmt.Println(err)
	} else {
		fmt.Println(myLibrary.Books[1])
	}
	fmt.Println("----Книга уже выдана----")
	err = myLibrary.IssueBookToReader(1, 1)
	if err != nil{
		fmt.Println(err)
	} else {
		fmt.Println(myLibrary.Books[1])
	}
	fmt.Println("----Такого читателя нет----")
	err = myLibrary.IssueBookToReader(2, 15)
	if err != nil{
		fmt.Println(err)
	} else {
		fmt.Println(myLibrary.Books[1])
	}
	fmt.Println("----Успешный возврат книги----")
	err = myLibrary.ReturnBook(1)
	if err != nil{
		fmt.Println(err)
	} else {
		fmt.Println("Статус книги после возварат:", myLibrary.Books[1])
	}
	fmt.Println("----Книга уже в библиотеке----")
	err = myLibrary.ReturnBook(1)
	if err != nil{
		fmt.Println(err)
	} else {
		fmt.Println("Статус книги после возварата:", myLibrary.Books[1])
	}
}