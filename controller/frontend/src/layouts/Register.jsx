import React, { Component } from 'react';
import {FormGroup,InputGroup,Button,RadioGroup, Radio,Tooltip,Intent} from '@blueprintjs/core';
// import { DateInput } from "@blueprintjs/datetime";
import axios from "axios";

import '@blueprintjs/core/lib/css/blueprint.css';
import "./style/register.css";
import { DateInput } from "@blueprintjs/datetime";

const emailRegex = /^(([^<>()[\]\\.,;:\s@\"]+(\.[^<>()[\]\\.,;:\s@\"]+)*)|(\".+\"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;




class RegisterLayOut extends Component {
    // see java class
  constructor(props){
    super(props);
  }

  state = {
    // all variables except gender are object type. because they are initialized to null.
    firstName : null,
    lastName : null,
    email : null,
    password : null,
    gender:"MALE",


    formErrors : {
      // No error for gender. because it has defult value
      firstName : "",
      lastName : "",
      email : "",
      password : "",
        
    }
  }

  // form validation function.
  // Passing parameters is the state of the class. formErrors is one parameter.
  // other variables of the state are taken as one parameter. (...variable_name)
  // when we passing objects we need to use curly brackets.
  formValid = ({formErrors,...rest})=>{
    let valid = true;
    // Reading variables of a object. If any variable of formErrors object have string (error) set valid as false
    Object.values(formErrors).forEach(val => {
      (val.length>0 ) &&(valid = false);
    });
  
    // checking the objects of the state except formErrors , whether they are null
    if(valid!==false){
      // rest is a object containg objects.
      Object.values(rest).forEach(val =>{
        if(val===null){
          valid = false;
        }
      }) 
    }

    return valid;
  };


  // When inputing datas to the form this function will be called. (like typing)
  // "e" means event. (It is a class.It has a object called target. 
  // In target object there are 2 variables called key and value. 
  // key = name of the element. name of the element should same as state variable.
  // value = the data that we input. 
  // When inputing datas , event is passing to the function.
  handleChange = (e)=>{
    
    // declare a new object and reference it to e.target object
    const {name , value} = e.target;
    console.log(name,value)

    // declare a new object and reference it to formErrors object
    let formErrors = this.state.formErrors;

  
    switch(name){
      // if name of the passing event is "firstName" , then run this.
      case "firstName" :
        formErrors.firstName = value.length < 3 
          ? "minimum 3 chars requires" : "";
        break;
  
      case "lastName" :
        formErrors.lastName = value.length < 3 
          ? "minimum 3 chars requires" : "";
        break;
  
      case "email" :
        // emailRegex.test(value) : verifying the entered email
        formErrors.email = value.length > 3 && emailRegex.test(value)
          ? "" : "Invalid email address";
        break;
  
      case "password" :
          formErrors.password = value.length < 6 
            ? "minimum 6 chars requires" : "";
          break;

        break;
    }

    // this.setState({formErrors})    :  when the state name and passed variable name is same , we can passed like this.
    // this.setState([key] : value)   :  when passing key and value
    //                                            call back function
    this.setState({formErrors , [name]: value} , () => console.log(this.state));
    
  }

// this is a another way of handling events
// -----------------------------------------
  // handleGenderChange=(e)=>{
  //   this.setState({gender:e.target.value} )
  // }
  

// when submitting the form this func will be called. 
  handleSubmit =(e)=>{
    // disable the pre buit funcs.
    e.preventDefault();

    // when formValid is true
    if (this.formValid(this.state)){
      // A way of printing in console
      // "$" is used to access variables
      console.log(`
                  --Submitting--
        FirstName : ${this.state.firstName}
        LastName : ${this.state.lastName}
        Email : ${this.state.email}
        Password : ${this.state.password}
        Gender : ${this.state.gender}
      `)
    }
    else{
      // if formValid is false give error
      console.error("INVALID SUBMIT")
    }
  }

  

  render() { 
    return ( 
      <div className = "register">
          <div className="register-form">
                <div className="register-inputs">
                    <h4 class="bp3-heading">Register</h4>
                    <FormGroup
                    intent = "danger"
                    label="First Name"
                    labelFor="firstName"
                    // form errors regarding to first name will be displayed
                    helperText = {this.state.formErrors.firstName}
                    >
                    <InputGroup 
                    // if form errors regarding to first name exists , intent will be danger
                    intent = {this.state.formErrors.firstName.length > 0 ? "danger" : "none"}
                    name="firstName" 
                    placeholder="Sachintha" 
                    // when typing handleChange func will be run
                    onChange = {this.handleChange}/>
                    </FormGroup>
                    
                    

                    <FormGroup
                    intent = "danger"
                    label="Last Name"
                    labelFor="lastName"
                    helperText = {this.state.formErrors.lastName}
                    >
                    <InputGroup 
                    intent = {this.state.formErrors.lastName.length > 0 ? "danger" : "none"}
                    name="lastName" 
                    placeholder="Gunathilaka" 
                    onChange = {this.handleChange}
                    />
                    </FormGroup>


                    <FormGroup
                    intent = "danger"
                    label="Email"
                    labelFor="email"
                    helperText = {this.state.formErrors.email}
                    >
                    <InputGroup 
                    intent = {this.state.formErrors.email.length > 0 ? "danger" : "none"}
                    name="email" 
                    placeholder="sachi.lifef@gmail.com" 
                    onChange = {this.handleChange}
                    />
                    </FormGroup>


                    <FormGroup
                    intent = "danger"
                    label="Password"
                    labelFor="password"
                    helperText = {this.state.formErrors.password}
                    >
                    <InputGroup 
                    intent = {this.state.formErrors.password.length > 0 ? "danger" : "none"}
                    name="password" 
                    placeholder="" 
                    onChange = {this.handleChange}
                    />
                    </FormGroup>


                    <RadioGroup
                        inline = "true"
                        label="Gender"
                        name="gender"
                        onChange={this.handleChange}
                        selectedValue={this.state.gender}
                    >
                        <Radio label="Male" value="MALE" />
                        <Radio label="Female" value="FEMALE" />
                        <Radio label="Other" value="OTHER" />
                    </RadioGroup>

                    Birth Day 
                    <br/><br/>

                    <Button 
                    fill = "true"
                    className = "submit-button"
                    text = "Register" 
                    type = "submit" 
                    onClick = {this.handleSubmit}
                    />

                </div>
            </div>
      </div>
      
     );
  }
}
 
export default RegisterLayOut ;