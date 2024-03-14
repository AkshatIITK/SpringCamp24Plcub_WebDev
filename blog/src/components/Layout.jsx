import { Outlet } from "react-router-dom";
import "../App.css"
import Navbar from "./navbar";
const Layout = () => {
  return (
    <>
    <div className="main-container-wrapper">
      <div className="navbar-wrapper">
        <Navbar/>
      </div>
      
      <div className='box-content' >
        <Outlet />
      </div>
    </div>
      
    </>
  )
};

export default Layout;