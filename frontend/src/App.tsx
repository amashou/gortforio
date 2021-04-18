import React, { useState, useEffect } from 'react';
import './App.css';
import CleanUp from './CleanUp';
import axios from 'axios'

import Header from './components/Header'
// import StockView from './StockView';

const App: React.FC = () => {
  const [status, setStatus] = useState("first step");
  const [display, setDisplay] = useState(true);
  const [stock, setStock] = useState({})

  const getXemBalance = () => {
    axios.get("http://localhost:8080/xem")
      .then((res: any) => {
        setStock(res.data)
      })
      .catch(
        (err: any) => {console.log("レスポンスに発生しました:", err)}
      )
  }

  return (
    <div className="App">
      <Header />
      <h4>{status}</h4>
      <header className="App-header">
        <button onClick={getXemBalance}>XENの情報を取得する</button>
      </header>
    </div>
  );
}

export default App;
