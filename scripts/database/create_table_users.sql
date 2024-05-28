CREATE TABLE users(
                        id SERIAL,
                        name varchar(80) NOT NULL,
                        login varchar(100) NOT NULL,
                        password varchar(200) NOT NULL,
                        created_at TIMESTAMP default current_timestamp,
                        modified_at TIMESTAMP NOT NULL,
                        deleted BOOL NOT NULL DEFAULT false,
                        last_login TIMESTAMP default current_timestamp,
                        PRIMARY KEY(id)
)