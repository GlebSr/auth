CREATE TABLE refresh (
                         user_id VARCHAR(255) NOT NULL,
                         expire timestamp NOT NULL,
                         token VARCHAR(255) NOT NULL PRIMARY KEY
);

CREATE TABLE oauth (
                       user_id VARCHAR(255) NOT NULL,
                       service VARCHAR(255) NOT NULL,
                       is_refresh boolean NOT NULL,
                       expire timestamp NOT NULL,
                       token VARCHAR(255) NOT NULL
);

CREATE TABLE two_factor (
    user_id VARCHAR(255) NOT NULL PRIMARY KEY,
    secret_key VARCHAR(255) NOT NULL,
    reserve_codes VARCHAR(255)[] NOT NULL
);

CREATE TABLE users (
                            id VARCHAR(255) NOT NULL PRIMARY KEY,
                            email VARCHAR(255),
                            encrypted_password VARCHAR(255),
                            enabled2fa boolean NOT NULL,
                            google_id VARCHAR(255),
                            vk_id VARCHAR(255),
                            yandex_id VARCHAR(255),
                            github_id VARCHAR(255)
);