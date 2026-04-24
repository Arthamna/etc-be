
INSERT INTO users (user_id, username, email, password_hash, profile_picture, registration_date, role, created_at, updated_at)
VALUES
('USR-001', 'admin', 'admin@library.test', '$2b$12$adminhash', decode('iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAQAAAC1HAwCAAAAC0lEQVR42mP8/x8AAwMCAO5XWfQAAAAASUVORK5CYII=', 'base64'), now(), 'admin', now(), now()),
('USR-002', 'budi',  'budi@library.test',  '$2b$12$userhash1', decode('iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAQAAAC1HAwCAAAAC0lEQVR42mP8/x8AAwMCAO5XWfQAAAAASUVORK5CYII=', 'base64'), now(), 'user',  now(), now()),
('USR-003', 'sari',  'sari@library.test',  '$2b$12$userhash2', decode('iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAQAAAC1HAwCAAAAC0lEQVR42mP8/x8AAwMCAO5XWfQAAAAASUVORK5CYII=', 'base64'), now(), 'user',  now(), now());

INSERT INTO categories (category_id, name, description, created_at, updated_at)
VALUES
('CAT-001', 'Programming', 'Buku tentang pemrograman dan teknologi', now(), now()),
('CAT-002', 'Productivity', 'Buku pengembangan diri dan produktivitas', now(), now()),
('CAT-003', 'Fantasy', 'Buku fiksi fantasi', now(), now());

INSERT INTO books (book_id, author, book_picture, title, description, status, created_at, updated_at)
VALUES
('BOK-001', 'Robert C. Martin', decode('iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAQAAAC1HAwCAAAAC0lEQVR42mP8/x8AAwMCAO5XWfQAAAAASUVORK5CYII=', 'base64'), 'Clean Code', 'Panduan menulis kode yang rapi dan mudah dipelihara', 'borrowed', now(), now()),
('BOK-002', 'James Clear',      decode('iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAQAAAC1HAwCAAAAC0lEQVR42mP8/x8AAwMCAO5XWfQAAAAASUVORK5CYII=', 'base64'), 'Atomic Habits', 'Buku tentang kebiasaan kecil yang berdampak besar', 'available', now(), now()),
('BOK-003', 'Frank Herbert',    decode('iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAQAAAC1HAwCAAAAC0lEQVR42mP8/x8AAwMCAO5XWfQAAAAASUVORK5CYII=', 'base64'), 'Dune', 'Novel fiksi ilmiah klasik', 'available', now(), now());

INSERT INTO book_categories (book_id, category_id)
VALUES
('BOK-001', 'CAT-001'),
('BOK-002', 'CAT-002'),
('BOK-003', 'CAT-003'),
('BOK-003', 'CAT-001');

INSERT INTO book_borrowings (borrowing_id, user_id, book_id, borrowed_at, returned_at)
VALUES
('BOR-001', 'USR-002', 'BOK-001', now() - interval '3 day', NULL),
('BOR-002', 'USR-003', 'BOK-002', now() - interval '10 day', now() - interval '2 day');

INSERT INTO books (book_id, author, book_picture, title, description, status, created_at, updated_at)
VALUES
('BOK-004', 'George Orwell',   decode('iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAQAAAC1HAwCAAAAC0lEQVR42mP8/x8AAwMCAO5XWfQAAAAASUVORK5CYII=', 'base64'), '1984', 'Novel distopia tentang pengawasan total', 'available', now(), now()),
('BOK-005', 'J.K. Rowling',    decode('iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAQAAAC1HAwCAAAAC0lEQVR42mP8/x8AAwMCAO5XWfQAAAAASUVORK5CYII=', 'base64'), 'Harry Potter', 'Petualangan penyihir muda di Hogwarts', 'borrowed', now(), now()),
('BOK-006', 'Yuval Noah Harari', decode('iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAQAAAC1HAwCAAAAC0lEQVR42mP8/x8AAwMCAO5XWfQAAAAASUVORK5CYII=', 'base64'), 'Sapiens', 'Sejarah singkat umat manusia', 'reserved', now(), now()),
('BOK-007', 'Paulo Coelho',    decode('iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAQAAAC1HAwCAAAAC0lEQVR42mP8/x8AAwMCAO5XWfQAAAAASUVORK5CYII=', 'base64'), 'The Alchemist', 'Kisah perjalanan mencari makna hidup', 'available', now(), now()),
('BOK-008', 'Neil Gaiman',     decode('iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAQAAAC1HAwCAAAAC0lEQVR42mP8/x8AAwMCAO5XWfQAAAAASUVORK5CYII=', 'base64'), 'Coraline', 'Cerita fantasi gelap penuh misteri', 'damaged', now(), now()),
('BOK-009', 'Toni Morrison',   decode('iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAQAAAC1HAwCAAAAC0lEQVR42mP8/x8AAwMCAO5XWfQAAAAASUVORK5CYII=', 'base64'), 'Beloved', 'Novel tentang trauma dan kenangan masa lalu', 'lost', now(), now()),
('BOK-010', 'Dan Brown',       decode('iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAQAAAC1HAwCAAAAC0lEQVR42mP8/x8AAwMCAO5XWfQAAAAASUVORK5CYII=', 'base64'), 'The Da Vinci Code', 'Thriller misteri dengan teka-teki sejarah', 'available', now(), now()),
('BOK-011', 'Malcolm Gladwell', decode('iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAQAAAC1HAwCAAAAC0lEQVR42mP8/x8AAwMCAO5XWfQAAAAASUVORK5CYII=', 'base64'), 'Outliers', 'Buku tentang faktor kesuksesan luar biasa', 'borrowed', now(), now()),
('BOK-012', 'Hanya Yanagihara', decode('iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAQAAAC1HAwCAAAAC0lEQVR42mP8/x8AAwMCAO5XWfQAAAAASUVORK5CYII=', 'base64'), 'A Little Life', 'Kisah persahabatan dan luka batin', 'maintenance', now(), now()),
('BOK-013', 'Stephen King',    decode('iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAQAAAC1HAwCAAAAC0lEQVR42mP8/x8AAwMCAO5XWfQAAAAASUVORK5CYII=', 'base64'), 'The Shining', 'Novel horor klasik di hotel terpencil', 'available', now(), now());

INSERT INTO categories (category_id, name)
VALUES
('CAT-004', 'Technology'),
('CAT-007', 'Fiction'),
('CAT-008', 'Philosophy'),
('CAT-009', 'Psychology'),
('CAT-010', 'History');

INSERT INTO book_categories (book_id, category_id)
VALUES
('BOK-004', 'CAT-007'),
('BOK-004', 'CAT-010'),

('BOK-005', 'CAT-006'),
('BOK-005', 'CAT-007'),

('BOK-006', 'CAT-010'),
('BOK-006', 'CAT-008'),

('BOK-007', 'CAT-008'),
('BOK-007', 'CAT-005'),

('BOK-008', 'CAT-006'),
('BOK-008', 'CAT-007'),

('BOK-009', 'CAT-009'),
('BOK-009', 'CAT-007'),

('BOK-010', 'CAT-007'),
('BOK-010', 'CAT-010'),

('BOK-011', 'CAT-005'),
('BOK-011', 'CAT-009'),

('BOK-012', 'CAT-009'),
('BOK-012', 'CAT-007'),

('BOK-013', 'CAT-006'),
('BOK-013', 'CAT-007');