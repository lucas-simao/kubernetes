import React from 'react'
import LeftIcon from '../../icons/chevron-left-icon.svg'
import { useLocation, useNavigate } from 'react-router'
import './index.scss'

const Header = () => {
  const location = useLocation();
  const navigate = useNavigate();

  return (
    <header className='header'>
      {location.pathname !== "/" && (
        <img onClick={() => navigate('')} src={LeftIcon} alt="back icon" />
      )}
      <h1>Simple app</h1>
    </header>
  )
}

export default Header;