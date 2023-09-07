CREATE TABLE chats (
                            id bigserial PRIMARY KEY,
                            user1_id bigserial NOT NULL,
                            user2_id bigserial NOT NULL,
                            FOREIGN KEY (user1_id) REFERENCES users (id) ON DELETE CASCADE ,
                            FOREIGN KEY (user2_id) REFERENCES users (id) ON DELETE CASCADE,
                            UNIQUE (user1_id, user2_id)
);

CREATE TABLE messages (
                          id bigserial PRIMARY KEY,
                          chat_id bigserial NOT NULL,
                          user_id bigserial NOT NULL,
                          content TEXT NOT NULL,
                          createdAt timestamp DEFAULT NOW(),
                          FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
                          FOREIGN KEY (chat_id) REFERENCES chats (id) ON DELETE CASCADE


);

-- CREATE TABLE user_chat_rooms (
--                                  id bigserial PRIMARY KEY,
--                                  user_id bigserial NOT NULL,
--                                  chat_room_id bigserial NOT NULL,
--                                  FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE ,
--                                  FOREIGN KEY (chat_room_id) REFERENCES chat_rooms (id) ON DELETE CASCADE
-- );