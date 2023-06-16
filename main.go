package main

import (
	"EDD_creative/pkg/estructuras"
	"fmt"
	"strconv"
)

func main() {

	matriz_csv := &estructuras.Matriz{Raiz: &estructuras.NodoMatriz{PosX: -1, PosY: -1, Color: "RAIZ"}}
	imagen := "mario"
	archivo := "body.csv"
	matriz_csv.LeerArchivo("csv/" + imagen + "/" + archivo)
	matriz_csv.Reporte()

	listaEmpleados := estructuras.ListaSimpleEnlazada{}
	listaImagenes := estructuras.ListaCircularSimpleEnlazada{}
	listaUsuarios := estructuras.ListaDobleEnlazada{}
	cola := estructuras.Cola{}
	pila := estructuras.Pila{}

	opcion := "0"
	exit1 := true
	for exit1 {
		fmt.Println("**********************************************")
		fmt.Println("*                 MENÚ PRINCIPAL             *")
		fmt.Println("**********************************************")
		fmt.Println("*   1. Iniciar sesion                        *")
		fmt.Println("*   2. Salir del sistema                     *")
		fmt.Println("**********************************************")
		fmt.Scanln(&opcion)
		switch opcion {
		case "1":
			usuario := ""
			contraseña := ""
			fmt.Println("**********************************************")
			fmt.Println("*                 LOGIN                      *")
			fmt.Println("**********************************************")
			fmt.Print("Ingrese el usuario: ")
			fmt.Scanln(&usuario)
			fmt.Print("Ingrese la contraseña: ")
			fmt.Scanln(&contraseña)

			if usuario == "1" && contraseña == "1" {
				opAdmin := "0"
				exit2 := true
				for exit2 {
					fmt.Println("**********************************************")
					fmt.Println("*          DASHBOARD ADMIN 202111556         *")
					fmt.Println("**********************************************")
					fmt.Println("*   1. Cargar Empleados                      *")
					fmt.Println("*   2. Cargar Imagenes                       *")
					fmt.Println("*   3. Cargar Usuarios                       *")
					fmt.Println("*   4. Actualizar cola                       *")
					fmt.Println("*   5. Reporte estructuras                   *")
					fmt.Println("*   6. Desloguearse                          *")
					fmt.Println("**********************************************")
					fmt.Scanln(&opAdmin)

					switch opAdmin {
					case "1":
						ruta := ""
						fmt.Print("Ingrese la ruta: ")
						fmt.Scanln(&ruta)
						listaEmpleados.LeerCSV(ruta)

					case "2":
						ruta := ""
						fmt.Print("Ingrese la ruta: ")
						fmt.Scanln(&ruta)
						listaImagenes.LeerCSV(ruta)

					case "3":
						ruta := ""
						fmt.Print("Ingrese la ruta: ")
						fmt.Scanln(&ruta)
						listaUsuarios.LeerCSV(ruta)

					case "4":
						ruta := ""
						fmt.Print("Ingrese la ruta: ")
						fmt.Scanln(&ruta)
						cola.LeerCSV(ruta)

					case "5":
						opReportes := "0"
						exit3 := true
						for exit3 {
							fmt.Println("**********************************************")
							fmt.Println("*             Escoge la lista                *")
							fmt.Println("**********************************************")
							fmt.Println("*   1. Empleados                             *")
							fmt.Println("*   2. Imagenes                              *")
							fmt.Println("*   3. Usuarios                              *")
							fmt.Println("*   4. Cola                                  *")
							fmt.Println("*   5. Pila                                  *")
							fmt.Println("*   6. Regresar                              *")
							fmt.Println("**********************************************")
							fmt.Scanln(&opReportes)
							switch opReportes {
							case "1":
								listaEmpleados.Reporte()

							case "2":
								listaImagenes.Reporte()

							case "3":
								listaUsuarios.Reporte()

							case "4":
								cola.Reporte()

							case "5":
								pila.Reporte()

							case "6":
								exit3 = false
							}

						}
					case "6":
						exit2 = false
					}
				}
			} else if listaEmpleados.Loging(usuario, contraseña) {
				opAdmin := "0"
				exit2 := true
				for exit2 {
					fmt.Println("**********************************************")
					fmt.Println("*           DASHBOARD EMPLEADOS              *")
					fmt.Println("**********************************************")
					fmt.Println("*   1. Ver imagenes                      *")
					fmt.Println("*   2. Realizar pedido                       *")
					fmt.Println("*   3. Desloguearse                          *")
					fmt.Println("**********************************************")
					fmt.Scanln(&opAdmin)

					switch opAdmin {
					case "1":
						img := ""
						fmt.Println("**********************************************")
						fmt.Println("*                IMAGENES                    *")
						fmt.Println("**********************************************")
						aux := listaImagenes.Raiz
						for i := 0; i < listaImagenes.Longitud; i++ {
							number := i + 1
							fmt.Println(strconv.Itoa(number) + " " + aux.Nombre)
							aux = aux.Siguiente
						}
						fmt.Print("Ingrese la imagen que quieres: ")
						fmt.Scanln(&img)

					case "2":
						idCliente := ""
						idEmpleado := ""
						nombreImagen := ""
						fmt.Println("**********************************************")
						fmt.Println("*                 PEDIDOS                    *")
						fmt.Println("**********************************************")
						fmt.Print("Ingrese el id del cliente: ")
						fmt.Scanln(&idCliente)
						fmt.Print("Ingrese el id del empleado: ")
						fmt.Scanln(&idEmpleado)
						fmt.Print("Ingrese el nombre de la imagen: ")
						fmt.Scanln(&nombreImagen)
						pila.Push(idCliente, idEmpleado, nombreImagen)

					case "3":
						exit2 = false

					}
				}
			}

		case "2":
			fmt.Println("Programa finalizado.")
			exit1 = false
		}
	}

}
