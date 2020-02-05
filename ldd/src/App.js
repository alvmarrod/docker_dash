import './App.css';
import React from 'react';
import logo from './logo.svg';

import RunningContainers from './RunningContainers';
import StoppedContainers from './StoppedContainers';
import ExistingImages from './ExistingImages';
import Others from './Others';

class App extends React.Component{

  state = {
    images: [],
    runningcontainers: [],
    stoppedcontainers: []
  }
  
  componentDidMount() {
    // fetch('http://172.17.0.3:8000/images')
    // fetch('http://172.17.0.1:8000/images')
    fetch('http://10.20.30.54:8000/images')
    .then(res => res.json())
    .then((data) => {
      this.setState({ images: data })
    })
    .catch(console.log)

    fetch('http://10.20.30.54:8000/runningcontainers')
    .then(res => res.json())
    .then((data) => {
      this.setState({ runningcontainers: data })
    })
    .catch(console.log)

    fetch('http://10.20.30.54:8000/stoppedcontainers')
    .then(res => res.json())
    .then((data) => {
      this.setState({ stoppedcontainers: data })
    })
    .catch(console.log)

  }

  render () {
    return(
      <div className="App">
        <header className="App-header">
          <span className="Dashboard-title">
            <img src={logo} className="React-logo" alt="logo" />
            Local Docker Dashboard
          </span>
          <RunningContainers containers={this.state.runningcontainers}/>
          <br/>
          <StoppedContainers containers={this.state.stoppedcontainers}/>
          <br/>
          <ExistingImages images={this.state.images} />
          <br/>
          <Others/>
        </header>
      </div>
    );
  }

}

export default App;