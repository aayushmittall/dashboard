import React, { Component } from "react";
import { Link } from "react-router-dom";

class SignIn extends Component {
  constructor(props) {
    super(props);
    this.state = { email: "", password: "" };
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

  async handleSubmit(e) {
    e.preventDefault();
    console.log("The form was submitted with the following data:");
    console.log(this.state);
    const data = this.state;
    let apiUrl = "http://localhost:8000/signin";
    const myHeaders = new Headers();
    myHeaders.append("Content-Type", "application/json");
    const options = {
      method: "POST",
      body: JSON.stringify(data),
      myHeaders
    };
    const response = await fetch(apiUrl, options);
    const responsedata = await response.json();
    console.log(responsedata);
  }
  render() {
    return (
      <div className="sign-in">
        <h2 className="form-title">Sign In</h2>

        <form onSubmit={this.handleSubmit} className="form-fields">
          <div className="form-field">
            <label className="form-field-label" htmlFor="email">
              E-Mail/Username
            </label>
            <input
              type="text"
              id="email"
              className="form-field-input"
              placeholder="email"
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
              placeholder="password"
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
