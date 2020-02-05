import React from 'react';
import Button from './Button';

class DockerContainer extends React.Component{

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
                <th><Button title={this.props.button["title"]} className={this.props.button["className"]}/></th>
            </tr>
        )

    }
}

export default DockerContainer;