import React from 'react';
import Button from './Button';

class DockerContainer extends React.Component{

    apiQueryForButton(){

        var query = "";

        if (this.props.button["title"] == "Stop") {
            query = "StopContainer";            
        } else if (this.props.button["title"] == "Start") {
            query = "StartContainer";
        } else {
            console.log("Error, unknown button title: " + this.props.button["title"]);
        }

        return query;
    }

    render(){

        return (
            <tr className="Regular-Row">
                <th>{this.props.container.name}</th>
                <th>{this.props.container.id}</th>
                <th>{this.props.container.image}</th>
                <th>{this.props.container.cmd}</th>
                <th>{this.props.container.created}</th>
                <th>{this.props.container.status}</th>
                <th>{this.props.container.ports}</th>
                <th><Button title={this.props.button["title"]}
                            className={this.props.button["className"]}
                            apiQuery={this.apiQueryForButton()} /></th>
            </tr>
        )

    }
}

export default DockerContainer;