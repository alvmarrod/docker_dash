import React from 'react';

import ButtonContainer from '../containers/ButtonContainer';

export const DCRow = (props) => {

    return (
        <tr className="Regular-Row">
        <th>{props.container.name}</th>
        <th>{props.container.id}</th>
        <th>{props.container.image}</th>
        <th>{props.container.cmd}</th>
        {props.ccCreated}
        <th>{props.container.status}</th>
        <th>{props.container.ports}</th>
        <th><ButtonContainer title={props.button["title"]}
                             className={props.button["className"]}
                             apiQuery={props.apiQuery}
                             itemID={props.container.id} /></th>
        </tr>
    );

}

export default DCRow;