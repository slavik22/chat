<template>
  <div class="container mt-4">
    <form class="form-outline mb-4" @submit="uploadImage">
        <input type="file" @change="handleFileUpload" accept="image/*" required class="form-label"><br>
        <input type="submit" class="btn btn-success" value="Upload">
    </form>

    <Form @submit="updateProfile" :validation-schema="schema">
      <div class="form-outline">
        <label for="name">Name:</label>
        <Field v-model="currentUser.name" name="name" type="text" class="form-control" />
        <ErrorMessage name="name" class="form-text error-feedback" />
      </div>

      <div class="form-outline">
        <label for="login" class="form-label">Login</label>
        <Field v-model="currentUser.login" name="login" type="text" class="form-control" />
        <ErrorMessage name="login" class="form-text error-feedback" />
      </div>

      <div class="form-outline mb-4">
        <label for="password" class="form-label">Password</label>
        <Field v-model="password" name="password" type="password" class="form-control" />
        <ErrorMessage name="password" class="form-text" />
      </div>

      <div class="form-outline mb-4">
        <input type="submit" class="btn btn-primary btn-block mb-4" value="Update" />
        <span v-show="loading" class="spinner-border spinner-border-sm"></span>
      </div>

      <div class="form-outline mb-4">
        <div v-if="message" class="form-text" role="alert">
          {{ message }}
        </div>
      </div>
    </Form>

    <div v-if="friends" class="friends">
      <h2>Friends</h2>

      <ul class="list-group">
        <li v-for="user in friends" :key="user.id" class="list-group-item">
          {{ user.name }}
          <button @click="deleteFriend(user.id)" class="btn btn-danger">Delete</button>
        </li>
      </ul>
    </div>

    <div v-if="blackList" class="blackList">
      <h2>BlackList</h2>

      <ul class="list-group">
        <li v-for="user in blackList" :key="user.id" class="list-group-item">
          {{ user.name }}
          <button @click="deleteBlackList(user.id)" class="btn btn-danger">Delete</button>
        </li>
      </ul>
    </div>
  </div>
</template>
  
<script>
import axios from 'axios';
import authHeader from '@/services/auth-header';

import { Form, Field, ErrorMessage } from "vee-validate";
import * as yup from "yup";

export default {
  data() {
    const schema = yup.object().shape({
      username: yup
        .string()
        .required("Username is required!")
        .min(3, "Must be at least 3 characters!")
        .max(20, "Must be maximum 20 characters!"),
      login: yup
        .string()
        .required("Login is required!")
        .min(3, "Must be at least 3 characters!")
        .max(50, "Must be maximum 50 characters!"),
    });

    return {
      file: null,
      password: "",
      friends: null,
      blackList: null,
      message: "",
      loading: false,
      schema,
    };
  },
  components: {
    Form,
    Field,
    ErrorMessage,
  },

  computed: {
    currentUser() {
      return this.$store.state.auth.user;
    },
  },
  mounted() {
    if (!this.currentUser) {
      this.$router.push("/login");
    }

    axios.get('http://localhost:8080/api/v1/users/friends/', { headers: authHeader() })
      .then((response) => {
        this.friends = response.data;
      })
      .catch((error) => {
        console.error('Error fetching users:', error);
      });

    axios.get('http://localhost:8080/api/v1/users/blackList/', { headers: authHeader() })
      .then((response) => {
        this.blackList = response.data;
      })
      .catch((error) => {
        console.error('Error fetching users:', error);
      });

  },
  methods: {
    handleFileUpload(event) {
      this.file = event.target.files[0];
    },
    uploadImage(e) {
      e.preventDefault()
      const formData = new FormData();
      formData.append('image', this.file);

      axios.post('http://localhost:8080/api/v1/users/upload', formData, {headers: authHeader() })
        .then(response => response.data)
        .then(message => {
          console.log(message);
        })
        .catch(error => {
          this.message =
              (error.response &&
                error.response.data &&
                error.response.data.message) ||
              error.message ||
              error.toString();
        });
    },
    async updateProfile() {
      this.loading = true;

      const formData = new FormData();
      formData.append('image', this.file);

      await axios.put("http://localhost:8080/api/v1/users/", { name: this.currentUser.name, login: this.currentUser.login, password: this.password, image: formData }, { headers: authHeader() })
        .then(
          () => {
            this.$store.dispatch("auth/logout")
            this.$router.push("/login");
          },
          (error) => {
            this.loading = false;
            this.message =
              (error.response &&
                error.response.data &&
                error.response.data.message) ||
              error.message ||
              error.toString();
          })

    },

    async deleteFriend(userId) {
      await axios.delete("http://localhost:8080/api/v1/users/friends/" + userId, { headers: authHeader() });
      location.reload()

    },
    async deleteBlackList(userId) {
      await axios.delete("http://localhost:8080/api/v1/users/blackList/" + userId, { headers: authHeader() });
      location.reload()
    }
  },
}
</script>
  