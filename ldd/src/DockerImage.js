import React from 'react';
import ButtonContainer from './containers/ButtonContainer';

class DockerImage extends React.Component{

    render(){

        return (
            <tr className="Regular-Row">
                
                <th>{this.props.image.repository}</th>
                <th>{this.props.image.tag}</th>
                <th>{this.props.image.id}</th>
                <th>{this.props.image.created}</th>
                <th>{this.props.image.size}</th>
                <th><ButtonContainer title="Run" 
                                     className="Run" />
                    <ButtonContainer title="X" 
                                     className="Stop" /></th>
            </tr>
        )

    }
}

export default DockerImage;