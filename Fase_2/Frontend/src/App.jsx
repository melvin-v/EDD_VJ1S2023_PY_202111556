import { createBrowserRouter, RouterProvider } from 'react-router-dom'
import Login from './routes/Login'
import Administrador from './routes/Administrador'
import Reportes from './routes/Reportes'
import Empleados from './routes/Empleados'
import Filtros from './routes/Filtros'
import GenerarFactura from './routes/GenerarFactura'
import { Factura } from './routes/Facturas'



const router = createBrowserRouter([
    {
      path: '/',
      element: <Login/>,
      errorElement: <h1>Error</h1>,
    },
    {
      path: '/admin',
      element: <Administrador/>,
    },
    {
      path: '/admin/reportes',
      element: <Reportes/>,
    },
    {
      path: '/empleados/reportes/facturas',
      element: <Factura/>,
    },
    {
      path: '/empleados',
      element: <Empleados/>,
    },
    {
      path: '/empleados/filtros',
      element: <Filtros/>,
    },
    {
      path: '/empleados/generarfactura',
      element: <GenerarFactura/>,
    }
  ])

export default function App() {
    return (
        <RouterProvider router={router}/>
    )
  }