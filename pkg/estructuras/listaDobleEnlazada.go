package estructuras

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

type imagenes struct {
	nombre    string
	cantidad  int
	anterior  *imagenes
	siguiente *imagenes
}

type ListaDobleEnlazada struct {
	raiz     *imagenes
	ultimo   *imagenes
	longitud int
}

func (l *ListaDobleEnlazada) insertar(nombre string, cantidad int) {
	node := &imagenes{nombre: nombre, cantidad: cantidad, anterior: nil, siguiente: nil}
	if l.raiz == nil {
		l.raiz = node
		l.ultimo = node
		l.longitud++
	} else if l.raiz.siguiente == nil {
		l.raiz.siguiente = node
		l.ultimo = node
		l.ultimo.anterior = l.raiz
		l.longitud++
	} else {
		node.anterior = l.ultimo
		l.ultimo.siguiente = node
		l.ultimo = node
		l.longitud++
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
	nombreArchivo := "./listadoble.dot"
	nombreImagen := "./listadoble.jpg"
	texto := "digraph lista{\n"
	texto += "rankdir=LR;\n"
	texto += "node[shape = record];\n"
	texto += "nodonull1[label=\"null\"];\n"
	texto += "nodonull2[label=\"null\"];\n"
	aux := l.raiz
	contador := 0
	texto += "nodonull1->nodo0 [dir=back];\n"
	for i := 0; i < l.longitud; i++ {
		texto += "nodo" + strconv.Itoa(i) + "[label=\"" + aux.nombre + "\"];\n"
		aux = aux.siguiente
	}
	for i := 0; i < l.longitud-1; i++ {
		c := i + 1
		texto += "nodo" + strconv.Itoa(i) + "->nodo" + strconv.Itoa(c) + ";\n"
		texto += "nodo" + strconv.Itoa(c) + "->nodo" + strconv.Itoa(i) + ";\n"
		contador = c
	}
	texto += "nodo" + strconv.Itoa(contador) + "->nodonull2;\n"
	texto += "}"
	crearArchivo(nombreArchivo)
	escribirArchivo(texto, nombreArchivo)
	ejecutar(nombreImagen, nombreArchivo)
}
