import React from 'react';
import { useHistory } from 'react-router-dom';
import Button from '@mui/material/Button';
import ArrowBackOutlinedIcon from '@mui/icons-material/ArrowBackOutlined';

const VolverButton = () => {
  const history = useHistory();

  const handleClick = () => {
    history.goBack();
  };

  return (
    <Button variant="contained" startIcon={<ArrowBackOutlinedIcon />}
      style={{margin: '5px', backgroundColor:"#141A2C", border: "1px solid #141A2C"}} 
      onClick={handleClick}
    >
      {"Volver"}
    </Button>
  )
}
  
export default VolverButton;