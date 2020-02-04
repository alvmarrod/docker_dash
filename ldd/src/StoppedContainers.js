import React from 'react';
import docker_logo from './Moby-logo.png';
import './Module.css';

class StoppedContainers extends React.Component{

    render(){

        return (
            <div className="Module">
              <header className="Module-header">
                <span className="Module-title">
                  <img src={docker_logo} className="Module-logo" alt="logo" />
                  Stopped Containers
                </span>
                <br/>
                <table>
                    <tr className="Table-header">
                        <th>Name</th>
                        <th>Container ID</th>
                        <th>Image</th>
                        <th>CMD</th>
                        <th>Created</th>
                        <th>Status</th>
                        <th>Ports</th>
                    </tr>
                    <tr className="Regular-Row">
                        <th>Col Data 1</th>
                        <th>Col Data 2</th>
                        <th>Col Data 3</th>
                        <th>Col Data 4</th>
                        <th>Col Data 5</th>
                        <th>Col Data 6</th>
                        <th>Col Data 7</th>
                    </tr>
                </table>
              </header>
            </div>
        );

    }
}

export default StoppedContainers;