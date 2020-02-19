import React from 'react';
import Switch from "react-switch";
import DockerContainer from './DockerContainer';

import './css/Module.css';
import docker_logo from './resources/Moby-logo.png';

class DockerContainers extends React.Component {

  constructor(props) {

    super(props);

    this.state = { 
      show_creation_date: true
    };

    this.handleChange = this.handleChange.bind(this);

  }

  handleChange(show_cdate) {
    this.setState({ 
      show_creation_date: show_cdate
    });
  }

  renderConditionalColumn(col_name) {

    switch (col_name) {
      case "Created":

        if (this.state.show_creation_date)
        {
          
          return <th>Created</th>
        }
        break;
      default:
        
    }

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

        return <DockerContainer key={this.props.title + "_DC_" + index} 
                                container={val}
                                button={buttonTemplate}
                                show_cdate={this.state.show_creation_date} />

      });
    }

    return DockerContainersInstances;

  }

  render() {

    const switchStyle = {
      width: 35,
      height: 20
    };

    return (
      <div className="Module">
        <header className="Module-header">
          <span className="Module-title">
            <img src={docker_logo} className="Module-logo" alt="logo" />
            {this.props.title} Containers
          </span>
          <br />
        </header>
        <div className="Module-body">
          <div className="Option">
            <span>
              Show creation date &nbsp;
              <Switch onChange={this.handleChange}
                      checked={this.state.show_creation_date}
                      width={switchStyle.width}
                      height={switchStyle.height}/>
            </span>
          </div>
          <div className="Table-Wrapper">
            <table>
              <thead>
                <tr className="Table-Header">
                  <th>Name</th>
                  <th>ID</th>
                  <th>Image</th>
                  <th>CMD</th>
                  {this.renderConditionalColumn("Created")}
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
        </div>
      </div>
    );

  }
}

export default DockerContainers;