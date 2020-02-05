import React from 'react';
import DockerImage from './DockerImage';

import './Module.css';
import docker_logo from './Moby-logo.png';

class ExistingImages extends React.Component{
      
  renderImages(){

    var DockerImages = this.props.images.map( (val, index) => {
      return <DockerImage key={"DockerImage_" + index} image={val} />
    });

    return DockerImages;

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
          <table>
            <tr className="Table-header">
              <th>Repository</th>
              <th>TAG</th>
              <th>Image ID</th>
              <th>Created</th>
              <th>Size</th>
            </tr>
            {this.renderImages()}
          </table>
        </header>
      </div>
    );

  }
}

export default ExistingImages;