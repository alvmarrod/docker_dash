import React from 'react';
import docker_logo from './Moby-logo.png';
import './Module.css';

class ExistingImages extends React.Component{

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
                    <tr className="Regular-Row">
                        <th>Col Data 1</th>
                        <th>Col Data 2</th>
                        <th>Col Data 3</th>
                        <th>Col Data 4</th>
                        <th>Col Data 5</th>
                    </tr>
                </table>
              </header>
            </div>
        );

    }
}

export default ExistingImages;