import React from "react";
import axios from "axios";

export default class Form extends React.Component {
  constructor(props) {
    super(props);
    this.state = { name: "", email: "" };

    this.handleNameChange = this.handleNameChange.bind(this);
    this.handleEmailChange = this.handleEmailChange.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);
  }

  handleNameChange(event) {
    this.setState({ name: event.target.value });
  }
  handleEmailChange(event) {
    this.setState({ email: event.target.value });
  }

  handleSubmit(event) {
    alert("A name was submitted: " + this.state.name);
    event.preventDefault();
    axios.post(`http://localhost:8080/addStudent`, this.state).then(res => {
      console.log(res);
      console.log(res.data);
    });
  }

  render() {
    
    return (
      <div>
        <h1>{this.state.name}</h1>
        <form onSubmit={this.handleSubmit}>
          <label>
            Name:
            <input
              type="text"
              value={this.state.name}
              onChange={this.handleNameChange}
            />
          </label>
          <label>
            Email:
            <input
              type="email"
              value={this.state.email}
              onChange={this.handleEmailChange}
            />
          </label>
          <input type="submit" value="Submit" />
        </form>
      </div>
    );
  }
}
