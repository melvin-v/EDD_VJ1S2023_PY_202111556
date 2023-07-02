package estructuras

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

type clientes struct {
	id        string
	Nombre    string
	Siguiente *clientes
}

type ListaCircularSimpleEnlazada struct {
	Raiz     *clientes
	ultimo   *clientes
	Longitud int
}

func (l *ListaCircularSimpleEnlazada) Insertar(Nombre string, id string) {
	node := &clientes{id: id, Nombre: Nombre}
	if l.Raiz == nil {
		l.Raiz = node
		l.Raiz.Siguiente = node
		l.ultimo = node
		l.Longitud++
	} else if l.Raiz.Siguiente == nil {
		l.Raiz.Siguiente = node
		l.ultimo = node
		l.ultimo.Siguiente = l.Raiz
		l.Longitud++
	} else {
		l.ultimo.Siguiente = node
		l.ultimo = node
		l.ultimo.Siguiente = l.Raiz
		l.Longitud++
	}
}

func (l *ListaCircularSimpleEnlazada) LeerCSV(ruta string) {
	file, err := os.Open(ruta)
	if err != nil {
		fmt.Println("No pude abrir el archivo")
		return
	}
	defer file.Close()

	lectura := csv.NewReader(file)
	lectura.Comma = ','
	encabezado := true
	for {
		linea, err := lectura.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("No pude leer la linea del csv")
			continue
		}
		if encabezado {
			encabezado = false
			continue
		}
		l.Insertar(linea[0], linea[1])
	}
}

func (l *ListaCircularSimpleEnlazada) Reporte() {
	NombreArchivo := "./listacircular.dot"
	NombreImagen := "./listacircular.jpg"
	texto := "digraph lista{\n"
	texto += "rankdir=LR;\n"
	texto += "node[shape = record];\n"
	aux := l.Raiz
	contador := 0
	for i := 0; i < l.Longitud; i++ {
		texto += "nodo" + strconv.Itoa(i) + "[label=\"{ID: " + aux.id + "\\n" + "Nombre: " + aux.Nombre + "|}\"];\n"
		aux = aux.Siguiente
	}
	for i := 0; i < l.Longitud-1; i++ {
		c := i + 1
		texto += "nodo" + strconv.Itoa(i) + "->nodo" + strconv.Itoa(c) + "[dir=forward];\n"
		contador = c
	}
	texto += "nodo" + strconv.Itoa(contador) + "->nodo0;\n"

	texto += "}"
	crearArchivo(NombreArchivo)
	escribirArchivo(texto, NombreArchivo)
	ejecutar(NombreImagen, NombreArchivo)
}

func (l *ListaCircularSimpleEnlazada) SiExiste(IdCliente string) bool {
	flag := false
	aux := l.Raiz
	for aux.Siguiente != nil {
		if aux.id == IdCliente {
			flag = true
		}
		aux = aux.Siguiente
	}

	return flag
}
