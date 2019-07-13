import React, { Component } from "react";
import { Link } from "react-router-dom";

class SignIn extends Component {
  constructor(props) {
    super(props);
    this.state = { email: "", password: "", username: "" };

    this.handleChange = this.handleChange.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);
  }

  handleChange(e) {
    let target = e.target;
    let value = target.value;
    let name = target.name;

    this.setState({
      [name]: value
    });
  }

  handleSubmit(e) {
    e.preventDefault();
    console.log("The form was submitted with the following data:");
    console.log(this.state);
    const data = this.state;
    const myHeaders = new Headers();
    myHeaders.append("Content-Type", "application/json");
    const options = {
      method: "POST",
      myHeaders,
      body: JSON.stringify(data)
    };
    fetch("http://localhost:8000/signin", options).then(res => res.json());
  }
  render() {
    return (
      <div className="sign-in">
        <form onSubmit={this.handleSubmit} className="form-fields">
          <div className="form-field">
            <label className="form-field-label" htmlFor="email">
              E-Mail Address
            </label>
            <input
              type="email"
              id="email"
              className="form-field-input"
              placeholder="Enter your email"
              name="email"
              value={this.state.email}
              onChange={this.handleChange}
              required
            />
          </div>
          <div className="form-field">
            <label className="form-field-label" htmlFor="password">
              Password
            </label>
            <input
              type="password"
              id="password"
              className="form-field-input"
              placeholder="Enter your password"
              name="password"
              value={this.state.password}
              onChange={this.handleChange}
              required
              minLength="5"
              maxLength="15"
            />
          </div>

          <div className="form-field">
            <button type="submit" className="sign-button">
              Sign In
            </button>
          </div>
          <div className="form-field">
            <Link to="/" className="link-already">
              Not Registered?
            </Link>
          </div>
        </form>
      </div>
    );
  }
}
export default SignIn;
