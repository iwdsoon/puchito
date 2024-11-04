import React, { useContext } from "react";
import { MenuTemplate } from "../../components/Menu/MenuTemplate";
import { useParams } from "react-router-dom";
import { useState, useEffect } from "react";
import { api } from "../../services/api";
import { Header } from "../../components/Header";
import { PaginatorContext } from "../../providers/Paginator";
import Loading from "../../components/Loading";
import { useMediaQuery } from "react-responsive";


export const Layout = ({ children }) => {
  const { setEvent } = useContext(PaginatorContext)
  const isMobile = useMediaQuery({ query: `(max-width: 770px)` });
  const param = useParams();
  const [state, setState] = useState({
    showResult: false,
    apiMessage: "",
    error: null
  });
  const [evento, setEvento] = useState();
  const [loading, setLoading] = useState(true);


  useEffect(() => {
    const callApi = async () => {
      try {
        const response = await api.eventos.get(param.id);
        if (response.status === "success") {
          const data = response.data.evento;
          setEvento(data);
          setEvent(data.evento)
          return data;
        }
      } catch (error) {
        setState({
          ...state,
          error: "Error en la Red.",
        });
      } finally {
        setLoading(false);
      }
    };
    callApi();
  }, []);

  if (loading) {
    return <Loading />;
  }

  return (
    <>
      <div className="body d-flex" style={{ height: "100%" }}>
        {!isMobile && (
          <MenuTemplate id='menu-template' evento={evento} />
        )}
        <div style={{ flex: "1 1 0%" }}>
          {evento &&
            <div className="h-100 d-flex flex-column">
              <Header evento={evento}></Header>
              <div className="flex-grow-1">
                {children}
              </div>
            </div>
          }
        </div>
      </div>

      <Footer />
    </>
  )
}

export default Layout;