
import Icons from "views/admin/Icons.jsx";
import Question from "views/user/Question.jsx";

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
    name: "Icon",
    icon: "pe-7s-photo-gallery",
    component: Icons,
    layout: "/admin"
   }

  ]
export default {adminRoute} ;
