import React, { Component } from 'react';

// import AdminLayout from "layouts/Admin.jsx";
import  LoginLayOut  from "layouts/Login.jsx";
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
                <ProtectedRoute allowRoles={['ADMIN']} exact path="/admin/*" component={User} redirectPath="/login"  /> */}
                {/* { <ProtectedRoute allowRoles={['ADMIN']} exact path="/" component={User} redirectPath="/login"/> } */}
                <Route path="/login"  render={props => <LoginLayOut {...props}/>} />
                { auth.getRole() === "ADMIN" ?<Redirect to='/admin/dashboard'  />:<Redirect to='/login'  />}  }
            </Switch>
          </BrowserRouter> );
    }
}
 
export default RouteManager;