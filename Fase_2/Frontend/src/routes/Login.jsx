import { useState} from 'react';

import "bootstrap/dist/css/bootstrap.min.css";
import "./Login.css"

export default function Login() {
  const [user, setUser] = useState('');
  const [pass, setPass] = useState('');

  const handleLogin = () => {

    // Realiza la solicitud a la API aquí
    fetch('http://127.0.0.1:3001/login', {
      method: 'POST',
      body: JSON.stringify({ Username: user, Password: pass }),
      headers: {
        'Content-Type': 'application/json'
      }
    })
    .then(response => response.json())
    .then(data => validar(data))
  };

  const validar = (data) => {
    if(data.status == "400"){
        window.open("/admin","_self")
    }else if(data.status == "200"){
        localStorage.setItem("empleado", user)
        window.open("/empleados","_self")
    }else{
        console.log("Datos incorrectos")
    }
}

  return (
      <section className="vh-100 gradient-custom">
        <div className="container py-5 h-100">
          <div className="row d-flex justify-content-center align-items-center h-100">
            <div className="col-12 col-md-8 col-lg-6 col-xl-5">
              <div className="card bg-dark text-white" style={{ borderRadius: '1rem' }}>
                <form onSubmit={handleLogin} className="card-body p-5 text-center">
  
                <div className="mb-md-5 mt-md-4 pb-5">
                    <h2 className="fw-bold mb-2 text-uppercase">Inicio de sesión</h2>
                    <p className="text-white-50 mb-5"></p>
  
                    <div className="form-outline form-white mb-4">
                      <input type="text" id="typeEmailX" className="form-control form-control-lg" placeholder="Nombre Usuario" required onChange={e => setUser(e.target.value)} value={user} autoFocus/>
                      <label className="form-label" htmlFor="typeEmailX">Usuario</label>
                    </div>
  
                    <div className="form-outline form-white mb-4">
                      <input type="password" id="typePasswordX" className="form-control form-control-lg" placeholder="Password" aria-describedby="passwordHelpInline"  onChange={e => setPass(e.target.value)} value={pass}  autoFocus/>
                      <label className="form-label" htmlFor="typePasswordX">Contraseña</label>
                    </div>
  
                    <button className="btn btn-outline-light btn-lg px-5" type="submit">Iniciar sesión</button>
  
                  </div>
  
                </form>
              </div>
            </div>
          </div>
        </div>
      </section>
  );
}
