import React, { Component } from "react";

import { Button } from "@blueprintjs/core";

import { Grid, Row } from "react-bootstrap";

class Dashboard extends Component {
  render() {
    return (
      <div className="content">
        <Grid fluid>
          <Row>
            
            <Button intent="success" text="button content" />
          </Row>
        </Grid>
      </div>
    );
  }
}

export default Dashboard;
