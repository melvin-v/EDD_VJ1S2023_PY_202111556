import { useState } from "react";

import "bootstrap/dist/css/bootstrap.min.css";
import "./Reportes.css";

export default function Reportes  (){
    const [imagen, setImagen] = useState('https://images.squarespace-cdn.com/content/v1/5e10bdc20efb8f0d169f85f9/1590751925678-5XVSVXMC2BX38RNKKO19/music.png?format=750w')
    const salir = (e) => {
        e.preventDefault();
        console.log("Listo")
        window.open("/admin","_self")
    }

    const validar = (data) =>{
        console.log(data)
        setImagen(data.imagen.Imagenbase64)
    }

    const reporteGrafo = async(e) => {
        e.preventDefault();
        fetch('http://localhost:3001/reporte-grafo',{
        })
        .then(response => response.json())
        .then(data => validar(data));
    }

    const reporteArbol = async(e) => {
        e.preventDefault();
        fetch('http://localhost:3001/reporte-arbol',{
        })
        .then(response => response.json())
        .then(data => validar(data));
    }

    const reporteBlockchain = async(e) => {
        e.preventDefault();
        fetch('http://localhost:3001/reporte-bloque',{
        })
        .then(response => response.json())
        .then(data => validar(data));
    }

    return(

        <section className="vh-100 gradient-custom">
        <div className="container py-5 h-100">
          <div className="row d-flex justify-content-center align-items-center h-100">
            <div className="col-12 col-md-8 col-lg-6 col-xl-5">
              <div className="card bg-dark text-white" style={{ borderRadius: '1rem' }}>
              <div className="text-center">
                  <form>
                    <br />
                    <h1 className="h3 mb-3 fw-normal">Reportes empleados</h1>
                    <br/>
                    <center><button className="w-50 btn btn-outline-primary" onClick={reporteGrafo}>Grafo</button></center>
                    <br/>
                    <center><button className="w-50 btn btn-outline-primary" onClick={reporteArbol}>Arbol AVL</button></center>
                    <br/>
                    <center><button className="w-50 btn btn-outline-primary" onClick={reporteBlockchain}>Facturas</button></center>
                    <br/>
                    <center><button className="w-50 btn btn-outline-success" onClick={salir}>Salir</button></center>
                    <br/>
                    <center><img src={imagen} width="350" height="350" alt='some value' /></center>
                  </form>
            </div>
              </div>
            </div>
          </div>
        </div>
      </section>
    );
}