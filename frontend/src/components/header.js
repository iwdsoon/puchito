import { useLocation } from 'react-router-dom';
import { useState, useEffect } from "react";


export const Header = () => {
    const location = useLocation();
    const [titulo, setTitulo] = useState('')
    const [icono, setIcono] = useState('');

    useEffect(() => {
        const url = "/home";

        switch (location.pathname) {
            case url:
                setIcono("widgets")
                setTitulo("Menu")
                break;
            case url + "/usuarios":
                setIcono("person");
                setTitulo("Usuarios");
            break;
            case url + "/generos":
                setIcono("wc");
                setTitulo("Generos");
            break;
            case url + "/libros":
                setIcono("menu_book");
                setTitulo("Libros");
            break;
            case url + "/prestamos":
                setIcono("store");
                setTitulo("Prestamos");
            break;
        }
    },[location.pathname]);
    
    return (
        <header
            style={{
                padding: "20px 10px",
            }}
        >
            <div className="d-flex align-items-center">
                <div className="d-flex justify-content-between align-items-center" style={{flex: 1}}>
                    <div className="d-flex">
                        <div className="d-flex justify-content-center align-items-center mr-2 text-light" style={{width: "35px", height: "35px", borderRadius: "50%", backgroundColor: "#2d2e2c"}}>
                            <i className="material-symbols-outlined ">{icono}</i>
                        </div>
                        <div>
                           <h2 style={{margin: "0px", display: "flex", alignItems: "center", fontWeight: "400"}}>{titulo}</h2>
                        </div>
                        
                    </div>

                </div>
            </div>
        </header>

    )
}

export default Header;
