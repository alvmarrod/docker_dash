import React from 'react';
import Button from './Button';

import './Module.css';
import docker_logo from './Moby-logo.png';

class RunningContainers extends React.Component{

    render(){

        return (
            <div className="Module">
              <header className="Module-header">
                <span className="Module-title">
                  <img src={docker_logo} className="Module-logo" alt="logo" />
                  Running Containers
                </span>
                <br/>
                <div className="Table-Wrapper">
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
                          <th></th>
                      </tr>
                    </thead>
                    <tbody>
                      <tr className="Regular-Row">
                          <th>Col Data 1</th>
                          <th>Col Data 2</th>
                          <th>Col Data 3</th>
                          <th>Col Data 4</th>
                          <th>Col Data 5</th>
                          <th>Col Data 6</th>
                          <th>Col Data 7</th>
                          <th><Button title="Stop"/></th>
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
                      <tr className="Regular-Row">
                          <th>Col Data 1</th>
                          <th>Col Data 2</th>
                          <th>Col Data 3</th>
                          <th>Col Data 4</th>
                          <th>Col Data 5</th>
                          <th>Col Data 6</th>
                          <th>Col Data 7</th>
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
                      <tr className="Regular-Row">
                          <th>Col Data 1</th>
                          <th>Col Data 2</th>
                          <th>Col Data 3</th>
                          <th>Col Data 4</th>
                          <th>Col Data 5</th>
                          <th>Col Data 6</th>
                          <th>Col Data 7</th>
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
                      <tr className="Regular-Row">
                          <th>Col Data 1</th>
                          <th>Col Data 2</th>
                          <th>Col Data 3</th>
                          <th>Col Data 4</th>
                          <th>Col Data 5</th>
                          <th>Col Data 6</th>
                          <th>Col Data 7</th>
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
                      <tr className="Regular-Row">
                          <th>Col Data 1</th>
                          <th>Col Data 2</th>
                          <th>Col Data 3</th>
                          <th>Col Data 4</th>
                          <th>Col Data 5</th>
                          <th>Col Data 6</th>
                          <th>Col Data 7</th>
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
                    </tbody>
                  </table>
                </div>
              </header>
            </div>
        );

    }
}

export default RunningContainers;