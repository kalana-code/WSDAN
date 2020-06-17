
import React, { Component } from "react";
import Graph from "vis-react";
import axios from "axios"
import { ButtonGroup, Button } from "@blueprintjs/core";
 
// import "./styles.css";
// need to import the vis network css in order to show tooltip
import "./style/network/css/network-manipulation.css";

import { Grid } from "react-bootstrap";
import NodeStat from "components/NodeStat/NodeStat.jsx";
import config from "./../config/config"

class Network extends Component {
state ={
    graph : {
      nodes: [],
      edges: []
    }
    ,noNode:true
    ,selectNode:[]
    ,selectCount:0 // used for select two nodes
    ,selectNodeError:null
};


events = {
    select: function(event) {
      let { nodes } = event;
      this.setState({selectNodeError:null})
      let selectNode =[]
      let selectCount  = this.state.selectCount;
      if(nodes[0] !== undefined){
        if(selectCount === 0){
          selectNode[selectCount] = this.getNodeData(nodes[0]);
          console.log(this.getNodeData(nodes[0]))
          selectCount++
        }else{
            selectNode = Object.assign(this.state.selectNode);
            console.log(selectNode[0].id)
            if(selectNode[0].id !== nodes[0] ){
              selectNode[selectCount] = this.getNodeData(nodes[0]);
              selectCount--
            }else{
              this.setState({selectNodeError:{
                Error:true,
                Message:"Cannot Select Same"
              }})
            }
            
        }
       
      }else{
        // reset user select node when user select outside of nodes
        selectCount = 0  
        selectNode  = Object.assign([]);
      }
      // change statesa
      this.setState({
        selectNode:selectNode,selectCount:selectCount
      })   
    }.bind(this),
};


   
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
        Agents: {
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
          NotAP: {
            shape: "icon",
            color: "#FF9900", // orange
            icon: {
              face: "FontAwesome",
              code: "\uf2da",
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
    height: "300px",
    layout: {
      randomSeed: 55
    },
    physics: {
        enabled:true,
      // forceAtlas2Based: {
      //   gravitationalConstant: -26,
      //   centralGravity: 0.005,
      //   springLength: 230,
      //   springConstant: 0.18
      // },
      // maxVelocity: 146,
      // solver: "forceAtlas2Based",
      // timestep: 0.35,
      // stabilization: {
      //   enabled: true,
      //   iterations: 20,
      //   updateInterval: 25
      // }
    }
};

getData=()=>{
    axios.get(`http://`+config.host+`:8081/GetNodeInfo`).then(response => {
      if(response.status === 200){
        if(response.data.Data.graphData.nodes.length>0){
          this.setState({graph:response.data.Data.graphData});
          this.setState({noNode:false})
        }else{
          this.setState({noNode:true})
        }

      }
    },error=>{
      console.log(error);
    });
};

getEdges = data => {
  console.log(data);
};

getNodes = data => {
  console.log(data);
};

async componentDidMount() {
    // try {
    //   setInterval(async () => {
        this.getData();
    //   }, 1000);
    // } catch(e) {
    //   console.log(e);
    // }
};

// we have  clear time interval 
// componentWillUnmount = () => {             // ***
//   // Is our timer running?                 // ***
//   if (this.timerHandle) {                  // ***
//       // Yes, clear it                     // ***
//       clearInterval(this.timerHandle);      // ***
//       this.timerHandle = 0;                // ***
//   }                                        // ***
// };
getNodeData=(NodeID)=>{
  for (var key in this.state.graph.nodes) {
    var obj = this.state.graph.nodes[key];
    if(obj.id ===NodeID){
      return obj;
    }
    
  }

  return "kalana"

}

  render() {
    return (
      <div className="content">

        <Grid fluid>
        <ButtonGroup minimal={true}>
            <Button icon="refresh" onClick={this.getData}>refresh</Button>
          </ButtonGroup>
          <p>{this.setState.noNode ? "Not identified any node": "" }</p>
          <Graph
            graph={this.state.graph}
            options={this.options}
            events={this.events}
            getEdges={this.getEdges}
            getNodes={this.getNodes}
            vis={vis => (this.vis = vis)}
          />
        </Grid>
        <NodeStat SelectedNodes ={this.state.selectNode}/>
      </div>
    );
  }
}

export default Network;
