import React, { Component } from "react";
import { HashRouter as Router, Route } from "react-router-dom";
import SignUp from "./components/SignUp";
import "./App.css";
import SignIn from "./components/SignIn";

class App extends Component {
  constructor(props) {
    super(props);
    this.state = {
      user: "Guest"
    };
  }
  user(data) {
    this.setState({ user: data });
  }
  render() {
    return (
      <Router className="base">
        <nav className="nav">
          <div>Analytics Dashboard</div>
          <div> {this.state.user} </div>{" "}
        </nav>

        <div className="app">
          <div className="left" />
          <div className="app-form">
            <h2 className="form-title">Sign Up/Sign In</h2>
            <Route exact path="/" component={SignUp} />
            <Route exact path="/sign-in" component={SignIn} />
          </div>
        </div>
      </Router>
    );
  }
}

export default App;
