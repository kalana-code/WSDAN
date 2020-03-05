import React, { Component } from "react";
import { NavLink } from "react-router-dom";

class Sidebar extends Component {
  // constructor(props) {
  //   super(props);
  //   this.state = {
  //     width: window.innerWidth
  //   };
  // }
  activeRoute(routeName) {
    return this.props.location.pathname.indexOf(routeName) > -1 ? "active" : "";
  }
  // updateDimensions() {
  //   this.setState({ width: window.innerWidth });
  // }
  // componentDidMount() {
  //   this.updateDimensions();
  //   window.addEventListener("resize", this.updateDimensions.bind(this));
  // }
  render() {
    
    return (
      <div className="side-bar-left">                  
          <ul>
          {this.props.routes.map((prop, key) => {
              if (!prop.redirect)
              console.log(this.activeRoute(prop.layout + prop.path))
                return (
                    <NavLink
                      key={key}
                      to={prop.layout + prop.path}
                      className="user-nav-link"
                      activeClassName={this.activeRoute(prop.layout + prop.path)}
                    >
                     <i className={prop.icon}></i>
                    </NavLink> 
              )})}
          </ul>
      </div>
      );
  }
}

export default Sidebar;
