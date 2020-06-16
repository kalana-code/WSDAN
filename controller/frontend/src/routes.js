

import Home from "views/Home.jsx";
import Question from "views/Question.jsx";
import Network from "views/Network.jsx";
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
    path: "/icon",
    name: "Rule Manager",
    icon: "pe-7s-photo-gallery",
    component: Question,
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
