package estructuras

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

type imagenes struct {
	Nombre    string
	cantidad  int
	anterior  *imagenes
	Siguiente *imagenes
}

type ListaDobleEnlazada struct {
	Raiz     *imagenes
	ultimo   *imagenes
	Longitud int
}

func (l *ListaDobleEnlazada) insertar(Nombre string, cantidad int) {
	node := &imagenes{Nombre: Nombre, cantidad: cantidad, anterior: nil, Siguiente: nil}
	if l.Raiz == nil {
		l.Raiz = node
		l.ultimo = node
		l.Longitud++
	} else if l.Raiz.Siguiente == nil {
		l.Raiz.Siguiente = node
		l.ultimo = node
		l.ultimo.anterior = l.Raiz
		l.Longitud++
	} else {
		node.anterior = l.ultimo
		l.ultimo.Siguiente = node
		l.ultimo = node
		l.Longitud++
	}
}

func (l *ListaDobleEnlazada) LeerCSV(ruta string) {
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
		intVar1, _ := strconv.Atoi(linea[1])
		l.insertar(linea[0], intVar1)
	}
}

func (l *ListaDobleEnlazada) Reporte() {
	NombreArchivo := "./listadoble.dot"
	NombreImagen := "./listadoble.jpg"
	texto := "digraph lista{\n"
	texto += "rankdir=LR;\n"
	texto += "node[shape = record];\n"
	texto += "nodonull1[label=\"null\"];\n"
	texto += "nodonull2[label=\"null\"];\n"
	aux := l.Raiz
	contador := 0
	texto += "nodonull1->nodo0 [dir=back];\n"
	for i := 0; i < l.Longitud; i++ {
		texto += "nodo" + strconv.Itoa(i) + "[label=\"" + aux.Nombre + "\"];\n"
		aux = aux.Siguiente
	}
	for i := 0; i < l.Longitud-1; i++ {
		c := i + 1
		texto += "nodo" + strconv.Itoa(i) + "->nodo" + strconv.Itoa(c) + ";\n"
		texto += "nodo" + strconv.Itoa(c) + "->nodo" + strconv.Itoa(i) + ";\n"
		contador = c
	}
	texto += "nodo" + strconv.Itoa(contador) + "->nodonull2;\n"
	texto += "}"
	crearArchivo(NombreArchivo)
	escribirArchivo(texto, NombreArchivo)
	ejecutar(NombreImagen, NombreArchivo)
}
