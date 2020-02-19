import React from 'react';

import DCTable from '../components/DCTable';

class DCTableContainer extends React.Component {

  constructor(props) {

    super(props);

    this.state = {
      show_creation_date: true
    };

    this.handleChange = this.handleChange.bind(this);

  }

  handleChange(show_cdate) {
    this.setState({
      show_creation_date: show_cdate
    });
  }

  render() {

    return (
      <DCTable title={this.props.title}
               handleChange={this.handleChange}
               show_cd={this.state.show_creation_date}
               containers={this.props.containers}
               />
    );

  }
}

export default DCTableContainer;