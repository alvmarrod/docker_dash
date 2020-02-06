import React from 'react';
import Switch from "react-switch";
import DockerContainer from './DockerContainer';

import './Module.css';
import docker_logo from './Moby-logo.png';

class DockerContainers extends React.Component {

  constructor(props) {

    super(props);

    this.state = { 
      checked: true
    };

    this.handleChange = this.handleChange.bind(this);

  }

  handleChange(checked) {
    this.setState({ checked });
  }

  renderContainers() {

    var DockerContainersInstances = [];

    if (this.props.containers != null) {

      DockerContainersInstances = this.props.containers.map((val, index) => {

        const buttonTemplate = { className: "Stop", title: "Stop" };
        if (this.props.title === "Stopped") {
          buttonTemplate.className = "Run";
          buttonTemplate.title = "Start";
        }

        return <DockerContainer key={this.props.title + "_DC_" + index} container={val} button={buttonTemplate} />

      });
    }

    return DockerContainersInstances;

  }

  render() {

    return (
      <div className="Module">
        <header className="Module-header">
          <span className="Module-title">
            <img src={docker_logo} className="Module-logo" alt="logo" />
            {this.props.title} Containers
          </span>
          <br />
          <span>
            <span>Show containers without name</span>
            <Switch onChange={this.handleChange} checked={this.state.checked} />
          </span>
          <div className="Table-Wrapper">
            <table>
              <thead>
                <tr className="Table-header">
                  <th>Name</th>
                  <th>Container ID</th>
                  <th>Image</th>
                  <th>CMD</th>
                  <th>Created</th>
                  <th>Status</th>
                  <th>Ports</th>
                  <th></th>
                </tr>
              </thead>
              <tbody>
                {this.renderContainers()}
              </tbody>
            </table>
          </div>
        </header>
      </div>
    );

  }
}

export default DockerContainers;