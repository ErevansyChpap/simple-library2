package main

type Storable interface{
	Save() error
	Load() error
}

//у нас пока что нет структуры Library