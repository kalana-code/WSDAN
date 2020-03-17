
import React, { Component } from "react";
import Graph from "react-graph-vis";
import axios from "axios"
 
// import "./styles.css";
// need to import the vis network css in order to show tooltip
import "./style/network/css/network-manipulation.css";

import { Grid } from "react-bootstrap";

class Network extends Component {
  state ={
    graph : {
      nodes: [],
      edges: []
    }
  }
  options = {
    nodes: {
      shape: "dot",
      size: 16
    },
    groups: {
        switch: {
          shape: "icon",
          color: "#FF9900", // orange
          icon: {
            face: "FontAwesome",
            code: "\uf10b",
            size: 50,
            color: "orange"
          }
        },
        Node: {
            shape: "icon",
            color: "#FF9900", // orange
            icon: {
              face: "FontAwesome",
              code: "\uf21b",
              size: 50,
              color: "gray"
            }
          },
          Gateway: {
            shape: "icon",
            color: "#FF9900", // orange
            icon: {
              face: "FontAwesome",
              code: "\uf233",
              size: 50,
              color: "gray"
            }
          },
          AP: {
            shape: "icon",
            color: "#FF9900", // orange
            icon: {
              face: "FontAwesome",
              code: "\uf1eb",
              size: 50,
              color: "gray"
            }
          },
          Laptop: {
            shape: "icon",
            color: "#FF9900", // orange
            icon: {
              face: "FontAwesome",
              code: "\uf109",
              size: 50,
              color: "gray"
            }
          }, Mobile: {
            shape: "icon",
            color: "#FF9900", // orange
            icon: {
              face: "FontAwesome",
              code: "\uf10b",
              size: 50,
              color: "gray"
            }
          },
          Sensor: {
            shape: "icon",
            color: "#FF9900", // orange
            icon: {
              face: "FontAwesome",
              code: "\uf512",
              size: 50,
              color: "gray"
            }
          },
       
      },
    height: "500px",
    layout: {
      randomSeed: 34
    },
    physics: {
        enabled:true,
      forceAtlas2Based: {
        gravitationalConstant: -26,
        centralGravity: 0.005,
        springLength: 230,
        springConstant: 0.18
      },
      maxVelocity: 146,
      solver: "forceAtlas2Based",
      timestep: 0.35,
      stabilization: {
        enabled: true,
        iterations: 2000,
        updateInterval: 25
      }
    }
  };
  getData=()=>{
    axios.get(`http://localhost:8081/GetNodeInfo`).then(response => {
      if(response.status === 200){
        this.setState({graph:response.data.Data.graphData});       
      }
    },error=>{
      console.log(error);
    });
  }

async componentDidMount() {
    try {
      setInterval(async () => {
        this.getData();
      }, 3000);
    } catch(e) {
      console.log(e);
    }
}

  render() {
    return (
      <div className="content">
        <Grid fluid>
          {/* <App /> */}
          <Graph
            graph={this.state.graph}
            options={this.options}
            //events={events}
            getNetwork={network => {
              //  if you want access to vis.js network api you can set the state in a parent component using this property
            }}
          />
        </Grid>
      </div>
    );
  }
}

export default Network;
