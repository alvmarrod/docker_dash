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
        // const newState = this.state.hover ? false : true;
        // this.setState({ hover: newState });
        alert("Data: " + this.state.container_id)
    }

    render(){

        return (
            <button onClick={this.launchFunction}>
                {this.props.title.toUpperCase()}
            </button>
        );

    }
}

export default Button;