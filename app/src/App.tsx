import React from 'react';
import { Route, Routes } from "react-router-dom";
import Menu from './screens/Menu';
import CreateUser from './screens/CreateUser';
import ListUser from './screens/ListUser';
import EditUser from './screens/EditUser';

import './App.scss'

function App() {
  return (
    <div className="App">
      <Routes>
        <Route path="/" element={<Menu/>} />
        <Route path="/create-user" element={<CreateUser/>} />
        <Route path="/list-user" element={<ListUser/>} />
        <Route path="/edit-user/:id" element={<EditUser/>} />
      </Routes>
    </div>
  );
}

export default App;
