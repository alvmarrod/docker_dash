import React from 'react';

import './Button.css';

class Button extends React.Component{

    constructor(props){

        super(props);

        this.state = {
            container_id: null
        };

        this.launchFunction = this.launchFunction.bind(this);

    }

    dictionaryToString(myDict){

        var result = "";

        result = Object.keys(myDict).map(function(key, index){
            if (typeof(myDict[key]) != 'object'){
                return "\t" + key + ": " + myDict[key] + "\n"
            } else
            {
                return this.dictionaryToString(myDict[key])
            }
            
        }, this)

        return result

    }

    launchFunction(){

        var apiQueryURL = "";
        var apiQuerySettings = {};

        var body = {};
        body['id'] = this.props.itemID;

        switch (this.props.apiQuery) {
            case "StartContainer":
                apiQueryURL = "http://10.20.30.54:8000//containers/" + this.props.itemID;
                apiQuerySettings['method'] = 'PUT';
                apiQuerySettings['headers'] = {
                    'Content-Type': 'application/json'
                };
                body['action'] = 'Start';
                apiQuerySettings['body'] = JSON.stringify(body);
                break;
            case "StopContainer":
                apiQueryURL = "http://10.20.30.54:8000//containers/" + this.props.itemID;
                apiQuerySettings['method'] = 'PUT';
                apiQuerySettings['headers'] = {
                    'Content-Type': 'application/json'
                };
                body['action'] = 'Stop';
                apiQuerySettings['body'] = JSON.stringify(body);
                break;
            case "RunImage":
                apiQueryURL = "http://10.20.30.54:8000/runningcontainers";
                break;
            default:
                // code
        }

        // const newState = this.state.hover ? false : true;
        // this.setState({ hover: newState });
        // alert("Data: " + this.state.container_id)

        // alert("Data: " + apiQueryURL + "\n\n" + "Settings:\n" + this.dictionaryToString(apiQuerySettings))

        fetch(apiQueryURL, apiQuerySettings)
        .then(res => res.json())
        .then((data) => {
        // this.setState({ runningcontainers: data })
            console.log(data);
        })
        .catch(console.log)

    }

    render(){

        return (
            <button className={this.props.className} onClick={this.launchFunction}>
                {this.props.title.toUpperCase()}
            </button>
        );

    }
}

export default Button;