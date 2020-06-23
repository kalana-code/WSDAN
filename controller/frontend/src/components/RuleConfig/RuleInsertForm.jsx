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

import config from "./../../config/config";
class RuleInsertForm extends Component {
  state = {
    destinationMac: [],
    isLoading: false,
    isDisable: false,
    FlowData:null,
    disableIPfield:false,
    IPfieldLoading:false,
    NodeIP: "",
    FlowID: "",
    Protocol: "",
    SrcIP:"",
    DstIP: "",
    Interface: "",
    DstMAC: "",
    Action: "ACCEPT",
    Error: {
      NodeIP: false,
      FlowID: false,
      SrcIP:false,
      DstIP: false,
      Interface: false,
      DstMAC: false,
    },
    ErrorMessage: {
      NodeIP: "",
      FlowID: "",
      SrcIP:"",
      DstIP: "",
      Interface: "",
      DstMAC: "",
    },
  };

  // handle Functions
  handleChange = (event) => {
    const { name, value } = event.target;
    if (name === "FlowID"){
        if(this.state.FlowData[value]!=undefined){
          this.setState({
            SrcIP:this.state.FlowData[value].SrcIP,
            DstIP:this.state.FlowData[value].DstIP,
            disableIPfield:true
          })
        }else{
          this.setState({
            SrcIP:"",
            DstIP: "",
            disableIPfield:false
          })
        }
    }
    this.setState({ [name]: value });
  };
isEmpty=(str)=> {
    return (!str || 0 === str.length);
}

  VerifyAndSubmit = () => {
    let Error = {
      NodeIP: false,
      FlowID: false,
      DstIP: false,
      SrcIP:false,
      Interface: false,
      DstMAC: false,
    };
    let ErrorMessage = {
      NodeIP: "",
      FlowID: "",
      DstIP: "",
      SrcIP:"",
      Interface: "",
      DstMAC: "",
    };

    if (this.isEmpty(this.state.FlowID)) {
      Error.FlowID = true;
      ErrorMessage.FlowID = "FlowID cannot be empty";
    }
    if (this.isEmpty(this.state.DstIP)) {
      Error.DstIP = true;
      ErrorMessage.DstIP = "DstIP cannot be empty";
    }

    if (this.isEmpty(this.state.DstMAC) && this.state.Action !=="DROP" ) {
      Error.DstMAC = true;
      ErrorMessage.DstMAC = "DstMAC cannot be empty";
    }
   

    if (this.isEmpty(this.state.Interface)) {
      Error.Interface = true;
      ErrorMessage.Interface = "Interface cannot be empty";
    }

    if (this.isEmpty(this.state.Protocol)) {
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

    if (
      !this.state.SrcIP.match(
        "^((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?).){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$"
      )
    ) {
      Error.SrcIP = true;
      ErrorMessage.SrcIP = "Invalid IP4 Address";
    } 
    console.log(Error)

    // set State
    this.setState({ Error: Error, ErrorMessage: ErrorMessage },()=>{
      this.DataSubmit()
    });
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

    let isValid = true;
    console.log(this.state.Error)
    for (const property in this.state.Error) {
      console.log(this.state.Error[property])
       if(this.state.Error[property] ===  true){
        isValid = false;
        break
       }
    }
    
    console.log(isValid)
    // Submit Data
    if (isValid) {
      const Request_Body = {
        NodeIP: this.props.SelectedNodes[this.props.SelectedNodes.length - 1]
          .NodeData.Node.IP,
        NodeName: this.props.SelectedNodes[this.props.SelectedNodes.length - 1]
          .NodeData.Node.Name,
        FlowID: this.state.FlowID,
        Protocol: this.state.Protocol,
        ScrIP: this.state.ScrIP,
        DstIP: this.state.DstIP,
        Interface: this.state.Interface,
        DstMAC: this.state.DstMAC,
        Action: this.state.Action,
        IsActive : true
      };
      this.setState({ isLoading: true });
      axios.post(`http://` + config.host + `:8081/AddRule`, Request_Body).then(
        (response) => {
          if (response.status === 200) {
            this.setState({ isLoading: false });
            this.props.getRuleData();
          }
        },
        (error) => {
          this.setState({ isLoading: false });
        }
      );
      this.props.toggleDrawer();
      this.setState({
        destinationMac: [],
        isLoading: false,
        isDisable: false,
        NodeIP: "",
        SrcMAC: "",
        FlowID: "",
        Protocol: "",
        DstIP: "",
        Interface: "",
        DstMAC: "",
        Action: "ACCEPT",
        Error: {
          NodeIP: false,
          FlowID: false,
          SrcMAC:false,
          DstIP: false,
          Interface: false,
          DstMAC: false,
        },
        ErrorMessage: {
          NodeIP: "",
          SrcMAC:"",
          FlowID: "",
          DstIP: "",
          Interface: "",
          DstMAC: "",
        },
      });
    }
  };

  GetFlowData=()=>{
    axios.get(`http://`+config.host+`:8081/GetFlowData`).then((res) => {
      if (res.status === 200) {
        console.log(res.data.Data.FlowData)
        this.setState({
          FlowData: res.data.Data.FlowData,
        });
      }
    });
  }

  render() {
    let NodeIP = "";
    let NeighboursMac = [];
    if (this.props.SelectedNodes.length > 0) {
      NodeIP = this.props.SelectedNodes[this.props.SelectedNodes.length - 1]
        .NodeData.Node.IP;
      let a = this.props.SelectedNodes[this.props.SelectedNodes.length - 1]
        .NodeData.Neighbours;
      if (a !== null) {
        NeighboursMac.push(
          <option key={0} value="">
            {" "}
            --Select DstMAC--
          </option>
        );
        for (let i = 0; i < a.length; i++) {
          console.log(a)
          if (this.props.nodeNames[a[i].MAC] != null && a[i].MAC!=this.props.contrllerMAC) {
            NeighboursMac.push(
              <option key={i + 1} value={a[i].MAC}>
                {this.props.nodeNames[a[i].MAC]}
              </option>
            );
          }
        }
      }
    }

    return (
      <Drawer
        icon="info-sign"
        onClose={this.props.toggleDrawer}
        onOpening={this.GetFlowData}
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
                  intent={this.getIntent("FlowID")}
                >
                  <div className="bp3-select">
                    <select
                      name="FlowID"
                      onChange={this.handleChange}
                      value={this.state.FlowID}
                      intent={this.getIntent("FlowID")}
                    >
                      <option value="">--select a flow id--</option>
                      <option value="F0">Flow 00</option>
                      <option value="F1">Flow 01</option>
                      <option value="F2">Flow 02</option>
                      <option value="F3">Flow 03</option>
                      <option value="F4">Flow 04</option>
                      <option value="F5">Flow 05</option>
                      <option value="F6">Flow 06</option>
                      <option value="F7">Flow 07</option>
                      <option value="F8">Flow 08</option>
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
                      <option value="UDP">UDP Protocol</option>
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
                  helperText={this.state.ErrorMessage.SrcIP}
                  label="Source IP address"
                  intent={this.getIntent("SrcIP")}
                  disabled={this.state.disableIPfield}
                >
                  <InputGroup
                    name="SrcIP"
                    disabled={this.state.disableIPfield}
                    onChange={this.handleChange}
                    value={this.state.SrcIP}
                    id="text-input"
                    placeholder="192.168.3.2"
                    intent={this.getIntent("ScrIP")}
                  />
                </FormGroup>
                <FormGroup
                  helperText={this.state.ErrorMessage.DstIP}
                  label="Destination IP address"
                  intent={this.getIntent("DstIP")}
                  disabled={this.state.disableIPfield}
                >
                  <InputGroup
                    name="DstIP"
                    disabled={this.state.disableIPfield}
                    onChange={this.handleChange}
                    value={this.state.DstIP}
                    id="text-input"
                    placeholder="192.168.1.1"
                    intent={this.getIntent("DstIP")}
                  />
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
                <Callout title={"Note"} icon={"info-sign"} intent={"success"}>
                  Message About Project .
                </Callout>
                {this.state.Action != "DROP" && (
                  <FormGroup
                    helperText={this.state.ErrorMessage.DstMAC}
                    label="Next Node"
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
                )}
              </div>
            ) : (
              <Callout title={"Note"} icon={"error"} intent={"danger"}>
                You selected node is being configured.
              </Callout>
            )}
          </div>
        </div>

        <div className={Classes.DRAWER_FOOTER}>
          {NodeIP !== "" && (
            <FormGroup>
              <Button
                intent="primary"
                text="Add Rule"
                onClick={this.VerifyAndSubmit}
                disabled={false}
                loading={this.state.isLoading}
              />
            </FormGroup>
          )}
        </div>
      </Drawer>
    );
  }
}

export default RuleInsertForm;
