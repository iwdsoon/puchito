import React, { useState } from 'react';
import { Table, Button, Modal, Form, Dropdown, DropdownButton } from 'react-bootstrap';

const Generos = () => {
  const [generos, setGeneros] = useState([]);
  const [showAddModal, setShowAddModal] = useState(false);
  const [showEditModal, setShowEditModal] = useState(false);
  const [showDeleteModal, setShowDeleteModal] = useState(false);
  const [selectedGenero, setSelectedGenero] = useState(null);
  const [newGenero, setNewGenero] = useState('');
  
  const handleAddGenero = () => {
    const newId = generos.length ? generos[generos.length - 1].id + 1 : 1;
    setGeneros([...generos, { id: newId, genero: newGenero }]);
    setNewGenero('');
    setShowAddModal(false);
  };

  const handleEditGenero = () => {
    setGeneros(generos.map(g => g.id === selectedGenero.id ? { ...g, genero: newGenero } : g));
    setShowEditModal(false);
  };

  const handleDeleteGenero = () => {
    setGeneros(generos.filter(g => g.id !== selectedGenero.id));
    setShowDeleteModal(false);
  };

  return (
    <div className="container mt-4">
      <h2>Generos</h2>
      <Button variant="primary" onClick={() => setShowAddModal(true)}>Add Genero</Button>

      <Table striped bordered hover className="mt-3">
        <thead>
          <tr>
            <th>ID</th>
            <th>Genero</th>
            <th>Action</th>
          </tr>
        </thead>
        <tbody>
          {generos.map(g => (
            <tr key={g.id}>
              <td>{g.id}</td>
              <td>{g.genero}</td>
              <td>
                <DropdownButton id="dropdown-basic-button" title="...">
                  <Dropdown.Item onClick={() => { setSelectedGenero(g); setNewGenero(g.genero); setShowEditModal(true); }}>Edit</Dropdown.Item>
                  <Dropdown.Item onClick={() => { setSelectedGenero(g); setShowDeleteModal(true); }}>Delete</Dropdown.Item>
                </DropdownButton>
              </td>
            </tr>
          ))}
        </tbody>
      </Table>

      {/* Add Modal */}
      <Modal show={showAddModal} onHide={() => setShowAddModal(false)}>
        <Modal.Header closeButton>
          <Modal.Title>Add Genero</Modal.Title>
        </Modal.Header>
        <Modal.Body>
          <Form.Group controlId="formAddGenero">
            <Form.Label>Genero</Form.Label>
            <Form.Control
              type="text"
              placeholder="Enter genero"
              value={newGenero}
              onChange={(e) => setNewGenero(e.target.value)}
            />
          </Form.Group>
        </Modal.Body>
        <Modal.Footer>
          <Button variant="secondary" onClick={() => setShowAddModal(false)}>Cancel</Button>
          <Button variant="primary" onClick={handleAddGenero}>Add</Button>
        </Modal.Footer>
      </Modal>

      {/* Edit Modal */}
      <Modal show={showEditModal} onHide={() => setShowEditModal(false)}>
        <Modal.Header closeButton>
          <Modal.Title>Edit Genero</Modal.Title>
        </Modal.Header>
        <Modal.Body>
          <Form.Group controlId="formEditGenero">
            <Form.Label>Genero</Form.Label>
            <Form.Control
              type="text"
              placeholder="Edit genero"
              value={newGenero}
              onChange={(e) => setNewGenero(e.target.value)}
            />
          </Form.Group>
        </Modal.Body>
        <Modal.Footer>
          <Button variant="secondary" onClick={() => setShowEditModal(false)}>Cancel</Button>
          <Button variant="primary" onClick={handleEditGenero}>Save Changes</Button>
        </Modal.Footer>
      </Modal>

      {/* Delete Modal */}
      <Modal show={showDeleteModal} onHide={() => setShowDeleteModal(false)}>
        <Modal.Header closeButton>
          <Modal.Title>Confirm Delete</Modal.Title>
        </Modal.Header>
        <Modal.Body>Are you sure you want to delete this genero?</Modal.Body>
        <Modal.Footer>
          <Button variant="secondary" onClick={() => setShowDeleteModal(false)}>Cancel</Button>
          <Button variant="danger" onClick={handleDeleteGenero}>Delete</Button>
        </Modal.Footer>
      </Modal>
    </div>
  );
};

export default Generos;