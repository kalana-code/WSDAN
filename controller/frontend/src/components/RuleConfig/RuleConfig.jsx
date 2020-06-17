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
} from "@blueprintjs/core";
import "./style.css";
// import Trend from "react-trend";
import RuleInsertForm from "./RuleInsertForm";

class RuleConfig extends Component {
  constructor(props) {
    super(props);
    this.state = {
      navbarTabId: "rs",
      isOpen: false,
    };
  }

  handleNavbarTabChange = (navbarTabId) => this.setState({ navbarTabId });

  handleOpenRuleInsertForm = () =>
    this.setState({ isOpen: !this.state.isOpen });

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
      {props.SelectedNodes.length > 0 ? (
        <table className="bp3-text-muted">
          <tr>
            <td>
              <p className="bp3-text-small"># Rule 00123</p>
            </td>

            <td className="pull-right10">
              <p className="bp3-text-small">
                <Icon icon={"dot"} iconSize={17} intent="success" />
                <b>Active</b>
              </p>
            </td>
            <td className="pull-right10">
              <ButtonGroup minimal={true} small={true}>
                <Button icon="cube-add" />
                <Divider />
                <Button icon="cube-remove" />
                <Divider />
                <Button icon="swap-horizontal" />
              </ButtonGroup>
            </td>
          </tr>
        </table>
      ) : (
        <p className="bp3-text-muted bp3-text-small">Please select a Node</p>
      )}
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
//                 <ButtonGroup minimal={true} small={true}>
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
