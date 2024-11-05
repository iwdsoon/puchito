import React, { useEffect, useState } from 'react';
import { useHistory } from 'react-router-dom';
import moment from 'moment-timezone';
import {
  Container,
  TextField,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
  Paper,
  IconButton,
} from '@mui/material';
import DeleteForeverIcon from '@mui/icons-material/DeleteForever';
import SearchOutlinedIcon from '@mui/icons-material/SearchOutlined';
import CloseOutlinedIcon from '@mui/icons-material/CloseOutlined';
import VolverButton from '../components/VolverButton';
import { api } from '../services/api';
import CrearGenero from '../components/modales/CrearGenero';
import EditarGenero from '../components/modales/EditarGenero';


const Generos = () => {
  const [state, setState] = useState({ showResult: false, apiMessage: "", error: null });
  const history = useHistory();
  const [generos, setGeneros] = useState([]);
  const [pageSize] = useState(10);
  const [showSearch, setShowSearch] = useState(false);
  const [query, setQuery] = useState("");

  const formatFecha = (fecha) => {
      return moment(fecha).clone().local().format("DD/MM/YY")
  }

  useEffect(() => {
    const callApi = async () => {
      try {
        const response = await api.generos.getAll();
        if (response.status === "success") {
          setGeneros(response.data.generos || []);
        } else {
          setGeneros([]);
        }
      } catch (error) {
        setState(prevState => ({
          ...prevState,
          error: "Error en la Red.",
        }));
      }
    };

    callApi();
  }, []);


  const handleDelete = async (id) => {
    const confirmDelete = window.confirm(
      "¿Estás seguro de que deseas eliminar esta publicación?"
    );
    if (!confirmDelete) return;

    try {
      const response = await api.generos.delete(id);
      if (response.status === "success") {
        setGeneros(generos.filter((g) => g.id !== id));
      }
    } catch (error) {
      setState(prevState => ({
        ...prevState,
        error: "Error al eliminar el genero.",
      }));
    }
  };

  const handleToggleSearch = () => {
    setShowSearch(!showSearch);
    if (showSearch) setQuery('');
  };

  return (
    <Container>
      <div style={{ display: "flex", justifyContent: "space-between", alignItems: "center", marginBottom: "1rem" }}>

        {showSearch ? (
          <div style={{ display: "flex", alignItems: "center" }}>
            <TextField
              value={query}
              onChange={(e) => setQuery(e.target.value)}
              placeholder="Buscar.."
              variant="outlined"
              size="small"
            />
            <IconButton onClick={handleToggleSearch}>
              <CloseOutlinedIcon />
            </IconButton>
          </div>
        ) : (
          <IconButton onClick={handleToggleSearch}>
            <SearchOutlinedIcon />
          </IconButton>
        )}
            <CrearGenero
              generos={generos}
            />
      </div>
        <>
          <TableContainer component={Paper}>
            <Table size="small">
              <TableHead sx={{ backgroundColor: "#EDEDED" }}>
                <TableRow>
                  <TableCell sx={{ fontFamily: 'Inter, sans-serif' }}>{"Id"}</TableCell>
                  <TableCell sx={{ fontFamily: 'Inter, sans-serif' }}>{"Genero"}</TableCell>
                  <TableCell sx={{ fontFamily: 'Inter, sans-serif' }} align="right">
                    {"Acciones"}
                  </TableCell>
                </TableRow>
              </TableHead>
              <TableBody>
                  {generos.map((genero) => (
                    <TableRow key={genero.id}>
                      <TableCell sx={{ fontFamily: 'Inter, sans-serif' }}>{genero.id}</TableCell>
                      <TableCell sx={{ fontFamily: 'Inter, sans-serif' }}>{genero.genero}</TableCell>
{/*                    <TableCell sx={{ fontFamily: 'Inter, sans-serif' }}>{formatFecha(pub.desde)}</TableCell>
                      <TableCell sx={{ fontFamily: 'Inter, sans-serif' }}>{formatFecha(pub.hasta)}</TableCell> */}
                      <TableCell align="right">
                        <IconButton
                          sx={{ color: "red" }}
                          onClick={() => handleDelete(genero.id)}
                        >
                          <DeleteForeverIcon sx={{ fontSize: "27px" }} />
                        </IconButton>
                        <EditarGenero
                            generoId={genero.id}
                            initialGenero={genero.genero}
                        />
                      </TableCell>
                    </TableRow>
                  ))}
              </TableBody>
            </Table>
          </TableContainer>
        </>
      <div style={{ display: "flex", justifyContent: "flex-end", marginTop: "1rem" }}>
        <VolverButton />
      </div>
    </Container>
  );
};

export default Generos;