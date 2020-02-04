import './App.css';
import React from 'react';
import logo from './logo.svg';
import RunningContainers from './RunningContainers';
import StoppedContainers from './StoppedContainers';
import ExistingImages from './ExistingImages';
import Others from './Others';

function App(){

  return (
    <div className="App">
      <header className="App-header">
        <span className="Dashboard-title">
          <img src={logo} className="React-logo" alt="logo" />
          Local Docker Dashboard
        </span>
        <RunningContainers/>
        <br/>
        <StoppedContainers/>
        <br/>
        <ExistingImages/>
        <br/>
        <Others/>
      </header>
    </div>
  );

}

export default App;