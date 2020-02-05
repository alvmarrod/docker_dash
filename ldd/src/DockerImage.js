import React from 'react';
import Button from './Button';

class DockerImage extends React.Component{

    render(){

        return (
            <tr className="Regular-Row">
                <th>{this.props.image.repository}</th>
                <th>{this.props.image.tag}</th>
                <th>{this.props.image.id}</th>
                <th>{this.props.image.created}</th>
                <th>{this.props.image.size}</th>
                <th><Button title="Run" className="Run" /></th>
            </tr>
        )

    }
}

export default DockerImage;