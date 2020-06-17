/** @format */

import React, { Component } from "react";
import {
  H3,
  H4,
  Tab,
  Tabs,
  Icon,

} from "@blueprintjs/core";
import "./style.css";
import Trend from "react-trend";

class NodeStat extends Component {
  constructor(props) {
    super(props);
    this.state = { navbarTabId: "ni" };
  }

  handleNavbarTabChange = (navbarTabId) => this.setState({ navbarTabId });

  render() {
    if (this.props.SelectedNodes.length > 0) {
      console.log(this.props.SelectedNodes[0].NodeData.Node.MAC);
    }

    return (
      <Tabs
        id="TabsExample"
        onChange={this.handleNavbarTabChange}
        selectedTabId={this.state.navbarTabId}
      >
        <Tab id="ni" title="Info" panel={<InfoPanel {...this.props} />} />
        <Tab
          id="st"
          title="Statistics"
          panel={<StatPanel {...this.props} />}
          panelClassName="ember-panel"
        />
       <Tab id="cn" disabled title="Configuration" panel={<ConfigPanel />} />
        <Tabs.Expander />
        <input className="bp3-input" type="text" placeholder="Search..." />
      </Tabs>
    );
  }
}

export default NodeStat;

const InfoPanel = (props) => (
  <div>
    <H4>Information </H4>
    {props.SelectedNodes.length > 0 ? (
      <div>
        <table className="bp3-text-muted">
          {/* Node Address */}
          <tbody>
            <tr>
              {/* CPU usage */}
              <td>
                <p className="bp3-text-small">
                  <b>CPU Usage</b>
                </p>
              </td>
              <td className="pull-right2">
                <p className="bp3-text-small">:</p>
              </td>
              <td className="pull-right5">
                <p className="bp3-text-small">35%</p>
              </td>
              {/* -------- */}
              <td className="pull-right5">
                <p className="bp3-text-small">
                  <b>Memory Usage</b>
                </p>
              </td>
              <td className="pull-right2">
                <p className="bp3-text-small">:</p>
              </td>
              <td className="pull-right5">
                <p className="bp3-text-small">45%</p>
              </td>
            </tr>
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
                  {props.SelectedNodes[props.SelectedNodes.length -1].NodeData.Node.IP ===""
                    ? "Configuring ..."
                    : props.SelectedNodes[props.SelectedNodes.length -1].NodeData.Node.IP}
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
                  {props.SelectedNodes[props.SelectedNodes.length -1].NodeData.Node.MAC === ""
                    ? "Configuring ..."
                    : props.SelectedNodes[props.SelectedNodes.length -1].NodeData.Node.MAC}
                </p>
              </td>
            </tr>
            {/* Last Seen */}
            <tr>
              <td>
                <p className="bp3-text-small">
                  <b>Last Seen</b>
                </p>
              </td>
              <td className="pull-right2">
                <p className="bp3-text-small">:</p>
              </td>
              <td className="pull-right5">
                <p className="bp3-text-small">{Date.now()}</p>
              </td>
            </tr>
            {/* Configure as */}
            <tr>
              <td>
                <p className="bp3-text-small">
                  <b>Neighbours</b>
                </p>
              </td>
              <td className="pull-right2">
                <p className="bp3-text-small">:</p>
              </td>
              <td className="pull-right5">
                <p className="bp3-text-small">
                  <b>{props.SelectedNodes[props.SelectedNodes.length -1].NodeData.Neighbours == null
                    ? "(Identifying)"
                    : props.SelectedNodes[props.SelectedNodes.length -1].NodeData.Neighbours.length} &nbsp; Nodes</b>
                </p>
              </td>
            </tr>
            {/* State */}
            <tr>
              <td>
                <p className="bp3-text-small">
                  <b>State</b>
                </p>
              </td>
              <td className="pull-right2">
                <p className="bp3-text-small">:</p>
              </td>
              <td className="pull-right5">
                <p className="bp3-text-small">
                  <Icon icon={"dot"} iconSize={17} intent="success" />
                  <b>Active</b>
                </p>
              </td>
            </tr>

            <tr>
              <td>
                <p className="bp3-text-small">
                  <b>Trafic</b>
                </p>
              </td>
              <td className="pull-right2">
                <p className="bp3-text-small">:</p>
              </td>
              <td className="pull-right5">
                <Trend
                  smooth
                  height={35}
                  width={70}
                  autoDraw
                  // autoDrawDuration={3000}
                  autoDrawEasing="ease-out"
                  data={[0, 2, 5, 9, 5, 10, 3, 5, 0, 0, 1, 8, 2, 9, 0]}
                  gradient={["#222"]}
                  radius={10}
                  strokeWidth={0.5}
                  strokeLinecap={"butt"}
                />
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    ) : (
      <p className="bp3-text-muted bp3-text-small">Please select a Node</p>
    )}
  </div>
);

const StatPanel = (props) => (
  <div>
    <H4>Statistics</H4>

    {props.SelectedNodes.length > 0 ? (
      <div>
        <hr />

        <table className="bp3-text-muted">
          <tbody>
            <tr>
              <td>
                <p className="bp3-text">
                  <b>Network Monitoring</b>
                </p>
              </td>
              <td className="pull-right2">
                <p className="bp3-text-small">:</p>
              </td>
            </tr>
            <tr>
              <td>
                <p className="bp3-text-small">
                  {" "}
                  <b>Drop Count</b>
                </p>
              </td>
              <td className="pull-right2">
                <p className="bp3-text-small">:</p>
              </td>
              <td className="pull-right2">
                <p className="bp3-text-small">20</p>
              </td>
              <td className="pull-right10">
                <p className="bp3-text-small">
                  <b>Forward Count</b>
                </p>
              </td>
              <td className="pull-right2">
                <p className="bp3-text-small">:</p>
              </td>
              <td className="pull-right2">
                <p className="bp3-text-small">30</p>
              </td>
              <td className="pull-right10">
                <p className="bp3-text-small">
                  <b>Local Count</b>
                </p>
              </td>
              <td className="pull-right2">
                <p className="bp3-text-small">:</p>
              </td>
              <td className="pull-right2">
                <p className="bp3-text-small">50</p>
              </td>
            </tr>
            <tr>
              <td>
                <p className="bp3-text-small">
                  {" "}
                  <b>InNet Trafic </b>
                </p>
              </td>
              <td className="pull-right2">
                <p className="bp3-text-small">:</p>
              </td>
              <td className="pull-right2">
                <p className="bp3-text-small">
                  50 kbps
                  <Trend
                    smooth
                    height={35}
                    width={70}
                    autoDraw
                    // autoDrawDuration={3000}
                    autoDrawEasing="ease-out"
                    data={[5, 6, 7, 8, 5, 9, 3, 5, 5, 6, 1, 9, 5, 2, 9]}
                    gradient={["#0D8050"]}
                    radius={10}
                    strokeWidth={0.5}
                    strokeLinecap={"butt"}
                  />
                </p>
              </td>
              <td className="pull-right10">
                <p className="bp3-text-small">
                  <b>OutNet Trafic</b>
                </p>
              </td>
              <td className="pull-right2">
                <p className="bp3-text-small">:</p>
              </td>
              <td className="pull-right2">
                <p className="bp3-text-small">
                  0.8 Mbps
                  <Trend
                    smooth
                    height={35}
                    width={70}
                    autoDraw
                    // autoDrawDuration={3000}
                    autoDrawEasing="ease-out"
                    data={[0, 2, 5, 9, 5, 10, 3, 5, 0, 0, 1, 8, 2, 9, 0]}
                    gradient={["#C23030"]}
                    radius={10}
                    strokeWidth={0.5}
                    strokeLinecap={"butt"}
                  />
                </p>
              </td>
            </tr>
            <tr>
              <td>
                <p className="bp3-text-small">
                  {" "}
                  <b>Reliability </b>
                </p>
              </td>
              <td className="pull-right2">
                <p className="bp3-text-small">:</p>
              </td>
              <td className="pull-right2">
                <p className="bp3-text-small">
                  80%
                  <Trend
                    smooth
                    height={35}
                    width={70}
                    autoDraw
                    // autoDrawDuration={3000}
                    autoDrawEasing="ease-out"
                    data={[9, 7, 6, 4, 8, 9, 5, 3, 6, 7, 9, 11, 6, 3, 9]}
                    gradient={["#106BA3"]}
                    radius={10}
                    strokeWidth={0.5}
                    strokeLinecap={"butt"}
                  />
                </p>
              </td>
            </tr>
          </tbody>
        </table>
        <hr />
        <table className="bp3-text-muted">
          <tbody>
            <tr>
              <td>
                <p className="bp3-text">
                  <b>Rule Usage Monitoring</b>
                </p>
              </td>
              <td className="pull-right2">
                <p className="bp3-text-small">:</p>
              </td>
            </tr>

            <tr>
              <th>
                <p className="bp3-text-small">
                  <b>Rule</b>
                </p>
              </th>
              <th className="pull-right10">
                <p className="bp3-text-small">
                  <b>Hit Count</b>
                </p>
              </th>
              <th className="pull-right10">
                <p className="bp3-text-small">
                  <b>State</b>
                </p>
              </th>
              <th className="pull-right10">
                <p className="bp3-text-small">
                  <b>Transition Time</b>
                </p>
              </th>
            </tr>
            <tr>
              <td>
                <p className="bp3-text-small"># Rule 00123</p>
              </td>
              <td className="pull-right10">
                <p className="bp3-text-small">
                  <b>5</b>
                  <Trend
                    smooth
                    height={35}
                    width={70}
                    autoDraw
                    // autoDrawDuration={3000}
                    autoDrawEasing="ease-out"
                    data={[0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 1, 1, 1, 0]}
                    gradient={["#106BA3"]}
                    radius={10}
                    strokeWidth={0.5}
                    strokeLinecap={"butt"}
                  />
                </p>
              </td>
              <td>
                <p className="bp3-text-small">
                  <Icon icon={"dot"} iconSize={17} intent="success" />
                  <b>Active</b>
                </p>
              </td>
            </tr>
            <tr>
              <td>
                <p className="bp3-text-small"># Rule 00124</p>
              </td>
              <td className="pull-right10">
                <p className="bp3-text-small">
                  <b>9</b>
                  <Trend
                    smooth
                    height={35}
                    width={70}
                    autoDraw
                    // autoDrawDuration={3000}
                    autoDrawEasing="ease-out"
                    data={[0, 1, 0, 0, 0, 0, 0, 2, 0, 0, 0, 4, 1, 1, 0]}
                    gradient={["#752F75"]}
                    radius={10}
                    strokeWidth={0.5}
                    strokeLinecap={"butt"}
                  />
                </p>
              </td>
              <td>
                <p className="bp3-text-small">
                  <Icon icon={"dot"} iconSize={17} intent="danger" />
                  <b>Inactive</b>
                </p>
              </td>
            </tr>
            <tr>
              <td>
                <p className="bp3-text-small"># Rule 00130</p>
              </td>
              <td className="pull-right10">
                <p className="bp3-text-small">
                  <b>0</b>
                  <Trend
                    smooth
                    height={35}
                    width={70}
                    autoDraw
                    // autoDrawDuration={3000}
                    autoDrawEasing="ease-out"
                    data={[0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0]}
                    gradient={["#752F75"]}
                    radius={10}
                    strokeWidth={0.5}
                    strokeLinecap={"butt"}
                  />
                </p>
              </td>
              <td>
                <p className="bp3-text-small">
                  <Icon icon={"dot"} iconSize={17} intent="success" />
                  <b>Active</b>
                </p>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    ) : (
      <p className="bp3-text-muted bp3-text-small">Please select a Node</p>
    )}
  </div>
);

const ConfigPanel = () => (
    <div>
        <H3>Backbone</H3>
    </div>
);
