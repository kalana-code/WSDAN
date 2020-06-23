

import Home from "views/Home.jsx";
import RuleManager from "views/RuleManager.jsx";
import Network from "views/Network.jsx";
import FlowManager from "views/FlowManager.jsx"
import Settings from "views/Settings.jsx";

export const adminRoute = [
  {
    path: "/dashboard",
    name: "Home",
    icon: "pe-7s-photo-gallery",
    component: Home,
    layout: "/admin"
   },
   {
    path: "/network",
    name: "Network Topology",
    icon: "pe-7s-photo-gallery",
    component: Network,
    layout: "/admin"
   },
   {
    path: "/ruleManager",
    name: "Rule Manager",
    icon: "pe-7s-photo-gallery",
    component: RuleManager,
    layout: "/admin"
   },

   {
    path: "/flowManager",
    name: "Flow Manager",
    icon: "pe-7s-photo-gallery",
    component: FlowManager,
    layout: "/admin"
   },
   
   {
    path: "/settings",
    name: "Settings",
    icon: "pe-7s-photo-gallery",
    component: Settings,
    layout: "/admin"
   }

  ]
export default {adminRoute} ;
