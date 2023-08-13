import React, {useState} from 'react'
import FormUser from '../../components/FormUser';
import { postCustomer } from '../../api';
import { TCustomer } from '../../helpers/types';

import './index.scss'

const emptyCustomer: TCustomer = {
  firstName: '',
  lastName: '',
  birthDate: '',
}

const CreateUser = () => {
  const [customer, setCustomer] = useState<TCustomer>(emptyCustomer);

  const handleCreateUser = (u: TCustomer) => {
    postCustomer(u).then(() => {
      setCustomer(emptyCustomer)
      alert("created success")
    }).catch(e => alert(e))
  }

  return (
    <div className='createUser'>
      <h1>Create user</h1>
      <FormUser onSave={(u) => handleCreateUser(u)}></FormUser>
    </div>
  )
}

export default CreateUser;