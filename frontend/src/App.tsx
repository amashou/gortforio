import React, { useState, useEffect } from 'react';
import { Counter } from './features/counter/Counter';
import './App.css';
import CleanUp from './CleanUp';
import axios from 'axios'

const App: React.FC = () => {
  const [status, setStatus] = useState("first step");
  const [display, setDisplay] = useState(true);

  const getXemBalance = () => {
    // axios.defaults.baseURL = 'http://localhost:8080';
    // axios.defaults.headers.get['Content-Type'] = 'application/json';
    // axios.defaults.headers.get['Access-Control-Allow-Origin'] = '*';
    axios.get("http://localhost:8080/xem").then((res: any) => console.log(res.data)).catch((err) => {console.log("レスポンスに発生しました:", err)})
  }

  return (
    <div className="App">
      <h4>{status}</h4>
      <header className="App-header">
        {display &&  <CleanUp />}
        <button onClick={() => setDisplay(!display)}>Toggle Display</button>

        <button onClick={getXemBalance}>XENの情報を取得する</button>

        <Counter />
        <p>
          Edit <code>src/App.tsx</code> and save to reload.
        </p>
        <span>
          <span>Learn </span>
          <a
            className="App-link"
            href="https://reactjs.org/"
            target="_blank"
            rel="noopener noreferrer"
          >
            React
          </a>
          <span>, </span>
          <a
            className="App-link"
            href="https://redux.js.org/"
            target="_blank"
            rel="noopener noreferrer"
          >
            Redux
          </a>
          <span>, </span>
          <a
            className="App-link"
            href="https://redux-toolkit.js.org/"
            target="_blank"
            rel="noopener noreferrer"
          >
            Redux Toolkit
          </a>
          ,<span> and </span>
          <a
            className="App-link"
            href="https://react-redux.js.org/"
            target="_blank"
            rel="noopener noreferrer"
          >
            React Redux
          </a>
        </span>
      </header>
    </div>
  );
}

export default App;
