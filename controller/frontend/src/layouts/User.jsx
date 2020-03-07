import React, { Component } from 'react';
import "./style/user.css";
import { Route, Switch } from "react-router-dom";
import {Popover,Button,ButtonGroup,Classes,Divider,Icon} from '@blueprintjs/core';
import "@blueprintjs/icons"
import logo from "./assets/logo/logo_small.png"

import auth from '../auth/auth'
import Sidebar from "./../components/Sidebar/SidebarUser"

import {adminRoute} from "routes.js";

class User extends Component {
    state = { 
        showPopover: false,
        IsLoading:true   
    }
    getRoutes = routes => {
        return routes.map((prop, key) => {
          if (prop.layout === "/admin") {
            return (
              <Route
                path={prop.layout + prop.path}
                render={props => (
                  <prop.component
                    {...props}
                    handleClick={this.handleNotificationClick}
                  />
                )}
                key={key}
              />
            );
          } else {
            return null;
          }
        });
    };
    
    getBrandText = path => {
        for (let i = 0; i < adminRoute.length; i++) {
          if (
            this.props.location.pathname.indexOf(
                adminRoute[i].layout + adminRoute[i].path
            ) !== -1
          ) {
            return adminRoute[i].name;
          }
        }
        // this.props.history.push("/admin/dashboard");
        return "Brand";
    };
      
    handleInteraction=(nextOpenState)=> {
        this.setState({ showPopover: nextOpenState });
    }
    
    render() { 
        
       
        const user =<Icon icon={"user"} iconSize={Icon.SIZE_LARGE}/>
        console.log(this.props.location.pathname)
         return ( 
            <div className="main-container">
                <nav>
                    <img className="logo" alt="logo" src={logo} />
                    <div className="search-bar">
                        {/* <InputGroup leftIcon="search" style={{background:"#3C4144"}}/> */}
                    </div>
                    <div className="button-set">
                        <Popover
                            interactionKind={"hover"}
                            popoverClassName={Classes.POPOVER_CONTENT_SIZING }
                            isOpen={this.state.showPopover}
                            onInteraction={(state) => this.handleInteraction(state)}
                            >
                            <Button icon={user}    minimal="true"></Button>
                            <div className="user-details">
                                <div className="user-pic">
                                    <img alt ="profile" src={auth.getProfile()} />
                                    <Button className="user-pic-add" name="showPopover" icon="plus"/>
                                </div>
                                <p><b>Hi!.</b> Kalana . Welcome to the Beq</p>
                                <div style={{ display: "flex", justifyContent: "space-evenly", marginTop: 15 }}>
                                <ButtonGroup minimal={true} fill={true} large={true} >
                                    <Button icon="log-out" onClick={()=>{auth.logOut(
                                        ()=>{
                                            this.props.history.push("/login");
                                        }
                                    )}}/>
                                    <Divider/>
                                    <Button icon="list-detail-view"/>
                                    <Divider/>
                                    <Button icon="stacked-chart"/>
                                </ButtonGroup>
                                </div>
                            </div>
                        </Popover>
                    </div>  
                     
                </nav>
                <div className="container">
                   <Sidebar routes={adminRoute} {...this.props}/>
                    <div className="main">
                        <div className="Tag">
                            <p className="Tag-Name">{this.getBrandText(this.props.location.pathname)}</p>
                        </div>
                        <div  className="user-main-panel" >
                            <Switch>{this.getRoutes(adminRoute)}</Switch>
                        </div>
                    </div>
                    
                </div>
                <div className="footer">
                        
                </div>
                
         </div> )
    }
}
 
export default User;