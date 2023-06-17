package estructuras

import (
	"encoding/csv"
	"fmt"
	"io"

	"os"
	"strconv"
)

type nodoCola struct {
	Id        string
	Nombre    string
	Siguiente *nodoCola
}

type Cola struct {
	Primero  *nodoCola
	longitud int
}

func (c *Cola) Encolar(Id string, Nombre string) {
	node := &nodoCola{Id: Id, Nombre: Nombre}
	if c.longitud == 0 {
		c.Primero = node
		c.longitud++
	} else {
		aux := c.Primero
		for aux.Siguiente != nil {
			aux = aux.Siguiente
		}
		aux.Siguiente = node
		c.longitud++
	}

}

func (c *Cola) Desencolar() {
	c.Primero = c.Primero.Siguiente
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

		c.Encolar(linea[0], linea[1])
	}
}

func (c *Cola) Reporte() {
	Nombre_archivo := "./cola.dot"
	Nombre_imagen := "cola.jpg"
	texto := "digraph cola{\n"
	texto += "rankdir=LR;\n"
	texto += "node[shape = record];\n"
	texto += "nodonull2[label=\"null\"];\n"
	aux := c.Primero
	contador := 0
	for i := 0; i < c.longitud; i++ {
		texto = texto + "nodo" + strconv.Itoa(i) + "[label=\"{Nombre: " + aux.Nombre + "\\n" + "Id: " + aux.Id + "|}\"];\n"
		aux = aux.Siguiente
	}
	for i := 0; i < c.longitud-1; i++ {
		c := i + 1
		texto += "nodo" + strconv.Itoa(i) + "->nodo" + strconv.Itoa(c) + ";\n"
		contador = c
	}
	texto += "nodo" + strconv.Itoa(contador) + "->nodonull2;\n"
	texto += "}"
	crearArchivo(Nombre_archivo)
	escribirArchivo(texto, Nombre_archivo)
	ejecutar(Nombre_imagen, Nombre_archivo)
}

func (c *Cola) SiExiste(IdCliente string) bool {
	flag := false
	aux := c.Primero
	for aux.Siguiente != nil {
		if aux.Id == IdCliente {
			flag = true
		}
		aux = aux.Siguiente
	}

	return flag
}
