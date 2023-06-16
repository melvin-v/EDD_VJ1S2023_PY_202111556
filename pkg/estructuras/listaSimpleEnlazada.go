package estructuras

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

type empleados struct {
	id         string
	nombre     string
	cargo      string
	contraseña string
	siguiente  *empleados
}

type ListaSimpleEnlazada struct {
	raiz     *empleados
	ultimo   *empleados
	longitud int
}

func (l *ListaSimpleEnlazada) insertar(id string, nombre string, cargo string, contraseña string) {
	node := &empleados{id: id, nombre: nombre, cargo: cargo, contraseña: contraseña, siguiente: nil}
	if l.raiz == nil {
		l.raiz = node
		l.ultimo = node
		l.longitud++
	} else if l.raiz.siguiente == nil {
		l.raiz.siguiente = node
		l.ultimo = node
		l.longitud++
	} else {
		l.ultimo.siguiente = node
		l.ultimo = node
		l.longitud++
	}
}

func (l *ListaSimpleEnlazada) LeerCSV(ruta string) {
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
		l.insertar(linea[0], linea[1], linea[2], linea[3])
	}
}

func (l *ListaSimpleEnlazada) Reporte() {
	nombreArchivo := "./listasimple.dot"
	nombreImagen := "./listasimple.jpg"
	texto := "digraph lista{\n"
	texto += "rankdir=LR;\n"
	texto += "node[shape = record];\n"
	aux := l.raiz
	for i := 0; i < l.longitud; i++ {
		texto += "nodo" + strconv.Itoa(i) + "[label=\"{ID: " + aux.id + "\\n" + "Nombre: " + aux.nombre + "|}\"];\n"
		aux = aux.siguiente
	}
	for i := 0; i < l.longitud-1; i++ {
		c := i + 1
		texto += "nodo" + strconv.Itoa(i) + "->nodo" + strconv.Itoa(c) + "[dir=forward];\n"
	}

	texto += "}"
	crearArchivo(nombreArchivo)
	escribirArchivo(texto, nombreArchivo)
	ejecutar(nombreImagen, nombreArchivo)
}

func (l *ListaSimpleEnlazada) Loging(id string, password string) bool {
	aux := l.raiz
	flag := false
	for aux != nil {
		if aux.id == id && aux.contraseña == password {
			flag = true
			break
		}
		aux = aux.siguiente
	}
	return flag
}
