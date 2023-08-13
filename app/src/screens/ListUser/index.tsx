import React, {useEffect, useState} from 'react'
import { TCustomer } from '../../helpers/types';
import { deleteCustomer, getCustomers } from '../../api';
import dayjs from 'dayjs';
import { useNavigate } from 'react-router';

import './index.scss'

const ListUser = () => {
  const [users, setUsers] = useState<TCustomer[]>([])

  const navigate = useNavigate();

  const handleGetUsers = async () => {
    await getCustomers().then(({data}) => {
      if (data) {
        setUsers(data)
      } else {
        setUsers([])
      }
    })
  }

  useEffect(() => {
   handleGetUsers()
  }, [])

  const formatDate = (d: Date | undefined) => {
    if (d) return dayjs(d).format('DD/MM/YYYY HH:mm')
  }

  const handleDelete = async (id: string | undefined) => {
    if (!id) return
    await deleteCustomer(id)
    await handleGetUsers()
  }

  const handleEdit = async (id: string | undefined) => {
    if (!id) return
    navigate(`/edit-user/${id}`)
  }

  return (
    <div className='listUser'>
      <h1>List of users</h1>
      <table>
        <tr>
          <th>ID</th>
          <th>First name</th>
          <th>Last name</th>
          <th>Birthdate</th>
          <th>Updated at</th>
          <th>Created at</th>
          <th>Actions</th>
        </tr>
          {users.map((v, index) => {
            return (
              <tr key={index}>
                <td>{v.id}</td>
                <td>{v.firstName}</td>
                <td>{v.lastName}</td>
                <td>{v.birthDate}</td>
                <td>{formatDate(v.updatedAt)}</td>
                <td>{formatDate(v.createdAt)}</td>
                <td>
                  <button onClick={() => handleEdit(v.id)} className='listUser--edit'>Edit</button>
                  <button onClick={() => handleDelete(v.id)} className='listUser--delete'>Delete</button>
                </td>
              </tr>
            )
          })}
      </table>
    </div>
  )
}

export default ListUser;