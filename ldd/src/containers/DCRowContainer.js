import React from 'react';
import DCRow from '../components/DCRow';

class DCRowContainer extends React.Component {

  renderConditionalColumn(col_name) {

    switch (col_name) {
      case "Created":

        if (this.props.show_cdate) {
          return <th>{this.props.container.created}</th>
        }
        break;
      default:

    }

  }

  apiQueryForButton() {

    var query = "";

    if (this.props.button["title"] === "Stop") {
      query = "StopContainer";
    } else if (this.props.button["title"] === "Start") {
      query = "StartContainer";
    } else {
      console.log("Error, unknown button title: " + this.props.button["title"]);
    }

    return query;
  }

  render() {

    return <DCRow container={this.props.container}
                  button={this.props.button}
                  apiQuery={this.apiQueryForButton()}
                  ccCreated={this.renderConditionalColumn("Created")} />

  }
}

export default DCRowContainer;