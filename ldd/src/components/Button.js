import React from 'react';

import '../css/Button.css';

export const Button = (props) => {

    return (
        <button className={props.className} 
                onClick={props.launchFunction}>
                {props.text}
        </button>
    );

}

export default Button;