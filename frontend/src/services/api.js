import { config } from "./config";
import 'isomorphic-fetch';

const apiOrigin = config.URL_API;
const publicOrigin = config.URL_PUBLIC;


const publicOptions = (method, data) => {
  let opt = {
    method,
    headers: {
      'Content-Type': 'application/json',
    }
  };
  if (data) {
    opt.body = JSON.stringify(data);
  }
  return opt;
};
const pivateOptions = (method, data) => {
  let opt = {
    method,
    headers: {
      'Content-Type': 'application/json',
      'AUTHORIZATION': `Bearer ${sessionStorage.getItem('token')}`
    }
  };
  if (data) {
    opt.body = JSON.stringify(data);
  }

  return opt;
};

export const api = {

  usuarios: {
    async create(data) {
      const response = await fetch(publicOrigin + '/usuarios', publicOptions('POST', data));
      const res = await response.json();
      return res
    },
    async update(id, data) {
        const response = await fetch(publicOrigin + '/usuarios/' +id, publicOptions('PUT', data));
        const res = await response.json();
        return res
      },
    async get(id) {
        const response = await fetch(publicOrigin + '/usuarios/' +id, publicOptions('GET'));
        const res = await response.json();
        return res
      },
    async getAll() {
        const response = await fetch(publicOrigin + '/usuarios', publicOptions('GET'));
        const res = await response.json();
        return res
      },
      async delete(id) {
        const response = await fetch(publicOrigin + '/usuarios/' +id, publicOptions('DELETE'));
        const res = await response.json();
        return res
      }, 
  },
  generos: {
    async create(data) {
      const response = await fetch(publicOrigin + '/generos', publicOptions('POST', data));
      const res = await response.json();
      return res
    },
    async update(id, data) {
        const response = await fetch(publicOrigin + '/generos/' +id, publicOptions('PUT', data));
        const res = await response.json();
        return res
      },
    async get(id) {
        const response = await fetch(publicOrigin + '/generos/' +id, publicOptions('GET'));
        const res = await response.json();
        return res
      },
    async getAll() {
        const response = await fetch(publicOrigin + '/generos', publicOptions('GET'));
        const res = await response.json();
        return res
      },
      async delete(id) {
        const response = await fetch(publicOrigin + '/generos/' +id, publicOptions('DELETE'));
        const res = await response.json();
        return res
      }, 
  },
  libros: {
    async create(data) {
      const response = await fetch(publicOrigin + '/libros', publicOptions('POST', data));
      const res = await response.json();
      return res
    },
    async update(id, data) {
        const response = await fetch(publicOrigin + '/libros/' +id, publicOptions('PUT', data));
        const res = await response.json();
        return res
      },
    async get(id) {
        const response = await fetch(publicOrigin + '/libros/' +id, publicOptions('GET'));
        const res = await response.json();
        return res
      },
    async getAll() {
        const response = await fetch(publicOrigin + '/libros', publicOptions('GET'));
        const res = await response.json();
        return res
      },
      async delete(id) {
        const response = await fetch(publicOrigin + '/libros/' +id, publicOptions('DELETE'));
        const res = await response.json();
        return res
      }, 
  },
  prestamos: {
    async create(data) {
      const response = await fetch(publicOrigin + '/prestamos', publicOptions('POST', data));
      const res = await response.json();
      return res
    },
    async update(id, data) {
        const response = await fetch(publicOrigin + '/prestamos/' +id, publicOptions('PUT', data));
        const res = await response.json();
        return res
      },
    async get(id) {
        const response = await fetch(publicOrigin + '/prestamos/' +id, publicOptions('GET'));
        const res = await response.json();
        return res
      },
    async getAll() {
        const response = await fetch(publicOrigin + '/prestamos', publicOptions('GET'));
        const res = await response.json();
        return res
      },
      async delete(id) {
        const response = await fetch(publicOrigin + '/prestamos/' +id, publicOptions('DELETE'));
        const res = await response.json();
        return res
      }, 
  }


};
