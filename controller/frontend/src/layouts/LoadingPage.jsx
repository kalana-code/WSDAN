import React,{Component } from 'react';

import "./style/loading.css"

export default class LoadingPage extends Component {

    render(){
        return(
            <div className="centered">
                <div className="spinner">
                    <div className="double-bounce1"></div>
                    <div className="double-bounce2"></div>
                </div>
            </div>
        );
    }

} 