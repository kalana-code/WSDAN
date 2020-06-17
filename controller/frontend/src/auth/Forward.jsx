import React, { Component } from 'react';
import axios from 'axios';
import "./style/style.css"
import config from "./../config/config"

class Forward extends Component {
    state = { IsLoading:true }

    componentDidMount=()=>{
        if(localStorage.getItem('Token')){
            console.log(localStorage.getItem('Token'))
        }
        // const config = {
        //     method: "get",
        //     url: "https://`+config.host+`:8081/auth/verify",
        //     headers: { }
        // }
        axios.get(`http://`+config.host+`:8081/auth/verify`, {
   
        headers: {
            'x-access-token':localStorage.getItem('Token')
        }
        }).then(response=>{
            if(response.status===200){
                console.log("OK")
                setTimeout( this.setState({IsLoading:false}), 2000);
                
            }else{
                console.log("Error")
                localStorage.removeItem("Token")
                this.props.history.push("/login");
            }
        }).catch(error=>{
            console.log("Error")
            localStorage.removeItem("Token")
            this.props.history.push(this.props.redirectPath);
        })
    }
    render() { 
        return ( 
            <div>
             {this.state.IsLoading?
             <div className="forward-Loading"> <div className="lds-ripple"><div></div><div></div><div></div><div></div></div></div>
             :
            <this.props.component {...this.props}/>}
            </div>
         );
    }
}
 
export default Forward;

Forward.propType ={
    component: Component,
    redirectPath:String

}