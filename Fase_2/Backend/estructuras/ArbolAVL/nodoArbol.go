package ArbolAVL

import "Backend/estructuras/Peticiones"

type NodoArbol struct {
	Izquierdo         *NodoArbol
	Derecho           *NodoArbol
	Valor             *Peticiones.Pedido
	Altura            int
	Factor_Equilibrio int
}
