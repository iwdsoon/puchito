import { Router, Route, Switch } from "react-router-dom";
import history from "./utils/history";
import Layout from "./views/home/layout";
import { Home } from "./views/home/home";
import Generos from "./views/generos";
import { Container } from "reactstrap";

function App() {
  return (
    <div id="App" className="d-flex h-100">
      <Router history={history}>
      <Container fluid className="flex-grow-1" style={{ paddingLeft: "0px", paddingRight: "0px" }}>
        <Switch>
        <Layout>
        <Route exact path="/home" component={Home} />
        <Route exact path="/home/generos" component={Generos} />
        </Layout>
        </Switch>
        </Container>
      </Router>
    </div>
  );
}

export default App;
