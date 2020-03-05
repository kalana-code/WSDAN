import React, { Component } from 'react';
import "./style/login.css";
import {FormGroup,InputGroup,Button,Checkbox,Intent,Tooltip} from '@blueprintjs/core';
import axios from "axios";
import auth from "./../auth/auth"
import logo from "./assets/logo/logo.png"

import Regex from "./util/regax"

export default class LoginLayOut extends Component {
    
    state = { 
        isLoading: false,
        UserName:"",
        Password:"",
        showPassword:false,
        Error:{
            UserName:false,
            Password:false
        },
        ErrorMessage:{
            UserName:"",
            Password:""
        }
    }
    // handle Functions 
    handleChange=(event)=>{
        const {name,value} =event.target
        this.setState({[name]:value})
        this.Verify()
    }

    handleLock=()=>{
        this.setState({showPassword:!this.state.showPassword})
    }

    Verify=()=>{
        
        let Error={
            UserName:false,
            Password:false
        }
        let ErrorMessage={
            UserName:"",
            Password:""
        }
        //Check Email 
        if(this.state.UserName ===""){
            Error.UserName = true
            ErrorMessage.UserName="Email cannot be empty"
        } else if(! Regex.Email.test(this.state.UserName)){
            Error.UserName = true
            ErrorMessage.UserName="Not valid Email"
        }

        //Check Password
        if(this.state.Password ===""){
            Error.Password = true
            ErrorMessage.Password="Password cannot be empty"
        } else if(this.state.Password.length<8){
            Error.Password = true
            ErrorMessage.Password="Pasword should have at least 8 character"
        }
        // set State
        this.setState({Error:Error,ErrorMessage:ErrorMessage})
    }

    // get intent 
    getIntent=(feild)=>{
        if(this.state.Error[feild]){
            return "danger"
        }
        return "primary"
    }

    //Send Request
    DataSubmit=()=> {
        let isValid= true;
        //check form input errors
        Object.keys(this.state.Error).map((value)=>{
            if(this.state.Error[value]){
                isValid = false 
            }
        })

        // Submit Data
        if(isValid) {
            const Request_Body =
                {
                    "Email":this.state.UserName,
                    "Password":this.state.Password
                }
            this.setState({isLoading:true})
            axios.post(`http://localhost:8081/Student/Login`, Request_Body).then(response => {
                if(response.status === 200){
                    localStorage.setItem('Token', 
                        response.data.Data.token);
                    this.setState({isLoading:false})
                    this.props.history.push("/user");

                    
            }
            },
            error=>{
                console.log(error);
                console.log(error.data);
                this.setState({isLoading:false}) 
                });
            }

        
        
      }
    componentWillMount(){
        if(auth.isAuthenticated(this.props.allowedRoles)){
            this.props.history.push("/user/dashboard");
        }
    }


    render() { 
        const lockButton = (
            <Tooltip content={`${this.state.showPassword ? "Hide" : "Show"} Password`}>
                <Button
                    // disabled={disabled}
                    icon={this.state.showPassword ? "eye-open" : "eye-off"}
                    intent={Intent.WARNING}
                    minimal={true}
                    onClick={this.handleLock}
                />
            </Tooltip>
        );

            return (
                <div className="sign-in">
                    <div className="form1">
                        <div className="logo">
                            <img width="100px"  />
                        </div>
                        <div className="inputs">
                            
                            {/* <p className="login-head">LOGIN</p> */}
                            <FormGroup
                                helperText={this.state.ErrorMessage.UserName}
                                label="User Name"
                                // labelFor="text-input"
                            >
                                <InputGroup id="text-input" name="UserName" placeholder="User Name" intent={this.getIntent("UserName")} value={this.state.UserName}  onChange={this.handleChange} />
                            </FormGroup>
                            <FormGroup
                                helperText={this.state.ErrorMessage.Password}
                                label="Password"
                                // labelFor="text-input"
                            >
                                <InputGroup id="text-input" name="Password"  intent={this.getIntent("Password")} placeholder="Password" value={this.state.Password} type={this.state.showPassword ? "text":"password"} onChange={this.handleChange} rightElement={lockButton} />
                            </FormGroup>
                            <FormGroup>
                                <Checkbox checked={true} label="Remember Me"  />
                            </FormGroup>

                            <FormGroup>
                                <Button intent="primary" text="Log In" fill={true} onClick={this.DataSubmit}  loading={this.state.isLoading} />
                            </FormGroup>
                            
                        </div>
                        <div className="linkContainer">
                                <div className="link">
                                    
                                    {/* <a  href="/Student/Register">Register</a> */}
                                </div>
                                <div className="link">
                                    {/* <a href="/Register">Forget Your Password</a> */}
                                </div>
                                
                                
                        </div>
                        
                    </div>
                   
                    <div className="footer"> All copyright reserve , 2020 Beq  </div>
                </div>
            );
    }
}