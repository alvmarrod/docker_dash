import React from 'react';

import './Button.css';

class Button extends React.Component{

    constructor(props){

        super(props);

        this.state = {
            container_id: null
        };

        this.launchFunction = this.launchFunction.bind(this);

    }


    launchFunction(){

        var apiQuery = "";

        switch (this.props.apiQuery) {
            case "StartContainer":
                apiQuery = "http://10.20.30.54:8000/runningcontainers";
                break;
            case "StopContainer":
                apiQuery = "http://10.20.30.54:8000/runningcontainers";
                break;
            case "RunImage":
                apiQuery = "http://10.20.30.54:8000/runningcontainers";
                break;
            default:
                // code
        }

        // const newState = this.state.hover ? false : true;
        // this.setState({ hover: newState });
        // alert("Data: " + this.state.container_id)

        alert("Data: " + apiQuery)

        // fetch('http://10.20.30.54:8000/runningcontainers')
        // .then(res => res.json())
        // .then((data) => {
        // this.setState({ runningcontainers: data })
        // })
        // .catch(console.log)

    }

    render(){

        return (
            <button className={this.props.className} onClick={this.launchFunction}>
                {this.props.title.toUpperCase()}
            </button>
        );

    }
}

export default Button;