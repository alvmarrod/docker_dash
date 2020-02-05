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
    containers: []
  }
  
  componentDidMount() {
    fetch('http://172.17.0.3:8000/images')
    .then(res => res.json())
    .then((data) => {
      this.setState({ images: data })
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
          <RunningContainers/>
          <br/>
          <StoppedContainers/>
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