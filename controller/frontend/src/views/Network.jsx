
import React, { Component } from "react";
import Graph from "react-graph-vis";
 
// import "./styles.css";
// need to import the vis network css in order to show tooltip
import "./style/network/css/network-manipulation.css";

import { Grid } from "react-bootstrap";

class Network extends Component {


  
  render() {
    return (
      <div className="content">
        <Grid fluid>
          <App />
        </Grid>
      </div>
    );
  }
}



 
function App() {
  const graph = {
    nodes: [
        { id: 0, label: "Node A", group: "Node" },
        { id: 1, label: "Gateway", group: "Gateway" },
        { id: 2, label: "Access Point", group: "AP" },
        { id: 3, label: "Node C", group: "Node" },
        { id: 4, label: "Node D", group: "Node" },
        { id: 5, label: "Node E", group: "Node" },
        { id: 6, label: "Node F", group: "Node" },
        { id: 7, label: "Laptop", group: "Laptop" },
        { id: 8, label: "Mobile", group: "Mobile" },
        { id: 9, label: "Sensor", group: "Sensor" },

      ],
    edges: [
        
        { from: 1, to: 3, label: "0.71 mbps",arrowStrikethrough:true,hoverWidth:6 ,
        arrows: {
            to: {
              enabled: false,
              imageHeight: undefined,
              imageWidth: undefined,
              scaleFactor: 1,
              src: undefined,
              type: "arrow"
            },
            middle: {
              enabled: false,
              imageHeight: 32,
              imageWidth: 32,
              scaleFactor: 1,
              src: "https://visjs.org/images/visjs_logo.png",
              type: "image"
            },
            from: {
              enabled: false,
              imageHeight: undefined,
              imageWidth: undefined,
              scaleFactor: 1,
              src: undefined,
              type: "arrow"
            }
          },
          arrowStrikethrough: true,
          chosen: true,
          color: {
            color:'#848484',
            highlight:'#848484',
            hover: '#848484',
            inherit: 'from',
            opacity:1.0
          },
          dashes: false,
          font: {
            color: '#343434',
            size: 14, // px
            face: 'arial',
            background: 'none',
            strokeWidth: 2, // px
            strokeColor: '#ffffff',
            align: 'horizontal',
            multi: false,
            vadjust: 0,
            bold: {
              color: '#343434',
              size: 14, // px
              face: 'arial',
              vadjust: 0,
              mod: 'bold'
            },
            ital: {
              color: '#343434',
              size: 14, // px
              face: 'arial',
              vadjust: 0,
              mod: 'italic',
            },
            boldital: {
              color: '#343434',
              size: 14, // px
              face: 'arial',
              vadjust: 0,
              mod: 'bold italic'
            },
            mono: {
              color: '#343434',
              size: 15, // px
              face: 'courier new',
              vadjust: 2,
              mod: ''
            }
          },
          hidden: false,
          hoverWidth: 1.5,
          label: undefined,
          labelHighlightBold: true,
          length: undefined,
          physics: true,
          scaling:{
            min: 1,
            max: 15,
            label: {
              enabled: true,
              min: 14,
              max: 30,
              maxVisible: 30,
              drawThreshold: 5
            },
            customScalingFunction: function (min,max,total,value) {
              if (max === min) {
                return 0.5;
              }
              else {
                var scale = 1 / (max - min);
                return Math.max(0,(value - min)*scale);
              }
            }
          },
          selectionWidth: 1,
          selfReferenceSize: 20,
          selfReference:{
              size: 20,
              angle: Math.PI / 4,
              renderBehindTheNode: true
          },
          shadow:{
            enabled: false,
            color: 'rgba(0,0,0,0.5)',
            size:10,
            x:5,
            y:5
          },
          smooth: {
            enabled: true,
            type: "dynamic",
            roundness: 0.5
          },
          title:undefined,
          value: undefined,
          width: 1,
          widthConstraint: false
        
        } ,
        { from: 4, to: 3, label: "0.71 mbps",arrowStrikethrough:true,hoverWidth:6 ,
        arrows: {
            to: {
              enabled: false,
              imageHeight: undefined,
              imageWidth: undefined,
              scaleFactor: 1,
              src: undefined,
              type: "arrow"
            },
            middle: {
              enabled: false,
              imageHeight: 32,
              imageWidth: 32,
              scaleFactor: 1,
              src: "https://visjs.org/images/visjs_logo.png",
              type: "image"
            },
            from: {
              enabled: false,
              imageHeight: undefined,
              imageWidth: undefined,
              scaleFactor: 1,
              src: undefined,
              type: "arrow"
            }
          },
          arrowStrikethrough: true,
          chosen: true,
          color: {
            color:'#848484',
            highlight:'#848484',
            hover: '#848484',
            inherit: 'from',
            opacity:1.0
          },
          dashes: false,
          font: {
            color: '#343434',
            size: 14, // px
            face: 'arial',
            background: 'none',
            strokeWidth: 2, // px
            strokeColor: '#ffffff',
            align: 'horizontal',
            multi: false,
            vadjust: 0,
            bold: {
              color: '#343434',
              size: 14, // px
              face: 'arial',
              vadjust: 0,
              mod: 'bold'
            },
            ital: {
              color: '#343434',
              size: 14, // px
              face: 'arial',
              vadjust: 0,
              mod: 'italic',
            },
            boldital: {
              color: '#343434',
              size: 14, // px
              face: 'arial',
              vadjust: 0,
              mod: 'bold italic'
            },
            mono: {
              color: '#343434',
              size: 15, // px
              face: 'courier new',
              vadjust: 2,
              mod: ''
            }
          },
          hidden: false,
          hoverWidth: 1.5,
          label: undefined,
          labelHighlightBold: true,
          length: undefined,
          physics: true,
          scaling:{
            min: 1,
            max: 15,
            label: {
              enabled: true,
              min: 14,
              max: 30,
              maxVisible: 30,
              drawThreshold: 5
            },
            customScalingFunction: function (min,max,total,value) {
              if (max === min) {
                return 0.5;
              }
              else {
                var scale = 1 / (max - min);
                return Math.max(0,(value - min)*scale);
              }
            }
          },
          selectionWidth: 1,
          selfReferenceSize: 20,
          selfReference:{
              size: 20,
              angle: Math.PI / 4,
              renderBehindTheNode: true
          },
          shadow:{
            enabled: false,
            color: 'rgba(0,0,0,0.5)',
            size:10,
            x:5,
            y:5
          },
          smooth: {
            enabled: true,
            type: "dynamic",
            roundness: 0.5
          },
          title:undefined,
          value: undefined,
          width: 1,
          widthConstraint: false
        
        } ,
        { from: 0, to: 1, label: "0.71 mbps",arrowStrikethrough:true,hoverWidth:6, 
        arrows: {
            to: {
              enabled: false,
              imageHeight: undefined,
              imageWidth: undefined,
              scaleFactor: 1,
              src: undefined,
              type: "arrow"
            },
            middle: {
              enabled: false,
              
            },
            from: {
              enabled: false,
              imageHeight: undefined,
              imageWidth: undefined,
              scaleFactor: 1,
              src: undefined,
              type: "arrow"}}}
        , 
        { from: 0, to: 3, label: "0.71 mbps",arrowStrikethrough:true,hoverWidth:6, 
        arrows: {
            to: {
              enabled: false,
              imageHeight: undefined,
              imageWidth: undefined,
              scaleFactor: 1,
              src: undefined,
              type: "arrow"
            },
            middle: {
              enabled: false,
              
            },
            from: {
              enabled: false,
              imageHeight: undefined,
              imageWidth: undefined,
              scaleFactor: 1,
              src: undefined,
              type: "arrow"}}}
        ,
        { from: 4, to: 5, label: "0.71 mbps",arrowStrikethrough:true,hoverWidth:6, 
        arrows: {
            to: {
              enabled: false,
              imageHeight: undefined,
              imageWidth: undefined,
              scaleFactor: 1,
              src: undefined,
              type: "arrow"
            },
            middle: {
              enabled: false,
              
            },
            from: {
              enabled: false,
              imageHeight: undefined,
              imageWidth: undefined,
              scaleFactor: 1,
              src: undefined,
              type: "arrow"}}}
        , 
        { from: 4, to: 6, label: "0.71 mbps",arrowStrikethrough:true,hoverWidth:6, 
        arrows: {
            to: {
              enabled: false,
              imageHeight: undefined,
              imageWidth: undefined,
              scaleFactor: 1,
              src: undefined,
              type: "arrow"
            },
            middle: {
              enabled: false,
              
            },
            from: {
              enabled: false,
              imageHeight: undefined,
              imageWidth: undefined,
              scaleFactor: 1,
              src: undefined,
              type: "arrow"}}}
        ,
        { from: 5, to: 0, label: "0.71 mbps",arrowStrikethrough:true,hoverWidth:6, 
        arrows: {
            to: {
              enabled: false,
              imageHeight: undefined,
              imageWidth: undefined,
              scaleFactor: 1,
              src: undefined,
              type: "arrow"
            },
            middle: {
              enabled: false,
              
            },
            from: {
              enabled: false,
              imageHeight: undefined,
              imageWidth: undefined,
              scaleFactor: 1,
              src: undefined,
              type: "arrow"}}}
        ,  
        { from: 3, to: 6, label: "0.71 mbps",arrowStrikethrough:true,hoverWidth:6, 
        arrows: {
            to: {
              enabled: false,
              imageHeight: undefined,
              imageWidth: undefined,
              scaleFactor: 1,
              src: undefined,
              type: "arrow"
            },
            middle: {
              enabled: false,
              
            },
            from: {
              enabled: false,
              imageHeight: undefined,
              imageWidth: undefined,
              scaleFactor: 1,
              src: undefined,
              type: "arrow"}}}
        ,
        { from: 2, to: 5, label: "0.71 mbps",arrowStrikethrough:true,hoverWidth:6, 
        arrows: {
            to: {
              enabled: false,
              imageHeight: undefined,
              imageWidth: undefined,
              scaleFactor: 1,
              src: undefined,
              type: "arrow"
            },
            middle: {
              enabled: false,
              
            },
            from: {
              enabled: false,
              imageHeight: undefined,
              imageWidth: undefined,
              scaleFactor: 1,
              src: undefined,
              type: "arrow"}}}
        , 
        { from: 2, to: 7, label: "0.71 mbps",arrowStrikethrough:true,hoverWidth:6, 
          arrows: {
            to: {
              enabled: false,
              imageHeight: undefined,
              imageWidth: undefined,
              scaleFactor: 1,
              src: undefined,
              type: "arrow"
            },
            middle: {
              enabled: false,
              
            },
            from: {
              enabled: false,
              imageHeight: undefined,
              imageWidth: undefined,
              scaleFactor: 1,
              src: undefined,
              type: "arrow"
            }
          }, },
        { from: 2, to: 8, label: "0.71 mbps",arrowStrikethrough:true,hoverWidth:6, 
          arrows: {
            to: {
              enabled: false,
              imageHeight: undefined,
              imageWidth: undefined,
              scaleFactor: 1,
              src: undefined,
              type: "arrow"
            },
            middle: {
              enabled: false,
              
            },
            from: {
              enabled: false,
              imageHeight: undefined,
              imageWidth: undefined,
              scaleFactor: 1,
              src: undefined,
              type: "arrow"
            }
          }, },
          { from: 2, to: 9, label: "0.71 mbps",arrowStrikethrough:true,hoverWidth:6, 
          arrows: {
            to: {
              enabled: false,
              imageHeight: undefined,
              imageWidth: undefined,
              scaleFactor: 1,
              src: undefined,
              type: "arrow"
            },
            middle: {
              enabled: false,
              
            },
            from: {
              enabled: false,
              imageHeight: undefined,
              imageWidth: undefined,
              scaleFactor: 1,
              src: undefined,
              type: "arrow"
            }
          }, }
        

      ]
  };
 
  const options = {
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
        enabled:false
    //   forceAtlas2Based: {
    //     gravitationalConstant: -26,
    //     centralGravity: 0.005,
    //     springLength: 230,
    //     springConstant: 0.18
    //   },
    //   maxVelocity: 146,
    //   solver: "forceAtlas2Based",
    //   timestep: 0.35,
    //   stabilization: {
    //     enabled: true,
    //     iterations: 2000,
    //     updateInterval: 25
    //   }
    }
  };
 
  const events = {
    select: function(event) {
      var { nodes, edges } = event;
      console.log(event)
    }
  };
  return (
    <Graph
      graph={graph}
      options={options}
      events={events}
      getNetwork={network => {
        //  if you want access to vis.js network api you can set the state in a parent component using this property
      }}
    />
  );
}

export default Network;
