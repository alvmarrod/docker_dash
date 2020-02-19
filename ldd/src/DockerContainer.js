import React from 'react';
import ButtonContainer from './containers/ButtonContainer';

class DockerContainer extends React.Component {

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

    return (
      <tr className="Regular-Row">
        <th>{this.props.container.name}</th>
        <th>{this.props.container.id}</th>
        <th>{this.props.container.image}</th>
        <th>{this.props.container.cmd}</th>
        {this.renderConditionalColumn("Created")}
        <th>{this.props.container.status}</th>
        <th>{this.props.container.ports}</th>
        <th><ButtonContainer title={this.props.button["title"]}
                             className={this.props.button["className"]}
                             apiQuery={this.apiQueryForButton()}
                             itemID={this.props.container.id} /></th>
      </tr>
    )

  }
}

export default DockerContainer;