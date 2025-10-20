package main

import "fmt"

type Library struct {
	Books   []*Book
	Readers []*Reader

	LastBookID   int
	LastReaderID int
}

type Book struct {
	ID       int
	Year     int
	Title    string
	Author   string
	IsIssued bool
	ReaderId *int
}

func (lib *Library) AddReader(firstName, lastName string) *Reader {
	lib.LastReaderID++

	newReader := &Reader{
		ID:        lib.LastReaderID,
		FirstName: firstName,
		LastName:  lastName,
		IsActive:  true,
	}

	lib.Readers = append(lib.Readers, newReader)

	fmt.Printf("Зарегистрирован новый читатель: %s\n", newReader)
	return newReader
}

func (lib *Library) AddBook(title, author string, year int) *Book {
	lib.LastBookID++

	newBook := &Book{
		ID:       lib.LastBookID,
		Title:    title,
		Author:   author,
		Year:     year,
		IsIssued: false,
	}

	lib.Books = append(lib.Books, newBook)

	fmt.Printf("Добавлена новая книга: %s\n", newBook)
	return newBook
}

func (l *Library) FindBookByID(id int) (*Book, error) {
	for i := 0; i < len(l.Books); i++ {
		if i == id {
			return l.Books[i], nil
		}
	}
	return nil, fmt.Errorf("книга с ID %d не найдена", id)
}

func (l *Library) FindReaderByID(id int) (*Reader, error) {
	for i := 0; i < len(l.Readers); i++ {
		if i == id {
			return l.Readers[i], nil
		}
	}
	return nil, fmt.Errorf("читатель с ID %d не найден", id)
}

func (l *Library) IssueBookToReader(bookID int, readerID int) error {
	book, err := l.FindBookByID(bookID)
	if book == nil {
		return err
	}
	reader, err2 := l.FindReaderByID(readerID)
	if reader == nil {
		return err2
	}
	book.IssueBook(reader)
	return nil
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

func (b *Book) ReturnBook() {
	if !b.IsIssued {
		fmt.Println("Книга уже в библиотеке.")
	} else {
		b.IsIssued = false
		b.ReaderId = nil
		fmt.Println("Книга возвращена.")
	}
}

func (r *Reader) AssignBook(b *Book) {
	fmt.Printf("Читатель %s %s взял книгу %s(%s %d)\n", r.FirstName, r.LastName, b.Title, b.Author, b.Year)
}

type Reader struct {
	ID        int
	FirstName string
	LastName  string
	IsActive  bool
}
