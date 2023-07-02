import "bootstrap/dist/css/bootstrap.min.css";
import "./Empleados.css";

export default function Empleados() {
  const filtro = (e) => {
    e.preventDefault();
    console.log("Listo");
    window.open("/empleados/filtros", "_self");
  };

    const reportes = (e) => {
      e.preventDefault();
      console.log("Listo");
      window.open("/empleados/generarfactura", "_self");
    };

    const ver = (e) => {
      e.preventDefault();
      console.log("Listo");
      window.open("/empleados/reportes/facturas", "_self");
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
              <button
                type="button"
                className="btn btn-success btn-square-md mb-3"
                onClick={filtro}
              >
                Aplicacion de filtros
              </button>
              <button
                type="button"
                className="btn btn-primary btn-square-md mb-3"
                onClick={reportes}
              >
                Generar factura
              </button>
              <button
                type="button"
                className="btn btn-warning btn-square-md mb-5"
                onClick={ver}
              >
                Ver facturas
              </button>
              <button
                type="button"
                className="btn btn-danger btn-square-md-exit mb-3"
                onClick={salir}
              >
                Cerrar sesión
              </button>
            </div>
          </div>
        </div>
      </section>
    );
  }