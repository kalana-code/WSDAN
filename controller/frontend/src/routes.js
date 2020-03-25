

import Question from "views/Question.jsx";
import Network from "views/Network.jsx";

export const adminRoute = [
  {
    path: "/dashboard",
    name: "Home",
    icon: "pe-7s-photo-gallery",
    component: Question,
    layout: "/admin"
   },
   {
    path: "/icon",
    name: "Rule Setter",
    icon: "pe-7s-photo-gallery",
    component: Question,
    layout: "/admin"
   },
   {
    path: "/network",
    name: "Network Topology",
    icon: "pe-7s-photo-gallery",
    component: Network,
    layout: "/admin"
   }

  ]
export default {adminRoute} ;
