import React, { Component } from "react";
import { Link } from "react-router-dom";
import { HashRouter as Router, Route } from "react-router-dom";
import SignUp from "./components/SignUp";
import "./App.css";
import SignIn from "./components/SignIn";

class App extends Component {
  constructor(props) {
    super(props);
    this.state = { isLoggedIn: false };
  }
  render() {
    return (
      <Router className="base">
        <nav className="nav">
          <div className="logo">
            <Link to="/" className="nav-edit">
              Analytics Dashboard
            </Link>
          </div>
        </nav>

        <div className="app">
          <div className="left" />
          <div className="app-form">
            <Route exact path="/" component={SignUp} />
            <Route exact path="/sign-in" component={SignIn} />
          </div>
        </div>
      </Router>
    );
  }
}

export default App;
