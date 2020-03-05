
import React from "react";
import ReactDOM from "react-dom";

import "bootstrap/dist/css/bootstrap.min.css";
import "./assets/css/animate.min.css";
import "./assets/sass/light-bootstrap-dashboard-react.scss?v=1.3.0";
import "./assets/css/demo.css";
import "./assets/css/pe-icon-7-stroke.css";

// Blue print Styles
import "@blueprintjs/core/lib/css/blueprint.css";


import RouteManager from "./auth/RouteManager"



ReactDOM.render(<RouteManager/>,
  document.getElementById("root")
);
