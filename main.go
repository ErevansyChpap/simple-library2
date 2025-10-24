package main

import "fmt"

func main() {
	fmt.Println("Запуск системы управления библиотекой...")

	myLibrary := &Library{}

	fmt.Println("\n--- Наполняем библиотеку ---")
	myLibrary.AddReader("Тамерлан", "Джигкаев")
	myLibrary.AddReader("Линда", "Элбакянц")

	myLibrary.AddBook(1984, "1984", "Джордж Оруэлл")
	myLibrary.AddBook(1967, "Мастер и Маргарита", "Михаил Булгаков")

	fmt.Println("\n--- Библиотека готова к работе ---")
	fmt.Println("Количество читателей:", len(myLibrary.Readers))
	fmt.Println("Количество книг:", len(myLibrary.Books))

	fmt.Println("---Тестируем выдачу книг---")
	//Выдаем книгу 1 читателю 1
	err := myLibrary.IssueBookToReader(1, 1)
	if err != nil {
		fmt.Println("Ошибка выдачи", err)
	}

	//Проверить статус книги после выдачи
	book, _ := myLibrary.FindBookById(1)
	if book != nil {
		fmt.Println("Статус книги после выдачи:", book)
	}

	//Тестирование возврата книги
	test := myLibrary.ReturnBook(1)
	if test != nil {
		fmt.Println(test)
	}

	if book != nil {
		fmt.Println("Статус книги после возврата:", book)
	}
}