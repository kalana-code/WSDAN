import React from "react";
import { Route, Redirect } from "react-router-dom";
import auth from "./auth";
import Forward from './Forward'

export const ProtectedRoute = ({
  component: Component,
  allowRoles,
  ...rest
}) => {
  
  return (
    <Route
      {...rest}
      render={props => {
        if (auth.isAuthenticated(allowRoles) ) {
          return <Forward component={Component} {...props} />;
        } else {
          return (
            <Redirect
              to={{
                pathname: "/login",
                state: {
                  from: props.location
                }
              }}
            />
          );
        }
      }}
    />
  );
};
