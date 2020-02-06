import React from 'react';

import './Module.css';
import docker_logo from './Moby-logo.png';

class Others extends React.Component{

    render(){

        return (
            <div className="Module">
              <header className="Module-header">
                <span className="Module-title">
                  <img src={docker_logo} className="Module-logo" alt="logo" />
                  Other capabilities
                </span>
                <br/>
                <p>Here will be more options!</p>
              </header>
            </div>
        );

    }
}

export default Others;