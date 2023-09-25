package lista

type nodoLista[T any] struct {
	dato      T
	siguiente *nodoLista[T]
}

type listaEnlazada[T any] struct {
	primero *nodoLista[T]
	ultimo  *nodoLista[T]
	largo   int
}

func CrearListaEnlazada[T any]() Lista[T] {
	return &listaEnlazada[T]{
		primero: nil,
		ultimo:  nil,
	}
}

// EstaVacia() bool
// 	InsertarPrimero(T)
// 	InsertarUltimo(T)
// 	BorrarPrimero() T
// 	VerPrimero() T
// 	VerUltimo() T
// 	Largo() int

// 	//Primitivas iterador externo
// 	Iterar(visitar func(T) bool)
// 	Iterador() IteradorLista[T]
