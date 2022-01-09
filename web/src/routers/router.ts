import Index from "../views/index/index";
import Home from "../views/home/home";
// Route parse
let router: any[] = [
  {
    path: "/",
    component: Index,
    auth: true,
  },
  {
    path: "/home",
    component: Home,
  },
];

export { router };
