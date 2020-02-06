import React from 'react';
import Switch from "react-switch";
import DockerImage from './DockerImage';

import './Module.css';
import docker_logo from './Moby-logo.png';

class DockerImages extends React.Component{

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
      
  renderImages(){

    var DockerImagesInstances = this.props.images.map( (val, index) => {
      return <DockerImage key={"DockerImage_" + index} image={val} />
    });

    return DockerImagesInstances;

  }

  render(){

    const switchStyle = {
      width: 35,
      height: 20
    };

    return (
      <div className="Module">
        <header className="Module-header">
          <span className="Module-title">
            <img src={docker_logo} className="Module-logo" alt="logo" />
            Existing Images
          </span>
          <br/>
        </header>
        <div className="Module-body">
          <div className="Option">
              <span>
                Show images without name &nbsp;
                <Switch onChange={this.handleChange}
                        checked={this.state.checked}
                        width={switchStyle.width}
                        height={switchStyle.height}/>
              </span>
            </div>
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
        </div>
      </div>
    );

  }
}

export default DockerImages;