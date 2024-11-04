import React, { useState, useEffect } from "react";
import { api } from "../../services/api";
import { useParams } from "react-router-dom";

export const Home = () => {
  const [state, setState] = useState({
    showResult: false,
    apiMessage: "",
    error: null,
  });
  const [evento, setEvento] = useState();
  const [eventosumary, setEventoSumary] = useState();
  const [tarifas, setTarifas] = useState([]);
  const [tickets, setTickets] = useState([]);
  const [vouchers, setVouchers] = useState([]);
  const param = useParams();


  const [dropdownOpen, setDropdownOpen] = useState(false);

  const toggle = () => setDropdownOpen((prevState) => !prevState);

  
  /*---------------------------------------------------------------------*/
  return (
    <>
        <div>
          {data.map((item, index) => (
            <div
              key={index}
              style={{ display: "flex", alignItems: "center", justifyContent: "center", gap: "25px", marginBottom: "25px", flexWrap: "wrap"}}
            >
              <div
                className="shadow"
                style={{
                  width: "230px",
                  height: "180px",
                  display: "flex",
                  flexDirection: "column",
                  alignItems: "center",
                  justifyContent: "center",
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
                    confirmation_number
                  </i>
                </div>
                <h2 className="m-0">{item.Cantidad}</h2>
                <p className="text-secondary m-0">Total {item.Tipo}</p>
              </div>

              <div
                className="shadow"
                style={{
                  width: "230px",
                  height: "180px",
                  display: "flex",
                  flexDirection: "column",
                  alignItems: "center",
                  justifyContent: "center",
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
                    qr_code_scanner
                  </i>
                </div>
                <h2 className="m-0">{item.Escaneados}</h2>
                <p className="text-secondary m-0">{t("dashboard.scanned")}</p>
              </div>

              <div
                className="shadow"
                style={{
                  width: "230px",
                  height: "180px",
                  display: "flex",
                  flexDirection: "column",
                  alignItems: "center",
                  justifyContent: "center",
                }}
              >
                <div
                  style={{
                    width: "60px",
                    height: "60px",
                    borderRadius: "50%",
                    marginBottom: "10px",
                    backgroundColor: "#fffbeb",
                    display: "flex",
                    alignItems: "center",
                    justifyContent: "center",
                  }}
                >
                  <i
                    className="material-symbols-outlined"
                    style={{ color: "#f7d43f", fontSize: "30px" }}
                  >
                    pending_actions
                  </i>
                </div>
                <h2 className="m-0">{item.Pendientes}</h2>
                <p className="text-secondary m-0">{t("dashboard.pending")}</p>
              </div>
            </div>
          ))}

          <div style={{ display: "flex", gap: "25px", marginBottom: "25px", flexWrap: "wrap", justifyContent: "center"}}>
            <div
              className="shadow"
              style={{
                width: "270px",
                height: "fit-content",
                display: "flex",
                flexDirection: "column",
                alignItems: "center",
                justifyContent: "center",
                padding: "20px"
              }}
            >
              <h2 className="mb-2">Tickets</h2>
              <div style={{width: "200px", height: "200px"}}>
                <Doughnut data={ticketsPorConcepto} options={options}/>
              </div>
              <div style={{width: '100%', marginTop: '10px'}}>
                {!noHayTickets ? (
                  <div>
                    <div className="d-flex justify-content-between">
                      <div>
                        <p><b>{t("dashboard.type")}</b></p>
                      </div>
                      <div>
                        <p><b>{t("dashboard.quantity")}</b></p>
                      </div>
                    </div>
                    {Object.entries(conteoTickets).map(([concepto, { cantidad }]) => (
                      <div key={concepto} className="w-100 d-flex mb-2">
                        <div style={{width: '50%'}}>
                          <p className="my-0 ml-2">{concepto}</p>
                        </div>
                        <div style={{width: '50%'}}>
                          <p className="text-end my-0 mr-2">{cantidad}</p>
                        </div>
                      </div>
                    ))}
                  </div>
                ) : (
                  <div className="w-100 d-flex justify-content-center">
                    <p className="m-0">{t("dashboard.noTicketsYet")}</p>
                  </div>
                )}
              </div>
            </div>

            <div
              className="shadow"
              style={{
                width: "270px",
                height: "fit-content",
                display: "flex",
                flexDirection: "column",
                alignItems: "center",
                padding: "20px"
              }}
            >
              <h2 className="mb-2">Vouchers</h2>
              <div style={{width: "200px", height: "200px"}}>
                <Doughnut data={vouchersPorConcepto} options={options}/>
              </div>
              <div style={{width: '100%', marginTop: '10px'}}>
                {!noHayVouchers ? (
                  <div>
                    <div className="d-flex justify-content-between">
                      <div>
                        <p><b>{t("dashboard.type")}</b></p>
                      </div>
                      <div>
                        <p><b>{t("dashboard.quantity")}</b></p>
                      </div>
                    </div>
                    {Object.entries(conteoVouchers).map(([concepto, { cantidad }]) => (
                      <div key={concepto} className="w-100 d-flex mb-2">
                        <div style={{width: '50%'}}>
                          <p className="my-0 ml-2">{concepto}</p>
                        </div>
                        <div style={{width: '50%'}}>
                          <p className="text-end my-0 mr-2">{cantidad}</p>
                        </div>
                      </div>
                    ))}
                  </div>
                ) : (
                  <div className="w-100 d-flex justify-content-center">
                    <p className="m-0">{t("dashboard.noVouchersYet")}</p>
                  </div>
                )}   
              </div>
            </div>
          </div>
        </div>
    </>
  );
};
export default Home;