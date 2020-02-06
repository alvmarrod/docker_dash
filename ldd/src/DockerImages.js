import React from 'react';
import DockerImage from './DockerImage';

import './Module.css';
import docker_logo from './Moby-logo.png';

class DockerImages extends React.Component{
      
  renderImages(){

    var DockerImagesInstances = this.props.images.map( (val, index) => {
      return <DockerImage key={"DockerImage_" + index} image={val} />
    });

    return DockerImagesInstances;

  }

  render(){

    return (
      <div className="Module">
        <header className="Module-header">
          <span className="Module-title">
            <img src={docker_logo} className="Module-logo" alt="logo" />
            Existing Images
          </span>
          <br/>
          <div className="Table-Wrapper">
            <table>
              <thead>
                <tr className="Table-header">
                  <th>Repository</th>
                  <th>TAG</th>
                  <th>Image ID</th>
                  <th>Created</th>
                  <th>Size</th>
                  <th></th>
                </tr>
              </thead>
              <tbody>
                {this.renderImages()}
              </tbody>
            </table>
          </div>
        </header>
      </div>
    );

  }
}

export default DockerImages;