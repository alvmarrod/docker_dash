import React from 'react';
import DockerContainer from './DockerContainer';

import './Module.css';
import docker_logo from './Moby-logo.png';

class StoppedContainers extends React.Component {

  renderContainers() {

    var DockerContainers = [];

    if (this.props.containers != null) {
      DockerContainers = this.props.containers.map((val, index) => {
        return <DockerContainer key={"StoppedDockerContainer_" + index} container={val} />
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
            Stopped Containers
                </span>
          <br />
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
              </tr>
            </thead>
            <tbody>
              {this.renderContainers()}
            </tbody>
          </table>
        </header>
      </div>
    );

  }
}

export default StoppedContainers;