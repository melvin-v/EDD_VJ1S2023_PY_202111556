import { useState } from "react";

import "bootstrap/dist/css/bootstrap.min.css";
import "./Administrador.css";

export default function Administrador() {
  const [mensajes, setMensajes] = useState([]);
  const reportes = (e) => {
    e.preventDefault();
    console.log("Listo");
    window.open("/admin/reportes", "_self");
  };

  const onChange = (e) => {
    e.preventDefault();
    var reader = new FileReader();
    reader.onload = (e) => {
      var obj = JSON.parse(e.target.result);
      console.log(obj.pedidos);
      fetch("http://localhost:3001/cargarpedidos", {
        method: "POST",
        body: JSON.stringify({
          Pedidos: obj.pedidos,
        }),
        headers: {
          "Content-Type": "application/json",
        },
      })
        .then((response) => response.json())
        .then((data) => validar(data));
    };
    reader.readAsText(e.target.files[0]);
    setMensajes('Pedidos cargados exitosamente')
  };

  const onChange1 = (e, file1) => {
    var file = file1 || e.target.files[0];
    console.log(file.name);
    fetch("http://localhost:3001/cargarempleados", {
      method: "POST",
      body: JSON.stringify({
        Nombre: file.name,
      }),
      headers: {
        "Content-Type": "application/json",
      },
    })
      .then((response) => response.json())
      .then((data) => validar(data));
      setMensajes('Empleados cargados exitosamente')
  };

  const validar = (data) => {
    console.log(data);
  };

  const salir = (e) => {
    e.preventDefault();
    console.log("Listo");
    window.open("/", "_self");
  };
  return (
    <section className="vh-100 gradient-custom">
      <div className="container py-5 h-100">
        <div className="row d-flex justify-content-center align-items-center h-100">
          <div className="d-flex flex-column align-items-center">
          <div className="col-md-4 mb-5">
              <p>{mensajes}</p>
          </div>
            <div className="col-md-4 mb-5">
              <label className="input-group-text text-justify">
                Cargar Pedidos
              </label>
              <input
                className="form-control"
                id="inputGroupFile01"
                type="file"
                accept="application/json"
                onChange={onChange}
              />
            </div>
            <div className="col-md-4 mb-5">
              <label className="input-group-text">Cargar Empleados</label>
              <input
                className="form-control"
                id="inputGroupFile02"
                type="file"
                accept=".csv, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet, application/vnd.ms-excel"
                onChange={onChange1}
              />
            </div>
            <button
              type="button"
              className="btn btn-success btn-square-md mb-5"
              onClick={reportes}
            >
              Reportes
            </button>
            <button
              type="button"
              className="btn btn-danger btn-square-md-exit mb-3"
              onClick={salir}
            >
              Regresar
            </button>
          </div>
        </div>
      </div>
    </section>
  );
}
