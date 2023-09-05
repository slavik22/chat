CREATE TABLE chat_rooms (
                            id bigserial PRIMARY KEY,
                            name VARCHAR(255) NOT NULL UNIQUE
);

CREATE TABLE messages (
                          id bigserial PRIMARY KEY,
                          chat_room_id bigserial NOT NULL,
                          user_id bigserial NOT NULL,
                          content TEXT NOT NULL,
                          createdAt timestamp DEFAULT NOW(),
                          FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
                          FOREIGN KEY (chat_room_id) REFERENCES chat_rooms (id) ON DELETE CASCADE


);

CREATE TABLE user_chat_rooms (
                                 id bigserial PRIMARY KEY,
                                 user_id bigserial NOT NULL,
                                 chat_room_id bigserial NOT NULL,
                                 FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE ,
                                 FOREIGN KEY (chat_room_id) REFERENCES chat_rooms (id) ON DELETE CASCADE
);