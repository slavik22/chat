<template>
  <div class="d-flex justify-content-center container">
    <Form @submit="handleRegister" :validation-schema="schema" style="width: 300px;">
      <div class="form-outline mb-4">
        <label for="username">Username</label>
        <Field name="username" type="text" class="form-control" />
        <ErrorMessage name="username" class="form-text error-feedback" />
      </div>

      <div class="form-outline mb-4">
        <label for="login" class="form-label">Login</label>
        <Field name="login" type="text" class="form-control" />
        <ErrorMessage name="login" class="form-text" />
      </div>

      <div class="form-outline mb-4">
        <label for="password" class="form-label">Password</label>
        <Field name="password" type="password" class="form-control" />
        <ErrorMessage name="password" class="form-text" />
      </div>

      <div class="form-outline mb-4">
        <input type="submit" class="btn btn-primary btn-block mb-4" :disabled="loading" value="Sign up" />
        <span v-show="loading" class="spinner-border spinner-border-sm"></span>
      </div>

      <div class="form-outline mb-4">
        <div v-if="message" class="form-text" role="alert">
          {{ message }}
        </div>
      </div>

      <div class="text-center">
        <p>Not a member? <a href="#!">Register</a></p>
      </div>
    </Form>
  </div>
</template>
  
<script>
import { Form, Field, ErrorMessage } from "vee-validate";
import * as yup from "yup";

export default {
  name: "Register",
  components: {
    Form,
    Field,
    ErrorMessage,
  },
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
      password: yup
        .string()
        .required("Password is required!")
        .min(6, "Must be at least 6 characters!")
        .max(40, "Must be maximum 40 characters!"),
    });

    return {
      successful: false,
      loading: false,
      message: "",
      schema,
    };
  },
  computed: {
    loggedIn() {
      return this.$store.state.auth.status.loggedIn;
    },
  },
  mounted() {
    if (this.loggedIn) {
      this.$router.push("/");
    }
  },
  methods: {
    handleRegister(user) {
      this.message = "";
      this.successful = false;
      this.loading = true;

      this.$store.dispatch("auth/register", user).then(
        (data) => {
          this.message = data.message;
          this.successful = true;
          this.loading = false;
          this.$router.push("/login");
        },
        (error) => {
          this.message =
            (error.response &&
              error.response.data &&
              error.response.data.message) ||
            error.message ||
            error.toString();
          this.successful = false;
          this.loading = false;
        }
      );
    },
  },
};
</script>