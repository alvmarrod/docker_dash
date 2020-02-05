import React from 'react';

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
                <th>Appropriate button</th>
            </tr>
        )

    }
}

export default DockerContainer;