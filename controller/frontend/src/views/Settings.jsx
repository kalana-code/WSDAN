/** @format */

import axios from "axios";

import React, { Component } from "react";

import { Card, H5, FormGroup, Switch } from "@blueprintjs/core";

import { Grid, Row, Col } from "react-bootstrap";
import config from "./../config/config"

class Settings extends Component {
  state = {
    automationLoad: false,
    automation: false,

    dispurserMode:false,
    dispurserLoad:false
  };
  togleAutomation(event) {
    this.setState({
      automationLoad: true,
    });
    axios.get(`http://`+config.host+`:8081/StateToggle`).then((res) => {
      if (res.status === 200) {
        this.setState({
          automation: !this.state.automation,
        });
      }
    });
    this.setState({
      automationLoad: false,
    });
  }

  togleStateToggleForceDispurser(event) {
    this.setState({
      dispurserLoad: true,
    });
    axios.get(`http://`+config.host+`:8081/StateToggleForceDispurser`).then((res) => {
      if (res.status === 200) {
        this.setState({
          dispurserMode: !this.state.dispurserMode,
        });
      }
    });
    this.setState({
      dispurserLoad: false,
    });
  }

  
  render() {
    return (
      <div className="content">
        <Grid fluid>
          <Row>
            <Col sm={6}>
              <Card>
                <H5>System Settings</H5>
                {/* <p>Message For User</p> */}
                <FormGroup
                  // helperText="Helper text with details..."
                  label="Control Rules "
                  // labelInfo="(required)"
                >
                  <Switch
                    checked={this.state.automation}
                    disabled={this.state.automationLoad}
                    onChange={() => this.togleAutomation()}
                    labelElement={
                      <strong>Enable automation rule generation process</strong>
                    }
                  />
                  <Switch 
                  checked={this.state.dispurserMode}
                  disabled={this.state.dispurserLoad}
                  onChange={() => this.togleStateToggleForceDispurser()}
                    labelElement={<strong>Enable force rule dispursing</strong>}
                  />
                </FormGroup>

                <FormGroup
                  // helperText="Helper text with details..."
                  label="Other Settings"
                  // labelInfo="(required)"
                >
                  <Switch
                    disabled={true}
                    labelElement={<strong>Setting 1</strong>}
                  />
                  <Switch
                    disabled={true}
                    labelElement={<strong>Setting 2</strong>}
                  />
                  {/* <Switch  labelElement={<strong></strong>} /> */}
                </FormGroup>
              </Card>
            </Col>
          </Row>
        </Grid>
      </div>
    );
  }
}

export default Settings;
