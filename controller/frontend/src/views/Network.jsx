/** @format */

import React, { Component } from "react";
import Graph from "vis-react";
import axios from "axios";
import { ButtonGroup, Button, Callout } from "@blueprintjs/core";

// import "./styles.css";
// need to import the vis network css in order to show tooltip
import "./style/network/css/network-manipulation.css";

import { Grid } from "react-bootstrap";
import NodeStat from "components/NodeStat/NodeStat.jsx";
import config from "./../config/config";

class Network extends Component {
  state = {
    graph: {
      nodes: [],
      edges: [],
    },
    noNode: true,
    selectNode: [],
    selectCount: 0, // used for select two nodes
    selectNodeError: null,
  };

  events = {
    select: function (event) {
      let { nodes } = event;
      this.setState({ selectNodeError: null });
      let selectNode = [];
      let selectCount = this.state.selectCount;
      if (nodes[0] !== undefined) {
        if (selectCount === 0) {
          selectNode[selectCount] = this.getNodeData(nodes[0]);
          console.log(this.getNodeData(nodes[0]));
          selectCount++;
        } else {
          selectNode = Object.assign(this.state.selectNode);
          console.log(selectNode[0].id);
          if (selectNode[0].id !== nodes[0]) {
            selectNode[selectCount] = this.getNodeData(nodes[0]);
            selectCount--;
          } else {
            this.setState({
              selectNodeError: {
                Error: true,
                Message: "Cannot Select Same",
              },
            });
          }
        }
      } else {
        // reset user select node when user select outside of nodes
        selectCount = 0;
        selectNode = Object.assign([]);
      }
      // change statesa
      this.setState({
        selectNode: selectNode,
        selectCount: selectCount,
      });
    }.bind(this),
  };

  options = {
    nodes: {
      shape: "dot",
      size: 14,
      font: {
        color: "gray",
        size: 12,
      },
    },
    edges: {
      font: {
        color: "gray",
        size: 12,
      },
    },
    layout:{

    improvedLayout:true
  },
    groups: {
      AP: {
        shape: "icon",
        icon: {
          face: "FontAwesome",
          code: "\uf2ce",
          size: 40,
          color: "#0F9960",
        },
      },
      NotAP: {
        shape: "icon",
        icon: {
          face: "FontAwesome",
          code: "\uf2ce",
          size: 40,
          color: "#A7B6C2",
        },
      },
      Controller: {
        shape: "icon",
        icon: {
          face: "FontAwesome",
          code: "\uf20e",
          size: 40,
          color: "#5C255C",
        },
      },
    },
    height: "400px",
    layout: {
      randomSeed: 55,
    },
    physics: {
      forceAtlas2Based: {
        gravitationalConstant: -26,
        centralGravity: 0.005,
        springLength: 230,
        springConstant: 0.18,
      },
      maxVelocity: 146,
      solver: "forceAtlas2Based",
      timestep: 0.35,
      stabilization: { iterations: 150 },
    },
  };

  getData = () => {
    axios.get(`http://` + config.host + `:8081/GetNodeInfo`).then(
      (response) => {
        if (response.status === 200) {
          if (response.data.Data.graphData.nodes.length > 0) {
            this.setState({ graph: response.data.Data.graphData });
            this.setState({ noNode: false });
          } else {
            this.setState({ noNode: true });
          }
        }
      },
      (error) => {
        console.log(error);
      }
    );
  };

  getEdges = (data) => {
    console.log(data);
  };

  getNodes = (data) => {
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
  }

  // we have  clear time interval
  // componentWillUnmount = () => {             // ***
  //   // Is our timer running?                 // ***
  //   if (this.timerHandle) {                  // ***
  //       // Yes, clear it                     // ***
  //       clearInterval(this.timerHandle);      // ***
  //       this.timerHandle = 0;                // ***
  //   }                                        // ***
  // };
  getNodeData = (NodeID) => {
    for (var key in this.state.graph.nodes) {
      var obj = this.state.graph.nodes[key];
      if (obj.id === NodeID) {
        return obj;
      }
    }
  };

  

  render() {
    return (
      <div className="content">
        <Grid fluid>
          <div className="graph_holder">
            <div className="refresh_button_holder">
              <ButtonGroup minimal={true}>
                <Button icon="refresh" onClick={this.getData}>
                  refresh
                </Button>
              </ButtonGroup>
            </div>
            {this.state.noNode ? (
              <div className="msg_text bp3-text-muted bp3-text-small">
                <div class="sk-cube-grid">
                  <div class="sk-cube sk-cube1"></div>
                  <div class="sk-cube sk-cube2"></div>
                  <div class="sk-cube sk-cube3"></div>
                  <div class="sk-cube sk-cube4"></div>
                  <div class="sk-cube sk-cube5"></div>
                  <div class="sk-cube sk-cube6"></div>
                  <div class="sk-cube sk-cube7"></div>
                  <div class="sk-cube sk-cube8"></div>
                  <div class="sk-cube sk-cube9"></div>
                </div>
                Waiting for Nodes Information....
              </div>
            ) : (
              <Graph
                graph={this.state.graph}
                options={this.options}
                events={this.events}
                getEdges={this.getEdges}
                getNodes={this.getNodes}
                vis={(vis) => (this.vis = vis)}
              />
            )}
          </div>
        </Grid>
        <NodeStat SelectedNodes={this.state.selectNode} />
      </div>
    );
  }
}

export default Network;
