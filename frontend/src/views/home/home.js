import { Container } from "reactstrap";
import { useHistory } from "react-router-dom/cjs/react-router-dom.min";

export const Home = () => {

    const history = useHistory();


  return (
    <>
        <Container fluid="sm">
            <div
              style={{ display: "flex", alignItems: "center", justifyContent: "center", gap: "25px", marginBottom: "25px", flexWrap: "wrap"}}
            >
              <div
                onClick={() => history.push("/home/usuarios")}
                className="shadow"
                style={{
                  width: "230px",
                  height: "180px",
                  display: "flex",
                  flexDirection: "column",
                  alignItems: "center",
                  justifyContent: "center",
                  cursor:"pointer"
                }}
              >
                <div
                  style={{
                    width: "60px",
                    height: "60px",
                    borderRadius: "50%",
                    marginBottom: "10px",
                    backgroundColor: "#e4f3fc",
                    display: "flex",
                    alignItems: "center",
                    justifyContent: "center",
                  }}
                >
                  <i
                    className="material-symbols-outlined"
                    style={{ color: "#067bea", fontSize: "30px" }}
                  >
                    person
                  </i>
                </div>
                <h4 className="m-0">{"Usuarios"}</h4>
              </div>

              <div
                onClick={() => history.push("/home/generos")}
                className="shadow"
                style={{
                  width: "230px",
                  height: "180px",
                  display: "flex",
                  flexDirection: "column",
                  alignItems: "center",
                  justifyContent: "center",
                  cursor:"pointer"
                }}
              >
                <div
                  style={{
                    width: "60px",
                    height: "60px",
                    borderRadius: "50%",
                    marginBottom: "10px",
                    backgroundColor: "#f5f2ff",
                    display: "flex",
                    alignItems: "center",
                    justifyContent: "center",
                  }}
                >
                  <i
                    className="material-symbols-outlined"
                    style={{ color: "#8763f9", fontSize: "30px" }}
                  >
                    wc
                  </i>
                </div>
                <h4 className="m-0">{"Generos"}</h4>
              </div>

              <div
              onClick={() => history.push("/home/libros")}
                className="shadow"
                style={{
                  width: "230px",
                  height: "180px",
                  display: "flex",
                  flexDirection: "column",
                  alignItems: "center",
                  justifyContent: "center",
                  cursor:"pointer"
                }}
              >
                <div
                  style={{
                    width: "60px",
                    height: "60px",
                    borderRadius: "50%",
                    marginBottom: "10px",
                    backgroundColor: "#f5f2ff",
                    display: "flex",
                    alignItems: "center",
                    justifyContent: "center",
                  }}
                >
                  <i
                    className="material-symbols-outlined"
                    style={{ color: "#8763f9", fontSize: "30px" }}
                  >
                    menu_book
                  </i>
                </div>
                <h4 className="m-0">{"Libros"}</h4>
              </div>

              <div
                onClick={() => history.push("/home/prestamos")}
                className="shadow"
                style={{
                  width: "230px",
                  height: "180px",
                  display: "flex",
                  flexDirection: "column",
                  alignItems: "center",
                  justifyContent: "center",
                  cursor:"pointer"
                  
                }}
              >
                <div
                  style={{
                    width: "60px",
                    height: "60px",
                    borderRadius: "50%",
                    marginBottom: "10px",
                    backgroundColor: "#f5f2ff",
                    display: "flex",
                    alignItems: "center",
                    justifyContent: "center",
                  }}
                >
                  <i
                    className="material-symbols-outlined"
                    style={{ color: "#8763f9", fontSize: "30px" }}
                  >
                    store
                  </i>
                </div>
                <h4 className="m-0">{"Prestamos"}</h4>
              </div>

            </div>

        </Container>
    </>
  );
};
export default Home;