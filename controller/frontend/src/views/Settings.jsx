/** @format */

import React, { Component } from "react";

import { Card, H5, FormGroup, Switch } from "@blueprintjs/core";

import { Grid, Row,Col } from "react-bootstrap";

class Settings extends Component {
  render() {
    return (
      <div className="content">
        <Grid fluid>
          <Row>
            <Col sm="6">
              <Card >
                <H5>System Settings</H5>
                {/* <p>Message For User</p> */}
                <FormGroup
                  // helperText="Helper text with details..."
                  label="Control Rules "
                  // labelInfo="(required)"
                >
                   <Switch labelElement={<strong>Enable automation rule generation process</strong>} />
                   <Switch  labelElement={<strong>Enable force rule dispursing</strong>} />
                   {/* <Switch  labelElement={<strong></strong>} /> */}
                </FormGroup>

                <FormGroup
                  // helperText="Helper text with details..."
                  label="Other Settings"
                  // labelInfo="(required)"
                >
                   <Switch labelElement={<strong>Setting 1</strong>} />
                   <Switch  labelElement={<strong>Setting 2</strong>} />
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
