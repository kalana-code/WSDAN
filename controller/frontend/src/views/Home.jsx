/** @format */

import "./style/home.css";

import React from "react";
import { Grid, Row, Col } from "react-bootstrap";
import { Callout } from "@blueprintjs/core";

class Home extends React.Component {
  render() {
    return (
      <div className="content">
        <Grid fluid>
          <Row>
            <Col lg="6" sm="12">
              <Callout title="About" icon="info-sign" intent="primary">
                <div className="bp3-ui-text bp3-text-small">
                <p className="bp3-text-small" style={{ "textAlign": "justify" ,'paddingRight':"10px" }}>
                    Software Defined Network <code>(SDN)</code> is the modern
                    trend of networking towards implementing flexible and
                    programmable networks. Software defined networks decouples
                    the control plane from the data plane allowing them to
                    operate as two separate entities. Here this concept has been
                    extended to an ad-hoc networks. An ad-hoc network is a
                    network that is composed of individual devices communicating
                    with each other directly over wired or wireless interfaces.
                  </p>
                </div>
              </Callout>
            </Col>

            <Col lg="6" sm="12">
              <Callout title="What WSDN Can Do ?" icon="info-sign" intent="success">
                <div className="bp3-ui-text ">
                  <p className="bp3-text-small" style={{ "textAlign": "justify" ,'paddingRight':"10px" }}>
                    Due to all these characteristics WSDAN can be used to
                    overcome lots of practical problems that are arised in day
                    today life. Integrating WSDAN into traffic lighting system,
                    security robots, drone systems used in disaster situations
                    etc are some of the practical situations where WSDAN is or
                    can be used.
                    <br />
                    <br />

                  </p>
                </div>
              </Callout>
            </Col>
            <Col lg="12" sm="12">
              <div className="center">
                {/* <img
                  src="https://banner2.cleanpng.com/20180808/zfs/kisspng-systems-architecture-project-management-software-a-5b6a8f834d7ca6.7991397715337102113174.jpg"
                  alt="Girl in a jacket"
                /> */}
                <img src={ require('./img/Home.png') } />
              </div>
              s
            </Col>
            <Col lg="6" sm="12">
              <Callout title="Note" icon="info-sign" intent="success">
                Message About Project
              </Callout>
            </Col>
          </Row>
        </Grid>
      </div>
    );
  }
}

export default Home;
