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
      const response = await fetch(apiOrigin + '/usuarios', publicOptions('POST', data));
      const res = await response.json();
      return res
    },
    async edit(data, id) {
        const response = await fetch(apiOrigin + '/usuarios' +id, publicOptions('PUT', data));
        const res = await response.json();
        return res
      },
    async get(id) {
        const response = await fetch(apiOrigin + '/usuarios' +id, publicOptions('GET', id));
        const res = await response.json();
        return res
      },
    async getAll() {
        const response = await fetch(apiOrigin + '/usuarios', publicOptions('GET'));
        const res = await response.json();
        return res
      },
      async delete(id) {
        const response = await fetch(apiOrigin + '/usuarios' +id, publicOptions('DELETE'));
        const res = await response.json();
        return res
      }, 
  },
  generos: {
    async create(data) {
      const response = await fetch(apiOrigin + '/generos', publicOptions('POST', data));
      const res = await response.json();
      return res
    },
    async edit(data, id) {
        const response = await fetch(apiOrigin + '/generos' +id, publicOptions('PUT', data));
        const res = await response.json();
        return res
      },
    async get(id) {
        const response = await fetch(apiOrigin + '/generos' +id, publicOptions('GET', id));
        const res = await response.json();
        return res
      },
    async getAll() {
        const response = await fetch(apiOrigin + '/generos', publicOptions('GET'));
        const res = await response.json();
        return res
      },
      async delete(id) {
        const response = await fetch(apiOrigin + '/generos' +id, publicOptions('DELETE'));
        const res = await response.json();
        return res
      }, 
  },
  libros: {
    async create(data) {
      const response = await fetch(apiOrigin + '/libros', publicOptions('POST', data));
      const res = await response.json();
      return res
    },
    async edit(data, id) {
        const response = await fetch(apiOrigin + '/libros' +id, publicOptions('PUT', data));
        const res = await response.json();
        return res
      },
    async get(id) {
        const response = await fetch(apiOrigin + '/libros' +id, publicOptions('GET', id));
        const res = await response.json();
        return res
      },
    async getAll() {
        const response = await fetch(apiOrigin + '/libros', publicOptions('GET'));
        const res = await response.json();
        return res
      },
      async delete(id) {
        const response = await fetch(apiOrigin + '/libros' +id, publicOptions('DELETE'));
        const res = await response.json();
        return res
      }, 
  },
  prestamos: {
    async create(data) {
      const response = await fetch(apiOrigin + '/prestamos', publicOptions('POST', data));
      const res = await response.json();
      return res
    },
    async edit(data, id) {
        const response = await fetch(apiOrigin + '/prestamos' +id, publicOptions('PUT', data));
        const res = await response.json();
        return res
      },
    async get(id) {
        const response = await fetch(apiOrigin + '/prestamos' +id, publicOptions('GET', id));
        const res = await response.json();
        return res
      },
    async getAll() {
        const response = await fetch(apiOrigin + '/prestamos', publicOptions('GET'));
        const res = await response.json();
        return res
      },
      async delete(id) {
        const response = await fetch(apiOrigin + '/prestamos' +id, publicOptions('DELETE'));
        const res = await response.json();
        return res
      }, 
  }


};
