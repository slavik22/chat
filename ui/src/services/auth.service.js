import axios from 'axios';

const API_URL = 'http://localhost:8080/api/v1/auth/';

class AuthService {
  login(user) {
    return axios
      .post(API_URL + 'login', {
        login: user.login,
        password: user.password
      })
      .then(response => {
        if (response.data.access_token) {
          localStorage.setItem('user', JSON.stringify(response.data));
        }

        return response.data;
      });
  }

  logout() {
    localStorage.removeItem('user');
  }

  register(user) {
    return axios.post(API_URL + 'register', {
      name: user.username,
      login: user.login,
      password: user.password
    });
  }
}

export default new AuthService();