import { useState } from "react";
import "bootstrap/dist/css/bootstrap.min.css";
import "./Filtros.css";

export default function Filtros() {
    const [filtro, setFiltro] = useState(0)
    const salir = (e) => {
        e.preventDefault();
        console.log("Listo")
        window.open("/empleados","_self")
    }

    const validar = (data) =>{
        console.log(data)
    }

    const aplicarFiltros = async(e) => {
        e.preventDefault();
        fetch('http://localhost:3001/aplicarfiltro',{
            method: 'POST',
            body: JSON.stringify({
                Tipo: filtro,
                NombreImagen: ""
            }),
            headers:{
                'Content-Type': 'application/json'
            }
        })
        .then(response => response.json())
        .then(data => validar(data))
        
    }

    const handleChange = (e) => {
        var j = parseInt(e.target.value);
        setFiltro(j)
    }
    return (

        <section className="vh-100 gradient-custom">
        <div className="container py-5 h-100">
          <div className="row d-flex justify-content-center align-items-center h-100">
            <div className="col-12 col-md-8 col-lg-6 col-xl-5">
              <div className="card bg-dark text-white" style={{ borderRadius: '1rem' }}>
                
                 <form>
                 <center><h1>Empleado {localStorage.getItem("empleado")}</h1></center>
                    <br/>
                    <center><h4>Elige un Filtro</h4></center>
                    
                    <br/>
                    <center> <div className="col-md-4 mb-5">
                        <select className="form-control" aria-label=".form-select-lg example" onChange={handleChange}>
                            <option value={0}>Elegir....</option>
                            <option value={1}>Negativo</option>
                            <option value={2}>Escala de Grises</option>
                            <option value={3}>Espejo X</option>
                            <option value={4}>Espejo Y</option>
                            <option value={5}>Ambos Espejos</option>
                        </select>
                    </div></center>
                    <center><button className="w-50 btn btn-outline-primary" onClick={aplicarFiltros}>Generar Imagen con Filtro</button></center>
                    <br/>
                    <center><button className="w-50 btn btn-outline-success" onClick={salir}>Salir</button></center>
                    <br/>
                    <p className="mt-5 mb-3 text-muted">EDD 201700918</p>
                  </form>

              </div>
            </div>
          </div>
        </div>
      </section>


    );
  }