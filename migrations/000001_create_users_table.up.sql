CREATE TABLE IF NOT EXISTS users(
    id SERIAL PRIMARY KEY,
    display_name VARCHAR(200) NOT NULL,
    profile_image_url VARCHAR(200) NOT NULL,
    email VARCHAR(200) NOT NULL,
    twitch_code VARCHAR(200) NOT NULL,
    role VARCHAR(200) NOT NULL,
    access_token VARCHAR(200) NOT NULL,
    refresh_token VARCHAR(200) NOT NULL,
    jwt_token VARCHAR(200) NOT NULL
);