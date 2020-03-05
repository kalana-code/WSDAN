import React, { Component } from 'react';

// import AdminLayout from "layouts/Admin.jsx";
import  LoginLayOut  from "layouts/Login.jsx";
import RegisterLayOut from "layouts/Register.jsx";
import User from "layouts/User.jsx"
import { ProtectedRoute } from  "auth/privateRoute"
import { BrowserRouter, Route, Switch, Redirect } from "react-router-dom";
import auth from './auth'




class RouteManager extends Component {
    state = {  }
    render() { 
        return ( <BrowserRouter>
            <Switch>
                {/* <Route path="/Register" render={props => <RegisterLayOut {...props}/>} /> */}
                <Route path="/login"  render={props => <LoginLayOut {...props}/>} />
                <ProtectedRoute allowRoles={['STUDENT', 'ADMIN']} exact path="/user/*" component={User} redirectPath="/login"  />
                {/* <ProtectedRoute allowRoles={['ADMIN']} exact path="/" component={User} redirectPath="/login"/> */}
                { auth.getRole() =="STUDENT" ? <Redirect to='/user/dashboard'  />:
                    ( auth.getRole() =="ADMIN" ?<Redirect to='/admin/dashboard'  />:<Redirect to='/login'  />)}   
            </Switch>
          </BrowserRouter> );
    }
}
 
export default RouteManager;