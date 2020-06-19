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
      axios.get(`http://` + config.host + `:8081//RemoveRule`+id).then((res) => {
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
        <Button icon="database">get rules</Button>
        <Button icon="trash">delete</Button>
      </ButtonGroup>
      <hr />
      <Callout
        title={
          props.SelectedNodes.length > 0
            ? "Information About Selected Node"
            : "Please Select a Node"
        }
        icon={"info-sign"}
        intent={props.SelectedNodes.length > 0 ? "success" : "warning"}
      >
        {props.SelectedNodes.length > 0 ? (
          <table>
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
            </tr>
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
    />
  </div>
);

const ConfigPanel = () => (
  <div>
    <H3>Backbone</H3>
  </div>
);

// const RulesPanel = (props) => (
//   <div>
//     <div>
//       <H4>Rule Manager</H4>

//       {props.SelectedNodes.length > 0 ? (
//         <div>
//           <hr />
//           <table className="bp3-text-muted">
//             <tr>
//               <td>
//                 <p className="bp3-text">
//                   <b>Rule Manager</b>
//                 </p>
//               </td>
//               <td className="pull-right2">
//                 <p className="bp3-text-small">:</p>
//               </td>
//             </tr>
//             <tr>
//               <th>
//                 <p className="bp3-text-small">
//                   <b>Rule</b>
//                 </p>
//               </th>
//               <th className="pull-right10">
//                 <p className="bp3-text-small">
//                   <b>Hit Count</b>
//                 </p>
//               </th>
//               <th className="pull-right10">
//                 <p className="bp3-text-small">
//                   <b>State</b>
//                 </p>
//               </th>
//               <th className="pull-right10"></th>
//             </tr>
//             <tr>
//               <td>
//                 <p className="bp3-text-small"># Rule 00123</p>
//               </td>
//               <td className="pull-right10">
//                 <p className="bp3-text-small">
//                   <b>5</b>
//                   <Trend
//                     smooth
//                     height={35}
//                     width={70}
//                     autoDraw
//                     // autoDrawDuration={3000}
//                     autoDrawEasing="ease-out"
//                     data={[0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 1, 1, 1, 0]}
//                     gradient={["#106BA3"]}
//                     radius={10}
//                     strokeWidth={0.5}
//                     strokeLinecap={"butt"}
//                   />
//                 </p>
//               </td>
//               <td className="pull-right10">
//                 <p className="bp3-text-small">
//                   <Icon icon={"dot"} iconSize={17} intent="success" />
//                   <b>Active</b>
//                 </p>
//               </td>
//               <td className="pull-right10">
//                 <ButtonGroup minimal={true} >
//                   <Button icon="cube-add" />
//                   <Divider />
//                   <Button icon="cube-remove" />
//                   <Divider />
//                   <Button icon="swap-horizontal" />
//                 </ButtonGroup>
//               </td>
//             </tr>
//             <tr>
//               <td>
//                 <p className="bp3-text-small"># Rule 00124</p>
//               </td>
//               <td className="pull-right10">
//                 <p className="bp3-text-small">
//                   <b>9</b>
//                   <Trend
//                     smooth
//                     height={35}
//                     width={70}
//                     autoDraw
//                     // autoDrawDuration={3000}
//                     autoDrawEasing="ease-out"
//                     data={[0, 1, 0, 0, 0, 0, 0, 2, 0, 0, 0, 4, 1, 1, 0]}
//                     gradient={["#752F75"]}
//                     radius={10}
//                     strokeWidth={0.5}
//                     strokeLinecap={"butt"}
//                   />
//                 </p>
//               </td>
//               <td className="pull-right10">
//                 <p className="bp3-text-small">
//                   <Icon icon={"dot"} iconSize={17} intent="danger" />
//                   <b>Inactive</b>
//                 </p>
//               </td>
//               <td className="pull-right10"></td>
//             </tr>
//             <tr>
//               <td>
//                 <p className="bp3-text-small"># Rule 00130</p>
//               </td>
//               <td className="pull-right10">
//                 <p className="bp3-text-small">
//                   <b>0</b>
//                   <Trend
//                     smooth
//                     height={35}
//                     width={70}
//                     autoDraw
//                     // autoDrawDuration={3000}
//                     autoDrawEasing="ease-out"
//                     data={[0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0]}
//                     gradient={["#752F75"]}
//                     radius={10}
//                     strokeWidth={0.5}
//                     strokeLinecap={"butt"}
//                   />
//                 </p>
//               </td>
//               <td className="pull-right10">
//                 <p className="bp3-text-small">
//                   <Icon icon={"dot"} iconSize={17} intent="success" />
//                   <b>Active</b>
//                 </p>
//               </td>
//               <td className="pull-right10"></td>
//             </tr>
//           </table>
//           <hr />
//           <table className="bp3-text-muted">
//             <tr>
//               <td>
//                 <p className="bp3-text">
//                   <b>Rule Setter</b>
//                 </p>
//               </td>
//               <td className="pull-right2">
//                 <p className="bp3-text-small">:</p>
//               </td>
//             </tr>
//           </table>
//         </div>
//       ) : (
//         <p class="bp3-text-muted bp3-text-small">Please select a Node</p>
//       )}
//     </div>
//   </div>
// );
