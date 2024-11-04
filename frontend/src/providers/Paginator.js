import React, { createContext, useEffect, useState } from 'react';
import { useLocation } from 'react-router-dom/cjs/react-router-dom.min';


export const PaginatorContext = createContext();

export const PaginatorProvider = ({ children }) => {
  const [query, setQuery] = useState("");             // Filtro por busqueda textual
  const [totalLength, setTotalLength] = useState(0);  // Total de registros
  const [currentPage, setCurrentPage] = useState(0);  // Pagina activa
  const [pageSize, setPageSize] = useState(10);       // Cantidad de registros por pagina
  const [categoriaSelected, setCategoriaSelected] = useState("");
  const [isToggled, setIsToggled] = useState(false);

  const location = useLocation(); // Obtener la ubicación actual

  // Reiniciar valores en función del cambio de ubicación
  useEffect(() => {
    // Restablecer valores al cambiar de vista
    setQuery("");
    setTotalLength(0);
    setCurrentPage(0);
    setPageSize(10);
    setCategoriaSelected("");
    setIsToggled(false);
  }, [location]); // Ejecutar efecto cuando cambie la ubicación

  return (
    <PaginatorContext.Provider
      value={
        {
          query,
          setQuery,
          totalLength,
          setTotalLength,
          currentPage,
          setCurrentPage,
          pageSize,
          setPageSize,
          categoriaSelected,
          setCategoriaSelected,
          isToggled,
          setIsToggled,
        }
      }
    >
      {children}
    </PaginatorContext.Provider>
  );
}