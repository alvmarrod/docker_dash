import React from 'react';
import Switch from "react-switch";

import DCRowContainer from '../containers/DCRowContainer';

import docker_logo from '../resources/Moby-logo.png';

export const DCTable = (props) => {

  const switchStyle = {
    width: 35,
    height: 20
  };

  let DockerContainersInstances = [];

  if (props.containers != null) {

    DockerContainersInstances = props.containers.map((val, index) => {

      const buttonTemplate = { className: "Stop", title: "Stop" };
      if (props.title === "Stopped") {
        buttonTemplate.className = "Run";
        buttonTemplate.title = "Start";
      }

      // console.log("Container: " + val);
      
      return <DCRowContainer key={props.title + "_DC_" + index}
                             container={val}
                             button={buttonTemplate}
                             show_cdate={props.show_cd} />

    });
  }

  return (
    <div className="Module">
      <header className="Module-header">
        <span className="Module-title">
          <img src={docker_logo} className="Module-logo" alt="logo" />
          {props.title} Containers
          </span>
        <br />
      </header>
      <div className="Module-body">
        <div className="Option">
          <span>
            Show creation date &nbsp;
              <Switch onChange={props.handleChange}
                      checked={props.show_cd}
                      width={switchStyle.width}
                      height={switchStyle.height} />
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
                {props.show_cd ? <th>Created</th> : ""}
                <th>Status</th>
                <th>Ports</th>
                <th></th>
              </tr>
            </thead>
            <tbody>
              {DockerContainersInstances}
            </tbody>
          </table>
        </div>
      </div>
    </div>
  );

}

export default DCTable;