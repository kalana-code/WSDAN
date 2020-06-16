
import './style/home.css'

import React, { Component } from "react";
import { Grid, Row, Col } from "react-bootstrap";
import { Callout, } from "@blueprintjs/core";

class Home extends Component {
  render() {
    return (
      <div className="content">
        <Grid fluid>
          <Row>
            <Col lg="6" sm="12">
              <Callout title={"Note"} icon={"info-sign"} intent={"primary"}>
              Message About Project .
              </Callout>
            </Col>

            <Col lg="6" sm="12">
              <Callout title={"Note"} icon={"info-sign"} intent={"success"}>
                Message About Project .
              </Callout>
            </Col>
            <Col lg="12" sm="12">
              <div className="center">
              <img src="https://ak.picdn.net/shutterstock/videos/11811779/thumb/1.jpg" alt="Girl in a jacket"/>
              </div>
              
            </Col>
            <Col lg="6" sm="12">
              <Callout title={"Note"} icon={"info-sign"} intent={"success"}>
                Message About Project .
              </Callout>
            </Col>
          </Row>
        </Grid>
      </div>
    );
  }
}

export default Home;
