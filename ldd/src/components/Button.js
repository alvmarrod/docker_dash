import React from 'react';

export const Button = (props) => {

    return (
        <button className={props.className} 
                onClick={props.launchFunction}>
                {props.text}
        </button>
    );

}

export default Button;