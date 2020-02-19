import React from 'react';
import DCRowItem from '../components/DCRowItem';

class DCRowItemContainer extends React.Component {

  render() {
    return <DCRowItem value={this.props.value} />
  }

}

export default DCRowItemContainer;