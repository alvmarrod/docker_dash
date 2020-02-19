import './css/App.css';
import React from 'react';
import logo from './resources/react_logo.svg';

import Others from './components/Others';
import DockerImages from './DockerImages';
import DockerContainers from './DockerContainers';

class App extends React.Component{

  constructor(props){

    super(props);

    this.intervalID = null;

    this.state = {
      images: [],
      runningcontainers: [],
      stoppedcontainers: []
    };

  }

  fetchAllData() {

    fetch('http://localhost:8000/images')
    .then(res => res.json())
    .then((data) => {
      this.setState({ images: data })
    })
    .catch(console.log)

    fetch('http://localhost:8000/runningcontainers')
    .then(res => res.json())
    .then((data) => {
      this.setState({ runningcontainers: data })
    })
    .catch(console.log)

    fetch('http://localhost:8000/stoppedcontainers')
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