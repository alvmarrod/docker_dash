import React from 'react';

import docker_logo from '../resources/Moby-logo.png';

export class Others extends React.Component {

  render() {

    return (
      <div className="Module">
        <header className="Module-header">
          <span className="Module-title">
            <img src={docker_logo} className="Module-logo" alt="logo" />
            Other capabilities
                </span>
          <br />
          <p>Here will be more options!</p>
        </header>
      </div>
    );

  }
}

export default Others;