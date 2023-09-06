<template>
  <div class="d-flex justify-content-center container">
    <Form @submit="handleLogin" :validation-schema="schema" style="width: 300px;">
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
        <input type="submit" class="btn btn-primary btn-block mb-4" :disabled="loading" value="Sign in">
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
  name: "Login",
  components: {
    Form,
    Field,
    ErrorMessage,
  },
  data() {
    const schema = yup.object().shape({
      login: yup.string().required("Username is required!"),
      password: yup.string().required("Password is required!"),
    });

    return {
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
  created() {
    if (this.loggedIn) {
      this.$router.push("/");
    }
  },
  methods: {
    handleLogin(user) {
      this.loading = true;

      this.$store.dispatch("auth/login", user).then(
        () => {
          this.$router.push("/");
        },
        (error) => {
          this.loading = false;
          this.message =
            (error.response &&
              error.response.data &&
              error.response.data.message) ||
            error.message ||
            error.toString();
        }
      );
    },
  },
};
</script>