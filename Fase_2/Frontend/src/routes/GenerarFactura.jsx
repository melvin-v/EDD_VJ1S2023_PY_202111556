import {useState, useEffect} from 'react';
import './GenerarFactura.css'
import 'bootstrap/dist/css/bootstrap.min.css'

import moment from 'moment'

// 18:6:5
export default function GenerarFactura () {
    //(DD-MM-YY-::HH:MM:SS)
    const fecha = moment().format("DD-MM-yyyy-::hh:mm:ss")
    const idEmpleado = localStorage.getItem("empleado")
    const [cliente, setCliente] = useState('')
    const [pago, setPago] = useState('')

    useEffect(() => {
        peticion();
    },[])

    const peticion = () => {
        fetch("http://localhost:3001/obtenerPedido",{})
        .then(response => response.json())
        .then(data => validar(data))
    }

    const validar = (data) =>{
        var idcliente = data.datos.Id_Cliente
        setCliente(idcliente.toString())
    }

    const handleSubmit = (e) => {
        e.preventDefault();
        fetch('http://localhost:3001/generarfactura',{
            method: 'POST',
            body: JSON.stringify({
                Timestamp: fecha,
                Biller:    idEmpleado,
                Customer:  cliente,
                Payment:   pago
            }),
            headers:{
                'Content-Type': 'application/json'
            }
        })
        .then(response => response.json())
        .then(data => console.log(data))
        alert('Factura generada exitosamente')
    }

    const salir = (e) => {
        e.preventDefault();
        console.log("Listo")
        window.open("/empleados","_self")
    }

    return(
        <section className="vh-100 gradient-custom">
        <div className="container py-5 h-100">
          <div className="row d-flex justify-content-center align-items-center h-100">
            <div className="col-12 col-md-8 col-lg-6 col-xl-5">
              <div className="card bg-dark text-white" style={{ borderRadius: '1rem' }}>
              <div className="text-center">
                  <form onSubmit={handleSubmit}>
                  <h1 className="h3 mb-3 fw-normal">Dashboard Empleado {localStorage.getItem("empleado")}</h1>
                    <label htmlFor="inputEmail" className="visually-hidden">Fecha</label>
                    <input type="text" id="userI" className="form-control" placeholder="Nombre Usuario" disabled
                    value={fecha}  
                    autoFocus/>
                    <br/>
                    <label htmlFor="inputEmail" className="visually-hidden">Empleado Cobrador</label>
                    <input type="text" id="userI" className="form-control" placeholder="Nombre Usuario" disabled
                    value={idEmpleado}  
                    autoFocus/>
                    <br/>
                    <label htmlFor="inputEmail" className="visually-hidden">Usuario</label>
                    <input type="text" id="userI" className="form-control" placeholder="Nombre Usuario" disabled
                    value={cliente}  
                    autoFocus/>
                    <br/>
                    <label htmlFor="inputEmail" className="visually-hidden">Pago</label>
                    <input type="text" id="userI" className="form-control" placeholder="Q0.0" required
                    onChange={e => setPago(e.target.value)} 
                    value={pago}  
                    autoFocus/>
                    <br/>
                    <center><button className="w-50 btn btn-outline-success" type="submit">Generar Pago</button></center>
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
    )
}