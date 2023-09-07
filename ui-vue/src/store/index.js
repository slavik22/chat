import { createStore } from "vuex";
import { auth } from "./auth.module";
import { chat } from "./chat.module";
import { message } from "./message.module";

const store = createStore({
  modules: {
    auth,
    chat,
    message,
  },
});

export default store;