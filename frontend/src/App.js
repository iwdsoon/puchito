import { Router, Route, Switch } from "react-router-dom";
import history from "./utils/history";
import Generos from "./views/generos";

function App() {
  return (
    <div id="App" className="d-flex flex-column h-100">
      <Router history={history}>
        <Switch>
        <Route exact path="/generos" component={Generos} />
        </Switch>
      </Router>
    </div>
  );
}

export default App;
