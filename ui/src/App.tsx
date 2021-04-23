import React, { useState } from 'react'
import './App.css'
import "bootstrap/dist/css/bootstrap.css"; // Import precompiled Bootstrap css
import Table from './components/table';


function App() {
  return (
    <div className="App">
      <div className="app-layout">
        <div className="header box"><h1>Tunnel🤦‍♂️Man</h1></div>
        <div className="table box">
          <Table />
        </div>
        <div className="input box">
          <input type="text" placeholder=" 🔎 Filter ..."/>
        </div>
      </div>
    </div>
  )
}

export default App
