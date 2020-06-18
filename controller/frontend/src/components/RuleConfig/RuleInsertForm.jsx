/** @format */

import axios from "axios";
import React, { Component } from "react";
import {
  Drawer,
  Classes,
  FormGroup,
  InputGroup,
  Button,
  Callout,

} from "@blueprintjs/core";

import config from './../../config/config'
class RuleInsertForm extends Component {
  state = {
    destinationMac: [],
    isLoading: false,
    isDisable: false,
    NodeIP: "",
    FlowID: "",
    Protocol: "",
    DstIP: "",
    Interface: "",
    DstMAC: "",
    Action: "ACCEPT",
    Error: {
      NodeIP: false,
      FlowID: false,
      DstIP: false,
      Interface: false,
      DstMAC: false,
    },
    ErrorMessage: {
      NodeIP: "",
      FlowID: "",
      DstIP: "",
      Interface: "",
      DstMAC: "",
    },
  };

  // handle Functions
  handleChange = (event) => {
    const { name, value } = event.target;
    this.setState({ [name]: value }, () => this.Verify());
    console.log(this.state);
  };

  Verify = () => {
    console.log(this.state);
    let Error = {
      NodeIP: false,
      FlowID: false,
      DstIP: false,
      Interface: false,
      DstMAC: false,
    };
    let ErrorMessage = {
      NodeIP: "",
      FlowID: "",
      DstIP: "",
      Interface: "",
      DstMAC: "",
    };

    //Check Email
    if (this.state.NodeIP === "") {
      Error.NodeIP = true;
      ErrorMessage.NodeIP = "NodeIP cannot be empty";
    }
    if (this.state.FlowID === "") {
      Error.FlowID = true;
      ErrorMessage.FlowID = "FlowID cannot be empty";
    }
    if (this.state.DstIP === "") {
      Error.DstIP = true;
      ErrorMessage.DstIP = "DstIP cannot be empty";
    }

    if (this.state.DstMAC === "") {
        Error.DstMAC = true;
        ErrorMessage.DstMAC = "DstMAC cannot be empty";
      }
    if (this.state.DstMAC === null) {
        Error.DstMAC = true;
        ErrorMessage.DstMAC = "DstIP cannot be empty";
    }

    if (this.state.Interface === "") {
        Error.Interface = true;
        ErrorMessage.Interface = "Interface cannot be empty";
    }

    if (this.state.Protocol === "") {
        Error.Protocol = true;
        ErrorMessage.Protocol = "Protocol cannot be empty";
    }
    if (
      !this.state.DstIP.match(
        "^((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?).){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$"
      )
    ) {
      Error.DstIP = true;
      ErrorMessage.DstIP = "Invalid IP4 Address";
    }
    if (this.state.Interface === "") {
      Error.Interface = true;
      ErrorMessage.Interface = "Interface cannot be empty";
    }

    // set State
    this.setState({ Error: Error, ErrorMessage: ErrorMessage });
  };

  // get intent
  getIntent = (feild) => {
    if (this.state.Error[feild]) {
      return "danger";
    }
    return "primary";
  };

  //Send Request
  DataSubmit = () => {
    this.Verify();
    let isValid = true;
    //check form input errors
    Object.keys(this.state.Error).map(
      (value) =>
        function () {
          if (this.state.Error[value]) {
            isValid = false;
          }
        }
    );

    // Submit Data
    if (isValid) {
      const Request_Body = {
        NodeIP: this.props.SelectedNodes[this.props.SelectedNodes.length - 1].NodeData.Node.IP,
        FlowID: this.state.FlowID,
        Protocol: this.state.Protocol,
        DstIP: this.state.DstIP,
        Interface: this.state.DstMAC,
        DstMAC: this.state.DstMAC,
        Action: this.state.Action, 
      };
      this.setState({ isLoading: true });
      axios.post(`http://`+config.host+`:8081/AddRule`, Request_Body).then(
        (response) => {
          if (response.status === 200) {
            this.setState({ isLoading: false });
          }
        },
        (error) => {
          this.setState({ isLoading: false });
        }
      );
      this.props.toggleDrawer()
      this.setState({
        destinationMac: [],
        isLoading: false,
        isDisable: false,
        NodeIP: "",
        FlowID: "",
        Protocol: "",
        DstIP: "",
        Interface: "",
        DstMAC: "",
        Action: "ACCEPT",
        Error: {
          NodeIP: false,
          FlowID: false,
          DstIP: false,
          Interface: false,
          DstMAC: false,
        },
        ErrorMessage: {
          NodeIP: "",
          FlowID: "",
          DstIP: "",
          Interface: "",
          DstMAC: "",
        }
      })
    }
  };

  render() {
    let NodeIP = "";
    let NeighboursMac =[]
    if (this.props.SelectedNodes.length > 0) {
      NodeIP = this.props.SelectedNodes[this.props.SelectedNodes.length - 1].NodeData.Node.IP
      let a= this.props.SelectedNodes[this.props.SelectedNodes.length - 1].NodeData.Neighbours
      if(a!==null){
        NeighboursMac.push(<option  key={0} value=""> --Select DstMAC--</option>)
        for(let i = 0;i<a.length;i++){
            
            NeighboursMac.push(<option key ={i+1} value={a[i].MAC}>{a[i].MAC}</option>)
        }
      }
    }

    return (
      <Drawer
        icon="info-sign"
        onClose={this.props.toggleDrawer}
        title="Rule Insert"
        isOpen={this.props.isOpen}
      >
        <div className={Classes.DRAWER_BODY}>
          <div className={Classes.DIALOG_BODY}>
            {NodeIP !== "" ? (
              <div>
                <Callout title={"Note"} icon={"info-sign"} intent={"success"}>
                  Message About Project .
                </Callout>
                <FormGroup
                  helperText={this.state.ErrorMessage.FlowID}
                  label="Flow ID"
                >
                  <div className="bp3-select">
                    <select
                      name="FlowID"
                      onChange={this.handleChange}
                      value={this.state.FlowID}
                      intent={this.getIntent("FlowID")}
                    >
                      <option value="">--select a flow id--</option>
                      <option value="F0">
                        Flow 00
                      </option>
                      <option value="F1">Flow 01</option>
                      <option value="F2">Flow 01</option>
                    </select>
                  </div>
                </FormGroup>

                <FormGroup
                  helperText={this.state.ErrorMessage.Protocol}
                  label="Protocol"
                  intent={this.getIntent("Protocol")}
                >
                  <div className="bp3-select">
                    <select
                      name="Protocol"
                      onChange={this.handleChange}
                      value={this.state.Protocol}
                      intent={this.getIntent("Protocol")}
                    >
                    <option value="">--select protocol--</option>
                      <option value="ICMPv4">ICMP Protocol</option>
                      <option value="UDP">
                        UDP Protocol
                      </option>
                      {/* <option value="">UPD Protocol</option> */}
                    </select>
                  </div>
                </FormGroup>

                <FormGroup
                  helperText={this.state.ErrorMessage.Interface}
                  label="Interface"
                  intent={this.getIntent("Interface")}
                >
                  <div className="bp3-select">
                    <select
                      name="Interface"
                      onChange={this.handleChange}
                      value={this.state.Interface}
                      intent={this.getIntent("Interface")}
                    >
                    <option value="">--select interface--</option>
                      <option value="wlan0">WLAN0</option>
                    </select>
                  </div>
                </FormGroup>
                <FormGroup
                  helperText={this.state.ErrorMessage.DstIP}
                  label="Destination IP address"
                  intent={this.getIntent("DstIP")}
                >
                  <InputGroup
                    name="DstIP"
                    onChange={this.handleChange}
                    value={this.state.DstIP}
                    id="text-input"
                    placeholder="192.168.1.1"
                    intent={this.getIntent("DstIP")}
                  />
                </FormGroup>

                <Callout title={"Note"} icon={"info-sign"} intent={"success"}>
                  Message About Project .
                </Callout>
                <FormGroup
                  helperText={this.state.ErrorMessage.DstMAC}
                  label="Destination MAC address"
                  intent={this.getIntent("DstMAC")}
                >
                  <div className="bp3-select">
                    <select
                      name="DstMAC"
                      onChange={this.handleChange}
                      value={this.state.DstMAC}
                      intent={this.getIntent("DstMAC")}
                    >
                      {NeighboursMac}
                    </select>
                  </div>
                </FormGroup>
                <Callout
                  title={"What is Action"}
                  icon={"info-sign"}
                  intent={"success"}
                >
                  Message About Project .
                </Callout>
                <FormGroup
                  helperText={this.state.ErrorMessage.Action}
                  label="Action"
                >
                  <div className="bp3-select">
                    <select
                      name="Action"
                      onChange={this.handleChange}
                      value={this.state.Action}
                      intent={this.getIntent("Action")}
                    >
                      <option value="ACCEPT">ACCEPT</option>
                      <option value="DROP">DROP</option>
                    </select>
                  </div>
                </FormGroup>
              </div>
            ) : (
                <Callout title={"Note"} icon={"error"} intent={"danger"}>
                You selected node is being configured.
              </Callout>
            )}
          </div>
        </div>

        <div className={Classes.DRAWER_FOOTER}>
            {NodeIP !== ""&&<FormGroup>
            <Button
              
              intent="primary"
              text="Add Rule"
              onClick={this.DataSubmit}
              disabled={false}
              loading={this.state.isLoading}
            />
          </FormGroup>}
        </div>
      </Drawer>
    );
  }
}

export default RuleInsertForm;
