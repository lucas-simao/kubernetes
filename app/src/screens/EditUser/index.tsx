import React, {useState, useEffect} from 'react'
import FormUser from '../../components/FormUser';
import { useParams } from 'react-router-dom';
import { getCustomersById, patchCustomer } from '../../api';
import { TCustomer } from '../../helpers/types';

import './index.scss'

const EditUser = () => {
  const [user, setUser] = useState<TCustomer>()
  const {id} = useParams()
  
  const handleGetUser = async (id: string) => {
    await getCustomersById(id).then(({data}) => {
      setUser(data)
    })
  }

  const handleUpdateUser = (u: TCustomer) => {
    if (!id) return
    patchCustomer(id, u).then(() => {
      alert("updated success")
    }).catch(e => alert(e))
  }

  useEffect(() => {
    if (id) handleGetUser(id)
  }, [id])

  return (
    <div className='editUser'>
      <h1>Edit user</h1>
      <FormUser onSave={(u) => handleUpdateUser(u)} user={user}></FormUser>
    </div>
  )
}

export default EditUser;