import React from 'react';
import { useNavigate } from 'react-router';
import './index.scss'

const menu = [
  {
    title: 'Create User',
    path: '/create-user'
  },
  {
    title: 'List User',
    path: '/list-user'
  }
]

function App () {
  const navigate = useNavigate();

  return (
    <div className='menu'>
      {menu.map((v, index) => {
        return (
          <div onClick={() => navigate(v.path)} key={index} className='menu--item'>
            <p>{v.title}</p>
          </div>
        )
      })}
    </div>
  );
}

export default App;
