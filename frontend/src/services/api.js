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
    async createUsuario(data) {
      const response = await fetch(apiOrigin + '/usuarios', pivateOptions('POST', data));
      const res = await response.json();
      return res
    },
  }

};
