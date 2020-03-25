import React, { Component } from 'react';

class NodeStat extends Component {
    constructor(props) {
        super(props);
        this.state = {  }
    }

    render() {
        console.log(this.props) 
    return (  <p>{this.props.SelectedNodes.length() >0 && "Kalana" }</p> );
    }
}
 
export default NodeStat;