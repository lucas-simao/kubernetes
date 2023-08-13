import React, { FC } from 'react'
import './index.scss'

type Props = {
  id: string;
  type: string;
  label: string;
  value: string | number;
  onChange: (text: string) => void;
}

const Input: FC<Props> = ({ id, type, label, value, onChange }: Props) => {
  const handleChange = (v: string) => {
    onChange(v)
  }

  return (
    <div className='input'>
      <label htmlFor={id}>{label}</label>
      <input id={id} value={value} onChange={({ target }) => handleChange(target.value)} type={type} />
    </div>
  )
}

export default Input;