/** @format */

import React, { Component } from "react";
import {
  H3,
  // H4,
  Divider,
  Tab,
  Tabs,
  Icon,
  ButtonGroup,
  Button,
  Callout,
} from "@blueprintjs/core";
import "./style.css";
import axios from "axios";
import RuleInsertForm from "./RuleInsertForm";
import config from "./../../config/config";

class RuleConfig extends Component {
  constructor(props) {
    super(props);
    this.state = {
      navbarTabId: "rs",
      isOpen: false,
      CurrentRules: [],
    };
  }

  handleNavbarTabChange = (navbarTabId) => this.setState({ navbarTabId });

  handleOpenRuleInsertForm = () =>
    this.setState({ isOpen: !this.state.isOpen });
  
  handleEvent =(event)=>{
    let id = event.target.name;

    if (id === undefined) {
        id = event.target.parentElement.parentElement.name;
    }

    if(id!=null){
      axios.get(`http://` + config.host + `:8081/RemoveRule/`+id).then((res) => {
      if (res.status === 200) {
        this.getData()
      }});
    }
  }
  getData=()=>{
    axios.get(`http://` + config.host + `:8081/GetAllRules`).then((res) => {
      if (res.status === 200) {
        console.log(res.data.Data);
        let rows = [];
        let index = 0;
        res.data.Data.Rules.forEach((element) => {
          index++;
          rows.push(
            <tr key={index}>
              <td>
                <ButtonGroup minimal={true} >
                  <Button>{element.Name}</Button>
                </ButtonGroup>
              </td>
              <td>
                <ButtonGroup minimal={true} >
                  <Button>{element.NodeIP}</Button>
                </ButtonGroup>
              </td>
              <td>
                <ButtonGroup minimal={true} >
                  <Button>{element.RuleId}</Button>
                </ButtonGroup>
              </td>
              <td>
                <ButtonGroup minimal={true} >
                  <Button>{element.FlowId}</Button>
                </ButtonGroup>
              </td>
              <td>
                <ButtonGroup minimal={true} >
                  <Button>{element.Protocol}</Button>
                </ButtonGroup>
              </td>

              <td>
                <ButtonGroup minimal={true} >
                  <Button>{element.DstIP}</Button>
                </ButtonGroup>
              </td>
              <td>
                <ButtonGroup minimal={true} >
                  <Button>{element.DstMAC}</Button>
                </ButtonGroup>
              </td>
              <td>
                <ButtonGroup minimal={true} >
                  <Button>{element.Interface}</Button>
                </ButtonGroup>
              </td>
              <td>
                <ButtonGroup minimal={true} >
                  <Button>
                    <b>{element.Action}</b>
                  </Button>
                </ButtonGroup>
              </td>

              <td>
                <ButtonGroup minimal={true} >
                  <Button icon={"dot"} intent="success">
                    Active
                  </Button>
                </ButtonGroup>
              </td>
              <td className="pull-right10">
                <ButtonGroup minimal={true} >
                  <Button name={element.RuleId}  onClick={this.handleEvent} icon="trash" />
                  <Divider />
                  <Button icon="swap-horizontal" />
                </ButtonGroup>
              </td>
            </tr>
          );
        });

        this.setState({
          CurrentRules: rows,
        });
      }
    });
  }
  componentDidMount = () => {
    this.getData()
  };
  render() {
    console.log(this.props);
    return (
      <Tabs
        id="TabsExample"
        onChange={this.handleNavbarTabChange}
        selectedTabId={this.state.navbarTabId}
      >
        <Tab
          id="rs"
          title="Rules"
          panel={
            <RulesPanel
              {...this.props}
              {...this.state}
              toggleDrawer={this.handleOpenRuleInsertForm}
              getRuleData={this.getData}
            />
          }
        />
        <Tab id="cn" disabled title="Configuration" panel={<ConfigPanel />} />
      </Tabs>
    );
  }
}

export default RuleConfig;
const RulesPanel = (props) => (
  <div>
    <div>
      <ButtonGroup minimal={true}>
        <Button
          disabled={props.SelectedNodes.length < 1}
          icon="add"
          onClick={props.toggleDrawer}
        >
          add rule
        </Button>
      </ButtonGroup>
     
      <Callout
        title={
          props.SelectedNodes.length > 0
            ? "Information About Selected Node"
            : "Please Select a Node"
        }
        icon={"info-sign"}
        intent={props.SelectedNodes.length > 0 ? "success" : "minimal"}
      >
        {props.SelectedNodes.length > 0 ? (
          <table>
            <tbody>
            {/* Network Address */}
            <tr>
              <td>
                <p className="bp3-text-small">
                  <b>IP Address</b>
                </p>
              </td>
              <td className="pull-right2">
                <p className="bp3-text-small">:</p>
              </td>
              <td className="pull-right5">
                <p className="bp3-text-small">
                  {props.SelectedNodes[props.SelectedNodes.length - 1].NodeData
                    .Node.IP === ""
                    ? "Configuring ..."
                    : props.SelectedNodes[props.SelectedNodes.length - 1]
                        .NodeData.Node.IP}
                </p>
              </td>
              <td className="pull-right5">
                <p className="bp3-text-small">
                  <b>MAC Address</b>
                </p>
              </td>
              <td className="pull-right2">
                <p className="bp3-text-small">:</p>
              </td>
              <td className="pull-right5">
                <p className="bp3-text-small">
                  {props.SelectedNodes[props.SelectedNodes.length - 1].NodeData
                    .Node.MAC === ""
                    ? "Configuring ..."
                    : props.SelectedNodes[props.SelectedNodes.length - 1]
                        .NodeData.Node.MAC}
                </p>
              </td>
            </tr></tbody>
          </table>
        ) : (
          ""
        )}
      </Callout>
      <hr />
      <div>
        {props.CurrentRules.length > 0 ? (
          <table className="bp3-html-table bp3-interactive">
            <thead>
              <tr>
                <th>Node Name</th>
                <th>Node IP</th>
                <th>RuleID</th>
                <th>FlowId</th>
                <th>Protocol</th>

                <th>Destination IP</th>
                <th>Direct to MAC</th>
                <th>Interface</th>
                <th>Action</th>
                <th>State</th>
                <th>Settings</th>
              </tr>
            </thead>
            <tbody>{props.CurrentRules}</tbody>
          </table>
        ) : (
          <p className="bp3-text-small">No Rules Found</p>
        )}
      </div>
    </div>
    <RuleInsertForm
      {...props}
      isOpen={props.isOpen}
      toggleDrawer={props.toggleDrawer}
      getRuleData={props.getRuleData}
    />
  </div>
);

const ConfigPanel = () => (
  <div>
    <H3>Backbone</H3>
  </div>
);

