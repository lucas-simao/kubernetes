import React, {FC, useState, useEffect} from 'react'
import { TCustomer } from '../../helpers/types';
import Input from '../Input'
import './index.scss'

const emptyCustomer: TCustomer = {
  firstName: '',
  lastName: '',
  birthDate: '',
}

type Props = {
  user?: TCustomer;
  onSave: (user: TCustomer) => void
}

const FormUser:FC<Props> = ({user, onSave}:Props) => {
  const [customer, setCustomer] = useState<TCustomer>(emptyCustomer);

  useEffect(() => {
    if (user) setCustomer(user)
  }, [user])
  
  const handleUser = (key: string, value: string) => {
    setCustomer({
      ...customer, 
      [key]: value
    })
  }

  return (
    <form onSubmit={(e) => {
      e.preventDefault()
      onSave(customer)
    }} className='formUser'>
      <Input id='firstName' type='text' label='First Name' value={customer.firstName} onChange={(v) => handleUser('firstName', v)} />
      <Input id='lastName' type='text' label='Last Name' value={customer.lastName} onChange={(v) => handleUser('lastName', v)} />
      <Input id='birthDate' type='date' label='BirthDate' value={customer.birthDate} onChange={(v) => handleUser('birthDate', v)} />
      <button type='submit'>Save</button>
    </form>
  )
}

export default FormUser;