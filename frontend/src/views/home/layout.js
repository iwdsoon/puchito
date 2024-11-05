import { MenuTemplate } from "../../components/MenuTemplate";
import { useState, useEffect } from "react";
import { api } from "../../services/api";
import { Header } from "../../components/Header";

export const Layout = ({ children }) => {

  const [usuario, setUsuario] = useState();
  const [state, setState] = useState({ showResult: false, apiMessage: "", error: null });


  useEffect(() => {
    const callApi = async () => {
    try {
      const response = await api.usuarios.get(1);
      if (response.status === "success") {
        const data = response.data.usuario;
        setUsuario(data);
      }
    } catch (error) {
      setState(prevState => ({
        ...prevState,
        error: "Error en la Red.",
      }));
    }
  }
  callApi();
}, []);

  return (
    <>
      <div className="body d-flex" style={{ height: "100%" }}>
        <MenuTemplate id='menu-template' usuario={usuario} />
        <div style={{ flex: "1 1 0%" }}>
          
            <div className="h-100 d-flex flex-column">
              <Header></Header>
              <div className="flex-grow-1">
                {children}
              </div>
            </div>
          
        </div>
      </div>
    </>
  )
}

export default Layout;