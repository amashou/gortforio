import React from 'react';
import './App.css';
import { Switch, Route } from 'react-router-dom';

import Header from './components/Header'
import Coincheck from './pages/Coincheck';
import Gmo from './pages/Gmo';
import Binance from './pages/Binance';
import Zaif from './pages/Zaif';
import Top from './pages/Top';
import { Container, makeStyles } from '@material-ui/core';

const useStyle = makeStyles((theme) => ({
  container: {
    marginTop: '120px',
  }
}));

const App: React.FC = () => {
  const classes = useStyle();

  return (
    <div className="App">
      <Header />
      <Container className={classes.container}>
        <Switch>
          <Route exact path="/">
              <Top />
            </Route>
            <Route path="/coincheck">
              <Coincheck />
            </Route>
            <Route path="/gmocoin">
              <Gmo />
            </Route>
            <Route path="/binance">
              <Binance />
            </Route>
            <Route path="/zaif">
              <Zaif />
            </Route>
          </Switch>
      </Container>
    </div>
  );
}

export default App;
