package main

import "fmt"

type Book struct {
	ID       int
	Year     int
	Title    string
	Author   string
	IsIssued bool
	ReaderId *int
}

func (r *Reader) AssignBook(b *Book) {
	fmt.Printf("Читатель %s %s взял книгу %s(%s %d)\n", r.FirstName, r.LastName, b.Title, b.Author, b.Year)
}

func (r Reader) DisplayReader() {
	fmt.Printf("Читатель: %s %s (ID: %d)(Status: %v)\n", r.FirstName, r.LastName, r.ID, r.IsActive)
}

func (r *Reader) Deactivate() {
	r.IsActive = false
}

func (r Reader) String() string {
	status := ""
	if r.IsActive {
		status = "активен."
	} else {
		status = "неактивен."
	}
	return fmt.Sprintf("Пользователь %s %s, ID: %d, пользователь %s", r.FirstName, r.LastName, r.ID, status)
}

func (b Book) String() string {
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
	if b.IsIssued {
		fmt.Println("Книга уже используется.")
	} else {
		b.IsIssued = true
		b.ReaderId = &r.ID
		fmt.Printf("Книга выдана читателю %s %s.\n", r.FirstName, r.LastName)
	}
}

func (b *Book) ReturnBook() error {
	if !b.IsIssued {
		return fmt.Errorf("книга '%s' и так в библиотеке", b.Title)
	} else {
		b.IsIssued = false
		b.ReaderId = nil
	}
	return nil
}

type Library struct {
	Books   []*Book
	Readers []*Reader

	lastBookID   int
	lastReaderID int
}

func (l *Library) AddReader(firstName, lastName string) *Reader {
	l.lastReaderID++
	newReader := &Reader{
		ID:        l.lastReaderID,
		FirstName: firstName,
		LastName:  lastName,
		IsActive:  true,
	}
	l.Readers = append(l.Readers, newReader)
	fmt.Printf("Зарегистрирован новый читатель: %s\n", newReader)
	return newReader
}

func (l *Library) AddBook(year int, title, author string) *Book {
	l.lastBookID++

	newBook := &Book{
		ID:       l.lastBookID,
		Year:     year,
		Title:    title,
		Author:   author,
		IsIssued: false,
	}
	l.Books = append(l.Books, newBook)
	fmt.Printf("Добавлена новая книга: %s\n", newBook)
	return newBook
}

func (l *Library) FindBookById(id int) (*Book, error) {
	flag := false
	for i := 0; i < len(l.Books); i++ {
		if i == id {
			flag = true
		}
	}
	if flag {
		return l.Books[id], nil
	}
	return nil, fmt.Errorf("книга с ID %d не найдена", id)
}

func (l *Library) FindReaderById(id int) (*Reader, error) {
	flag := false
	for i := 0; i < len(l.Readers); i++ {
		if i == id-1 {
			flag = true
		}
	}
	if flag {
		return l.Readers[id-1], nil
	}
	return nil, fmt.Errorf("читатель с ID %d не найден", id)
}

func (l *Library) IssueBookToReader(bookId, readerId int) error {
	book, err := l.FindBookById(bookId)
	if book == nil {
		return err
	}
	reader, err := l.FindReaderById(readerId)
	if reader == nil {
		return err
	}
	book.IssueBook(reader)
	return nil
}

func (l *Library) ReturnBook(bookId int) error {
	book, err := l.FindBookById(bookId)
	if err != nil {
		return err
	}
	test := book.ReturnBook()
	if test != nil {
		return test
	}
	return nil
}

type Reader struct {
	ID        int
	FirstName string
	LastName  string
	IsActive  bool
}