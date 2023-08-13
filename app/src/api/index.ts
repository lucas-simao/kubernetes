import axios from 'axios'
import { TCustomer } from '../helpers/types';

const {REACT_APP_API_URL} = process.env;

const api = REACT_APP_API_URL as string;

export const postCustomer = (data: TCustomer) => {
  const url = `${api}/customers`

  return axios.post(url, data)
}

export const getCustomers = () => {
  const url = `${api}/customers`

  return axios.get(url)
}

export const getCustomersById = (id: string) => {
  const url = `${api}/customers/${id}`

  return axios.get(url)
}

export const patchCustomer = (id:string, data: TCustomer) => {
  const url = `${api}/customers/${id}`

  return axios.patch(url, data)
}

export const deleteCustomer = (id: string) => {
  const url = `${api}/customers/${id}`
  
  return axios.delete(url)
}
