import React from 'react';
import Button from './Button';
import DockerContainer from './DockerContainer';

import './Module.css';
import docker_logo from './Moby-logo.png';

class RunningContainers extends React.Component {

  renderContainers() {

    var DockerContainers = [];

    if (this.props.containers != null) {
      DockerContainers = this.props.containers.map((val, index) => {
        const buttonTemplate = {className: "Stop", title: "Stop"};
        return <DockerContainer key={"RunningDockerContainer_" + index} container={val} button={buttonTemplate} />
      });
    }

    return DockerContainers;

  }

  render() {

    return (
      <div className="Module">
        <header className="Module-header">
          <span className="Module-title">
            <img src={docker_logo} className="Module-logo" alt="logo" />
            Running Containers
          </span>
          <br />
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

export default RunningContainers;