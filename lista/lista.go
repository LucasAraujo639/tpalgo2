package lista

type Lista[T any] interface {
	EstaVacia() bool
	InsertarPrimero(T)
	InsertarUltimo(T)
	BorrarPrimero() T
	VerPrimero() T
	VerUltimo() T
	Largo() int

	Iterar(visitar func(T) bool) //It interno
	Iterador() IteradorLista[T]  // It externo
}

type IteradorLista[T any] interface {

	//Primitivas iterador externo
	VerActual() T
	HaySiguiente() bool
	Siguiente()
	Insertar(T)
	Borrar() T
}
