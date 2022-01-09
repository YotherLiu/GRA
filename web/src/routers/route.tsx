import { BrowserRouter as Router, Switch, Route } from "react-router-dom";
import { router } from "./router"

// 自定义路由组件
function BasicRoute() {
  return (
    <Router>
      <Switch>
        {
          router.map((val, index) => {
            return <Route
              key={index}
              path={val.path}
              component={val.component}
            />
          })
        }
      </Switch>
    </Router>
  );
}
export default BasicRoute
