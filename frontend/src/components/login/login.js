import React from 'react';
import './Login.css';

class Login extends React.Component {
  state = {
    username: '',
    password: '',
    errorMessage: ''
  };

  handleSubmit = (event) => {
    event.preventDefault();

    const { username, password } = this.state;

    if (username === '' || password === '') {
      this.setState({ errorMessage: 'Username and password are required!' });
      return;
    }

    // Make API call to the user microservice to check if the username and password are valid
    // ...

    // Clear the form
    this.setState({ username: '', password: '', errorMessage: '' });
  };

  render() {
    const { username, password, errorMessage } = this.state;

    return (
      <div className="login-container">
        <h1>Login</h1>
        {errorMessage && <p className="error-message">{errorMessage}</p>}
        <form onSubmit={this.handleSubmit}>
          <input
            type="text"
            placeholder="Username"
            value={username}
            onChange={(event) => this.setState({ username: event.target.value })}
          />
          <input
            type="password"
            placeholder="Password"
            value={password}
            onChange={(event) => this.setState({ password: event.target.value })}
          />
          <button type="submit">Submit</button>
        </form>
      </div>
    );
  }
}

export default Login;
