-- +migrate Up

INSERT INTO users (name, last_name, email, password, is_admin) VALUES (
    '',
    '',
    'admin',
    '$2a$10$LK.dwSquS83Lz.juOzn.0OfeQK95a2B5fXzlv/aoPb/7SrvGbNbBW',
    true
);
    
-- +migrate Down
DELETE FROM users WHERE email = 'admin';
