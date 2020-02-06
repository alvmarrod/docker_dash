import './App.css';
import React from 'react';
import logo from './logo.svg';

import Others from './Others';
import DockerImages from './DockerImages';
import DockerContainers from './DockerContainers';

class App extends React.Component{

  intervalID;

  constructor(props){
    super(props);

    this.state = {
      images: [],
      runningcontainers: [],
      stoppedcontainers: []
    };

  }

  fetchAllData() {

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

  componentDidMount() {
    // First data fetch to populate the interface
    this.fetchAllData()
    
    /*
      Now we need to make it run at a specified interval,
      bind the getData() call to `this`, and keep a reference
      to the invterval so we can clear it later.
    */
    this.intervalID = setInterval(this.fetchAllData.bind(this), 1000*5);
  }

  componentWillUnmount() {
    // stop fetchAllData() from continuing to run even after unmounting this component
    clearInterval(this.intervalID);
  }

  render () {
    return(
      <div className="App">
        <header className="App-header">
          <span className="Dashboard-title">
            <img src={logo} className="React-logo" alt="logo" />
            Local Docker Dashboard
          </span>
          <DockerContainers title="Running" containers={this.state.runningcontainers}/>
          <br/>
          <DockerContainers title="Stopped" containers={this.state.stoppedcontainers}/>
          <br/>
          <DockerImages images={this.state.images} />
          <br/>
          <Others/>
        </header>
      </div>
    );
  }

}

export default App;