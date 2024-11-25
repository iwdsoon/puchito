import React, { useState } from 'react';
import { api } from '../../services/api'; 
import { grey } from '@mui/material/colors';
import { Button, Box, Modal, TextField, Alert} from '@mui/material';
import CheckIcon from '@mui/icons-material/Check';
import AddOutlinedIcon from "@mui/icons-material/AddOutlined";

const CrearGenero = ({handleCreate}) => {
  const [genero, setGenero] = useState('');
  const [error, setError] = useState(false);
  const [errorMessage , setErrorMessage] = useState("");
  const [enviado, setEnviado] = useState(false);

  const handleSubmit = async (e) => {
    e.preventDefault();
    if (genero.trim() === '') {
      setError(true);
      setErrorMessage(<>{"Campo vacio"}</>)
    }
      try {

        let data = {
          genero: genero
        }
        
        const response = await api.generos.create(data);
        if (response.status === "success") {
          setError(false)
          setErrorMessage("")
          setEnviado(true)
          handleCreate(response.data)
          setTimeout(() => {
            handleClose();
          }, 2000);
        }
      } catch (error) {
          console.log(error);
          setGenero("")
          setError(false)
          setErrorMessage("")
      }
    
  };

  const [open, setOpen] = React.useState(false);
  const handleOpen = () => setOpen(true);
  const handleClose = () => {
    setOpen(false); 
    setGenero("")
    setError(false)
    setErrorMessage("")
    setEnviado(false)
  } 

  const style = {
    position: 'absolute',
    top: '50%',
    left: '50%',
    transform: 'translate(-50%, -50%)',
    bgcolor: 'background.paper',
    boxShadow: 24,
    p: 3,
  };  

  return (
    <>
      <Button
          variant="contained"
          sx={{
            backgroundColor: "#FFC300",
            color: "black",
            outline: "none",
            borderRadius: "5px",
            textTransform: "none",
            fontFamily: 'Inter, sans-serif',
            '&:hover': {
              backgroundColor: "#FFB000",
            }
          }}
          startIcon={<AddOutlinedIcon sx={{ marginLeft: "5px" }} />}
          onClick={handleOpen}
        >
          {"Agregar Genero"}
        </Button>

      <Modal
        open={open}
        onClose={handleClose}
        aria-labelledby="modal-modal-title"
        aria-describedby="modal-modal-description"
      >
        <Box sx={style} onSubmit={handleSubmit}>
          <h4 id="modal-modal-title">
            {"Agregar Genero"}
          </h4>
          <div id="modal-modal-description" sx={{ mt: 2 }}>
            <Box
              component="form"
              noValidate
              autoComplete="off"
            >
              <TextField 
                label={"Genero"}
                id="location"
                name="location"
                value={genero}
                onChange={(e) => setGenero(e.target.value)}
                variant="outlined"
                fullWidth
                error={error} 
                helperText= {error && errorMessage}
              />
              {enviado && (
                <Alert icon={<CheckIcon fontSize="inherit"/>} style={{marginTop: "10px"}} severity="success">
                  Genero Creado.
                </Alert>
              )}
              <div style={{ display: 'flex', justifyContent: 'flex-end', marginTop: "15px" }}>
                <Button size="small" variant="contained" sx={{background: grey[500]}} style={{marginRight: "10px"}} onClick={handleClose}>
                  {"Cancelar"}
                </Button>
                <Button variant="contained" type="submit" size="small">
                  {"Crear"}
                </Button>
              </div>
            </Box>
          </div>
        </Box>
      </Modal>
    </>
  );
};

export default CrearGenero;