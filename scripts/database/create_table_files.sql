CREATE TABLE files(
                        id SERIAL,
                        folder_id INT,
                        owner_id INT NOT NULL,
                        name varchar(200) NOT NULL,
                        type varchar(50) NOT NULL,
                        path varchar(50) NOT NULL,
                        created_at TIMESTAMP default current_timestamp,
                        modified_at TIMESTAMP NOT NULL,
                        deleted BOOL NOT NULL DEFAULT false,
                        PRIMARY KEY(id),
                        CONSTRAINT fk_folfers
                            FOREIGN KEY(folders_id)
                                REFERENCES folders(id),
                        CONSTRAINT fk_owners
                            FOREIGN KEY(owner_id)
                                REFERENCES users(id)

)