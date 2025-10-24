package main

import "fmt"

//Структура книги
type Book struct {
	ID       int
	Year     int
	Title    string
	Author   string
	IsIssued bool
	ReaderId *int
}


//Метод для деактивации читателя
func (r *Reader) Deactivate() error{
	if !r.IsActive {
		return fmt.Errorf("%v и так неактивен", r)
	} else{
		r.IsActive = false
	}
	return nil
}

//Метод для красивого выводы читателя
func (r Reader) String() string {
	status := ""
	if r.IsActive {
		status = "активен."
	} else {
		status = "неактивен."
	}
	return fmt.Sprintf("Пользователь %s %s, ID: %d, пользователь %s", r.FirstName, r.LastName, r.ID, status)
}

//Метод для красивого вывода книга
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
//Метод проверяющий используется ли книга
func (b *Book) IssueBook(r *Reader) error {
	if b.IsIssued {
		return fmt.Errorf("Книга уже используется.")
	} else {
		b.IsIssued = true
		b.ReaderId = &r.ID
	}
	return nil
}

//Метод возвращающий книгу
func (b *Book) ReturnBook() error {
	if !b.IsIssued {
		return fmt.Errorf("книга '%s' и так в библиотеке", b.Title)
	} else {
		b.IsIssued = false
		b.ReaderId = nil
	}
	return nil
}

//Структура библиотеки
type Library struct {
	Books   []*Book
	Readers []*Reader

	lastBookID   int
	lastReaderID int
}

//Метод добавляющтй читателя в библиотеку
func (l *Library) AddReader(firstName, lastName string) (*Reader, error) {
	if firstName == "" || lastName == ""{
		return nil, fmt.Errorf("Имя или фамилия не могут быть пустыми")
	}
	l.lastReaderID++
	newReader := &Reader{
		ID:        l.lastReaderID,
		FirstName: firstName,
		LastName:  lastName,
		IsActive:  true,
	}
	l.Readers = append(l.Readers, newReader)
	return newReader, nil
}

//Метод добавляющий книгу в библиотеку
func (l *Library) AddBook(year int, title, author string) (*Book, error) {
	for _, value := range l.Books {
		if value.Author == author && value.Title == title {
			return nil, fmt.Errorf("Такая книга уже есть в библиотеке.")
		}
	}
	l.lastBookID++
	newBook := &Book{
		ID:       l.lastBookID,
		Year:     year,
		Title:    title,
		Author:   author,
		IsIssued: false,
	}
	l.Books = append(l.Books, newBook)
	return newBook, nil
}

//Метод ищущий читателя по ID
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

//Метод ищущий книгу по ID
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

//Метод для выдачи книги читателю
func (l *Library) IssueBookToReader(bookId, readerId int) error {
	book, err := l.FindBookById(bookId)
	if book == nil {
		return err
	}
	reader, err := l.FindReaderById(readerId)
	if reader == nil {
		return err
	}
	err = book.IssueBook(reader)
	if err != nil{
		return err
	}
	return nil
}

//Метод возвращающий книгу по ID
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

//Структура читателя
type Reader struct {
	ID        int
	FirstName string
	LastName  string
	IsActive  bool
}