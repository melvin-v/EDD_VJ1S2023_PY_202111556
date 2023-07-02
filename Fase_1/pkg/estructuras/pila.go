package estructuras

import "fmt"

type nodoPila struct {
	idCliente    string
	idEmpleado   string
	nombreImagen string
	Siguiente    *nodoPila
}

type Pila struct {
	Primero  *nodoPila
	Longitud int
}

func (p *Pila) Push(idCliente string, idEmpleado string, nombreImagen string) {
	if p.Longitud == 0 {
		nuevoNodo := &nodoPila{idCliente: idCliente, idEmpleado: idEmpleado, nombreImagen: nombreImagen, Siguiente: nil}
		p.Primero = nuevoNodo
		p.Longitud++
	} else {
		nuevoNodo := &nodoPila{idCliente: idCliente, idEmpleado: idEmpleado, nombreImagen: nombreImagen, Siguiente: p.Primero}
		p.Primero = nuevoNodo
		p.Longitud++
	}
}

func (p *Pila) Pop() {
	if p.Longitud == 0 {
		fmt.Println("No hay elementos en la pila")
	} else {
		p.Primero = p.Primero.Siguiente
		p.Longitud--
	}
}

func (p *Pila) Reporte() {
	nombre_archivo := "./pila.dot"
	nombre_imagen := "pila.jpg"
	texto := "digraph pila{\n"
	texto += "rankdir=LR;\n"
	texto += "node[shape = record]"
	aux := p.Primero
	texto += "nodo0 [label=\""
	for i := 0; i < p.Longitud; i++ {
		texto = texto + "Carnet: " + aux.idCliente + "\\n" + "Imagen: " + aux.nombreImagen
		aux = aux.Siguiente
	}
	texto += "\"]; \n}"
	crearArchivo(nombre_archivo)
	escribirArchivo(texto, nombre_archivo)
	ejecutar(nombre_imagen, nombre_archivo)
}

func (p *Pila) JSON() {
	nombre_archivo := "./pila.json"
	texto := "{\"pedidos\":["
	aux := p.Primero
	for i := 0; i < p.Longitud; i++ {
		texto = texto + "{\"id_cliente\": " + "\"" + aux.idCliente + "\"" + "," + "\"imagen\": " + "\"" + aux.nombreImagen + "\"" + "}"
		aux = aux.Siguiente
	}
	texto += "]}"
	crearArchivo(nombre_archivo)
	escribirArchivo(texto, nombre_archivo)
}
