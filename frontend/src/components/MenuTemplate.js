import { useContext, useEffect, useState } from "react";
import { Link, useLocation} from "react-router-dom";
import { PaginatorContext } from "../providers/Paginator";
import "./styles.css";
import HomeOutlinedIcon from '@mui/icons-material/HomeOutlined';
import StoreIcon from '@mui/icons-material/Store';
import PersonIcon from '@mui/icons-material/Person';
import WcIcon from '@mui/icons-material/Wc';
import MenuBookIcon from '@mui/icons-material/MenuBook';


export const MenuTemplate = ({ usuario }) => {

  const location = useLocation();
  const imagen = "https://res-console.cloudinary.com/dzuhe3wkh/thumbnails/v1/image/upload/v1730745674/YWRtaW5fb25vcG1o/drilldown" 


  return (
    <>
      <div
        className="menu"
        style={{
          width: "270px",
          padding: "15px",
        }}
      >
        <div className="menu-header d-flex align-items-center gap-2">
          <div
            className="menu-header-img"
            style={{ width: "50px", height: "50px", flexShrink: "0" }}
          >
            <img
              src={imagen}
              alt="Usuario"
              className="rounded-circle"
              style={{
                width: "100%",
                height: "100%",
                objectFit: "cover",
                objectPosition: "center",
              }}
            ></img>
          </div>
          <div className="menu-header-info_evento">
            <p className="m-0">{usuario?.nombre}</p>
          </div>
        </div>

        <div className="menu-body" style={{ marginTop: "25px" }}>
          <ul>
            
              <li
                className={`menu-body-item ${
                  location.pathname === "/home" ? "active" : ""
                }`}
              >
                <Link className="link" to={"/home"}>
                  <HomeOutlinedIcon/>
                  <p style={{ fontSize: "16px" }}>{"Inicio"}</p>
                </Link>
              </li>

              <li
                className={`menu-body-item ${
                  location.pathname === "/home/usuarios"
                    ? "active"
                    : ""
                }`}
              >
                <Link className="link" to={`/home/usuarios`}>
                  <PersonIcon/>
                  <p>{"Usuarios"}</p>
                </Link>
              </li>

              <li
                className={`menu-body-item ${
                  location.pathname === "/home/generos"
                    ? "active"
                    : ""
                }`}
              >
                <Link className="link" to={`/home/generos`}>
                  <WcIcon/>
                  <p>{"Generos"}</p>
                </Link>
              </li>

              <li
                className={`menu-body-item ${
                  location.pathname === "/home/libros"
                    ? "active"
                    : ""
                }`}
              >
                <Link className="link" to={`/home/libros`}>
                  <MenuBookIcon/>
                  <p>{"Libros"}</p>
                </Link>
              </li>

              <li
                className={`menu-body-item ${
                  location.pathname === "/home/prestamos"
                    ? "active"
                    : ""
                }`}
              >
                <Link className="link" to={`/home/prestamos`}>
                  <StoreIcon/>
                  <p>{"Prestamos"}</p>
                </Link>
              </li>

          </ul>
        </div>
      </div>
    </>
  );
};