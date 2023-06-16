package estructuras

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

type nodoCola struct {
	id        string
	nombre    string
	Siguiente *nodoCola
}

type Cola struct {
	primero  *nodoCola
	longitud int
}

func (c *Cola) encolar(id string, nombre string) {
	node := &nodoCola{id: id, nombre: nombre}
	if c.longitud == 0 {
		c.primero = node
		c.longitud++
	} else {
		aux := c.primero
		for aux.Siguiente != nil {
			aux = aux.Siguiente
		}
		aux.Siguiente = node
		c.longitud++
	}

}

func (c *Cola) LeerCSV(ruta string) {
	file, err := os.Open(ruta)
	if err != nil {
		fmt.Println("No pude abrir el archivo")
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

		c.encolar(linea[0], linea[1])
	}
}

func (c *Cola) Reporte() {
	nombre_archivo := "./cola.dot"
	nombre_imagen := "cola.jpg"
	texto := "digraph cola{\n"
	texto += "rankdir=LR;\n"
	texto += "node[shape = record];\n"
	texto += "nodonull2[label=\"null\"];\n"
	aux := c.primero
	contador := 0
	for i := 0; i < c.longitud; i++ {
		texto = texto + "nodo" + strconv.Itoa(i) + "[label=\"{Nombre: " + aux.nombre + "\\n" + "Id: " + aux.id + "|}\"];\n"
		aux = aux.Siguiente
	}
	for i := 0; i < c.longitud-1; i++ {
		c := i + 1
		texto += "nodo" + strconv.Itoa(i) + "->nodo" + strconv.Itoa(c) + ";\n"
		contador = c
	}
	texto += "nodo" + strconv.Itoa(contador) + "->nodonull2;\n"
	texto += "}"
	crearArchivo(nombre_archivo)
	escribirArchivo(texto, nombre_archivo)
	ejecutar(nombre_imagen, nombre_archivo)
}
