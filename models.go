package main

import "fmt"

type Book struct {
	ID int
	Year int
	Title string
	Author string
	IsIssued bool
	ReaderId *int
}

func (r Reader) DisplayReader() {
	fmt.Printf("Читатель: %s %s (ID: %d)(Status: %v)\n", r.FirstName, r.LastName, r.ID, r.IsActive)
}

 func (r *Reader) Deactivate(){
	r.IsActive = false
}

func (r Reader) String() string{
	status := ""
	if r.IsActive {
		status = "активен."
	} else {
		status = "неактивен."
	}
	return fmt.Sprintf("Пользователь %s %s, ID: %d, пользователь %s", r.FirstName, r.LastName, r.ID, status)
}

func (b Book) String() string{
	status := ""
	if b.IsIssued {
		status = "используется."
		return fmt.Sprintf("ID: %d, %s (%s %d), книга %s читателем %d", b.ID, b.Title, b.Author, b.Year, status, *b.ReaderId)
	} else {
		status = "не используется."
		return fmt.Sprintf("ID: %d, %s (%s %d), книга %s", b.ID, b.Title, b.Author, b.Year, status)
	}
	
}

func (b *Book) IssueBook(r *Reader) {
	if b.IsIssued{
		fmt.Println("Книга уже используется.")
	} else {
		b.IsIssued = true
		b.ReaderId = &r.ID
		fmt.Printf("Книга выдана читателю %s %s.\n", r.FirstName, r.LastName)
	}
}

func (b *Book) ReturnBook() {
	if !b.IsIssued{
		fmt.Println("Книга уже в библиотеке.")
	} else {
		b.IsIssued = false
		b.ReaderId = nil
		fmt.Println("Книга возвращена.")
	}
}

func (r *Reader) AssignBook(b *Book) {
	fmt.Printf("Читатель %s %s взял книгу %s(%s %d)", r.FirstName, r.LastName, b.Title, b.Author, b.Year)
}

type Reader struct {
	ID        int
	FirstName string
	LastName  string
	IsActive  bool
}
