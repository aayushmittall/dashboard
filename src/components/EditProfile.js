import React, { Component } from "react";

class EditProfile extends Component {
  constructor(props) {
    super(props);

    this.initialState = {
      email: "",
      password: "",
      firstname: "",
      lastname: "",
      age: "",
      username: "",
      country: "",
      gender: ""
    };
    this.state = this.initialState;

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
    console.log(JSON.stringify(data));

    let apiUrl = "http://localhost:8000/editprofile";
    const myHeaders = new Headers();
    myHeaders.append("Content-Type", "application/json");
    const options = {
      method: "POST",
      body: JSON.stringify(data),
      myHeaders
    };
    fetch(apiUrl, options)
      .then(res => res.json())
      .then(result => {
        this.setState({ response: result.res });
      });
  }

  render() {
    return (
      <div className="sign-up">
        <h2>Edit Profile</h2>
        <form onSubmit={this.handleSubmit} className="form-fields">
          <div className="form-field">
            <label className="form-field-label" htmlFor="firstname">
              First Name
            </label>
            <input
              type="text"
              id="firstname"
              className="form-field-input"
              placeholder="Enter your first name"
              name="firstname"
              value={this.state.name}
              onChange={this.handleChange}
              required
            />
          </div>
          <div className="form-field">
            <label className="form-field-label" htmlFor="lastname">
              Last Name
            </label>
            <input
              type="text"
              id="lastname"
              className="form-field-input"
              placeholder="Enter your last name"
              name="lastname"
              value={this.state.name}
              onChange={this.handleChange}
              required
            />
          </div>
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
            <label className="form-field-label" htmlFor="username">
              Username
            </label>
            <input
              type="text"
              id="username"
              className="form-field-input"
              placeholder="Enter your first name"
              name="username"
              value={this.state.name}
              onChange={this.handleChange}
              required
            />
          </div>
          <div className="form-field">
            <label className="form-field-label">Enter your Gender </label>
            <label className="form-field-label" htmlFor="male">
              Male
            </label>
            <input
              className="form-field-label"
              type="radio"
              name="gender"
              id="male"
              value="male"
              onChange={this.handleChange}
              required
            />
            <label className="form-field-label" htmlFor="female">
              Female
            </label>
            <input
              type="radio"
              name="gender"
              id="female"
              value="female"
              onChange={this.handleChange}
              required
            />
          </div>

          <div className="form-field">
            <label htmlFor="age"> Age </label>
            <select onChange={this.handleChange} name="age" id="age" required>
              <option>age</option>
              <option value="18">18</option>
              <option value="19">19</option>
              <option value="20">20</option>
              <option value="21">21</option>
              <option value="22">22</option>
              <option value="23">23</option>
              <option value="24">24</option>
              <option value="25">25</option>
              <option value="26">26</option>
              <option value="27">27</option>
              <option value="28">28</option>
              <option value="29">29</option>
              <option value="30">30</option>
            </select>
          </div>
          <div className="form-field">
            <label className="form-field-label" htmlFor="country">
              Country
            </label>
            <input
              type="text"
              id="country"
              className="form-field-input"
              placeholder="Enter your country"
              name="country"
              value={this.state.name}
              onChange={this.handleChange}
              required
            />
          </div>

          <div className="form-field">
            <button type="submit" className="sign-button">
              Submit
            </button>
          </div>
        </form>
      </div>
    );
  }
}
export default EditProfile;
