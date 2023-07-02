import {useState, useEffect} from 'react';
import './Facturas.css'
import 'bootstrap/dist/css/bootstrap.min.css'

export function Factura () {
    //const idEmpleado = localStorage.getItem("empleado")
    const [facturas, setFacturas] = useState([])
    const salir = (e) => {
        e.preventDefault();
        console.log("Listo")
        window.open("/empleados","_self")
    }

    useEffect(() => {
        peticion()
    },[])

    const peticion = () => {
        fetch('http://localhost:3001/facturaempleado',{
        })
        .then(response => response.json())
        .then(data => validar(data))
    }

    const validar = (data) =>{
        console.log(data.factura)
        setFacturas(data.factura) 
    }

    return(
        <section className="vh-100 gradient-custom">
                <div className="container py-5 h-100">
                <div className="row d-flex justify-content-center align-items-center h-100">
                    <div className="col-12 col-md-8 col-lg-6 col-xl-5">
                    <div className="card bg-dark text-white" style={{ borderRadius: '1rem' }}>
                    <div className="text-center">
                  <form>
                    <h1 className="h3 mb-3 fw-normal">Facturas Generadas <br/> Empleado {localStorage.getItem("empleado")}</h1>
                    <br/>
                    <table className="table table-dark table-striped">
                        <thead>
                            <tr>
                                <th scope="col">#</th>
                                <th scope="col">ID Cliente</th>
                                <th scope="col">ID Factura</th>
                            </tr>
                        </thead>
                        <tbody>
                            {
                                facturas.map((element, j) => {
                                    if (element.Id_Cliente != '') {
                                        return <>
                                        <tr key={"fact"+j}>
                                            <th scope="row">{j+1}</th>
                                            <td>{element.Id_Cliente}</td>
                                            <td>{element.Id_Factura}</td>
                                        </tr>
                                    </>
                                    }
                                })
                            }
                        </tbody>
                    </table>
                    <br/>
                    <center><button className="w-50 btn btn-outline-success" onClick={salir}>Salir</button></center>
                    <br/>
                  </form>
            </div>
                    </div>
                    </div>
                </div>
                </div>
            </section>


    );
}