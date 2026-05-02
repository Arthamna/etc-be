-- ====================================================================
-- SEED DATA - PostgreSQL
-- Relasi yang dijamin:
--   • 10 dosen (user-01..10) menjadi pembuat rekrutmen & ketua tim
--   • 30 mahasiswa (user-11..40) mendaftar dan bergabung tim
--   • Setiap rekrutmen punya tim, setiap tim punya peserta
--   • Ada pendaftar accepted, rejected, dan pending
--   • History review dibuat dosen terhadap anggota timnya
-- ====================================================================
-- CATATAN: Nama tabel menggunakan konvensi default GORM (snake_case plural).
-- Sesuaikan jika kamu menggunakan TableName() custom di setiap model.
-- Contoh: Rekrutmen -> rekrutmens, Tim -> tims, dll.
-- ====================================================================

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Hapus data lama jika perlu (urutan penting karena ada FK):
-- TRUNCATE TABLE bookmarks, histories, tim_participants, pendaftars, tims, rekrutmens, users, setting_drives RESTART IDENTITY CASCADE;


-- ====================================================================
-- 1. USERS (40 rows: 10 dosen + 30 mahasiswa)
-- ====================================================================
INSERT INTO users (user_id, nama, jurusan, no_pengenal, role, no_telp, password_hash, profile_picture, created_at, updated_at, spesialisasi) VALUES

-- Dosen (user-01 s.d. user-10) — mereka yang create rekrutmen & jadi ketua tim
('user-01','Dr. Andi Wijaya',   'Teknik Informatika','NIP198001010001','dosen','081234560001','$2a$10$seedhash.user01',NULL,'2023-01-01 08:00:00','2023-01-01 08:00:00','["Machine Learning","Data Science"]'),
('user-02','Dr. Budi Santoso',  'Sistem Informasi',  'NIP197902020002','dosen','081234560002','$2a$10$seedhash.user02',NULL,'2023-01-01 08:00:00','2023-01-01 08:00:00','["Web Development","Cloud Computing"]'),
('user-03','Dr. Citra Dewi',    'Matematika',        'NIP198103030003','dosen','081234560003','$2a$10$seedhash.user03',NULL,'2023-01-01 08:00:00','2023-01-01 08:00:00','["Statistika","Data Science"]'),
('user-04','Dr. Dian Pratama',  'Teknik Elektro',    'NIP198204040004','dosen','081234560004','$2a$10$seedhash.user04',NULL,'2023-01-01 08:00:00','2023-01-01 08:00:00','["IoT","Embedded Systems"]'),
('user-05','Dr. Eka Susanti',   'Teknik Komputer',   'NIP198305050005','dosen','081234560005','$2a$10$seedhash.user05',NULL,'2023-01-01 08:00:00','2023-01-01 08:00:00','["Computer Vision","AI"]'),
('user-06','Dr. Fajar Nugroho', 'Teknik Informatika','NIP198406060006','dosen','081234560006','$2a$10$seedhash.user06',NULL,'2023-01-01 08:00:00','2023-01-01 08:00:00','["Cybersecurity","Networking"]'),
('user-07','Dr. Gita Rahayu',   'Sistem Informasi',  'NIP198507070007','dosen','081234560007','$2a$10$seedhash.user07',NULL,'2023-01-01 08:00:00','2023-01-01 08:00:00','["Business Intelligence","ERP"]'),
('user-08','Dr. Hendra Kusuma', 'Matematika',        'NIP198608080008','dosen','081234560008','$2a$10$seedhash.user08',NULL,'2023-01-01 08:00:00','2023-01-01 08:00:00','["Algoritma","Kriptografi"]'),
('user-09','Dr. Indira Sari',   'Teknik Elektro',    'NIP198709090009','dosen','081234560009','$2a$10$seedhash.user09',NULL,'2023-01-01 08:00:00','2023-01-01 08:00:00','["Robotika","Kontrol Sistem"]'),
('user-10','Dr. Joko Widodo',   'Teknik Komputer',   'NIP198810100010','dosen','081234560010','$2a$10$seedhash.user10',NULL,'2023-01-01 08:00:00','2023-01-01 08:00:00','["Game Development","VR"]'),

-- Mahasiswa batch 2021 (user-11 s.d. user-30)
('user-11','Aldi Firmansyah',  'Teknik Informatika','NIM20210001','mahasiswa','081234561101','$2a$10$seedhash.user11',NULL,'2023-06-01 08:00:00','2023-06-01 08:00:00','["Python","React"]'),
('user-12','Bella Octavia',    'Sistem Informasi',  'NIM20210002','mahasiswa','081234561102','$2a$10$seedhash.user12',NULL,'2023-06-01 08:00:00','2023-06-01 08:00:00','["UI/UX","Figma"]'),
('user-13','Cahya Pratama',    'Teknik Komputer',   'NIM20210003','mahasiswa','081234561103','$2a$10$seedhash.user13',NULL,'2023-06-01 08:00:00','2023-06-01 08:00:00','["C++","Arduino"]'),
('user-14','Dewi Anggraeni',   'Teknik Informatika','NIM20210004','mahasiswa','081234561104','$2a$10$seedhash.user14',NULL,'2023-06-01 08:00:00','2023-06-01 08:00:00','["Java","Spring Boot"]'),
('user-15','Eka Prasetyo',     'Sistem Informasi',  'NIM20210005','mahasiswa','081234561105','$2a$10$seedhash.user15',NULL,'2023-06-01 08:00:00','2023-06-01 08:00:00','["PHP","Laravel"]'),
('user-16','Fina Maharani',    'Matematika',        'NIM20210006','mahasiswa','081234561106','$2a$10$seedhash.user16',NULL,'2023-06-01 08:00:00','2023-06-01 08:00:00','["R","Statistika"]'),
('user-17','Gilang Saputra',   'Teknik Elektro',    'NIM20210007','mahasiswa','081234561107','$2a$10$seedhash.user17',NULL,'2023-06-01 08:00:00','2023-06-01 08:00:00','["PCB Design","Microcontroller"]'),
('user-18','Hana Puspita',     'Teknik Informatika','NIM20210008','mahasiswa','081234561108','$2a$10$seedhash.user18',NULL,'2023-06-01 08:00:00','2023-06-01 08:00:00','["Data Science","TensorFlow"]'),
('user-19','Ilham Maulana',    'Sistem Informasi',  'NIM20210009','mahasiswa','081234561109','$2a$10$seedhash.user19',NULL,'2023-06-01 08:00:00','2023-06-01 08:00:00','["Node.js","Express"]'),
('user-20','Jasmine Putri',    'Teknik Komputer',   'NIM20210010','mahasiswa','081234561110','$2a$10$seedhash.user20',NULL,'2023-06-01 08:00:00','2023-06-01 08:00:00','["Flutter","Dart"]'),
('user-21','Kevin Hidayat',    'Teknik Informatika','NIM20210011','mahasiswa','081234561111','$2a$10$seedhash.user21',NULL,'2023-06-01 08:00:00','2023-06-01 08:00:00','["Go","Microservices"]'),
('user-22','Laila Nur',        'Sistem Informasi',  'NIM20210012','mahasiswa','081234561112','$2a$10$seedhash.user22',NULL,'2023-06-01 08:00:00','2023-06-01 08:00:00','["SQL","PostgreSQL"]'),
('user-23','Mario Setiawan',   'Teknik Komputer',   'NIM20210013','mahasiswa','081234561113','$2a$10$seedhash.user23',NULL,'2023-06-01 08:00:00','2023-06-01 08:00:00','["Networking","Cisco"]'),
('user-24','Nadia Fitri',      'Teknik Informatika','NIM20210014','mahasiswa','081234561114','$2a$10$seedhash.user24',NULL,'2023-06-01 08:00:00','2023-06-01 08:00:00','["Machine Learning","Python"]'),
('user-25','Omar Abdullah',    'Sistem Informasi',  'NIM20210015','mahasiswa','081234561115','$2a$10$seedhash.user25',NULL,'2023-06-01 08:00:00','2023-06-01 08:00:00','["Blockchain","Web3"]'),
('user-26','Putri Ayu',        'Matematika',        'NIM20210016','mahasiswa','081234561116','$2a$10$seedhash.user26',NULL,'2023-06-01 08:00:00','2023-06-01 08:00:00','["Algoritma","Matematika Diskrit"]'),
('user-27','Rizki Ramadhan',   'Teknik Elektro',    'NIM20210017','mahasiswa','081234561117','$2a$10$seedhash.user27',NULL,'2023-06-01 08:00:00','2023-06-01 08:00:00','["PLC","SCADA"]'),
('user-28','Sari Wulandari',   'Teknik Informatika','NIM20210018','mahasiswa','081234561118','$2a$10$seedhash.user28',NULL,'2023-06-01 08:00:00','2023-06-01 08:00:00','["DevOps","Docker"]'),
('user-29','Taufik Hidayat',   'Sistem Informasi',  'NIM20210019','mahasiswa','081234561119','$2a$10$seedhash.user29',NULL,'2023-06-01 08:00:00','2023-06-01 08:00:00','["SAP","Business Process"]'),
('user-30','Ulfa Rahmawati',   'Teknik Komputer',   'NIM20210020','mahasiswa','081234561120','$2a$10$seedhash.user30',NULL,'2023-06-01 08:00:00','2023-06-01 08:00:00','["Android","Kotlin"]'),

-- Mahasiswa batch 2022 (user-31 s.d. user-40)
('user-31','Vino Pratama',   'Teknik Informatika','NIM20220001','mahasiswa','081234561121','$2a$10$seedhash.user31',NULL,'2023-09-01 08:00:00','2023-09-01 08:00:00','["React Native","JavaScript"]'),
('user-32','Winda Sari',     'Sistem Informasi',  'NIM20220002','mahasiswa','081234561122','$2a$10$seedhash.user32',NULL,'2023-09-01 08:00:00','2023-09-01 08:00:00','["Power BI","Tableau"]'),
('user-33','Xander Putra',   'Teknik Komputer',   'NIM20220003','mahasiswa','081234561123','$2a$10$seedhash.user33',NULL,'2023-09-01 08:00:00','2023-09-01 08:00:00','["Linux","Shell Scripting"]'),
('user-34','Yuni Astuti',    'Teknik Informatika','NIM20220004','mahasiswa','081234561124','$2a$10$seedhash.user34',NULL,'2023-09-01 08:00:00','2023-09-01 08:00:00','["Vue.js","TypeScript"]'),
('user-35','Zaki Maulana',   'Sistem Informasi',  'NIM20220005','mahasiswa','081234561125','$2a$10$seedhash.user35',NULL,'2023-09-01 08:00:00','2023-09-01 08:00:00','["Kubernetes","AWS"]'),
('user-36','Adelia Putri',   'Matematika',        'NIM20220006','mahasiswa','081234561126','$2a$10$seedhash.user36',NULL,'2023-09-01 08:00:00','2023-09-01 08:00:00','["MATLAB","Simulink"]'),
('user-37','Bagas Nugraha',  'Teknik Elektro',    'NIM20220007','mahasiswa','081234561127','$2a$10$seedhash.user37',NULL,'2023-09-01 08:00:00','2023-09-01 08:00:00','["Raspberry Pi","Python"]'),
('user-38','Clara Octaviani','Teknik Informatika','NIM20220008','mahasiswa','081234561128','$2a$10$seedhash.user38',NULL,'2023-09-01 08:00:00','2023-09-01 08:00:00','["Swift","iOS"]'),
('user-39','Dimas Saputro',  'Sistem Informasi',  'NIM20220009','mahasiswa','081234561129','$2a$10$seedhash.user39',NULL,'2023-09-01 08:00:00','2023-09-01 08:00:00','["GraphQL","Apollo"]'),
('user-40','Erni Susanti',   'Teknik Komputer',   'NIM20220010','mahasiswa','081234561130','$2a$10$seedhash.user40',NULL,'2023-09-01 08:00:00','2023-09-01 08:00:00','["Unity","C#"]');


-- ====================================================================
-- 2. REKRUTMEN (40 rows)
-- Pembuat (user_id) berputar user-01..user-10, sehingga setiap dosen
-- menjadi pembuat 4 rekrutmen dan otomatis menjadi ketua tim-nya.
-- Kegiatan berputar: riset, projek, lomba
-- ====================================================================
INSERT INTO rekrutmens (rekrutmen_id, user_id, kegiatan, kriteria, tanggal_mulai, tanggal_selesai, fee, role, contact_person) VALUES
('rek-01','user-01','riset',  'IPK min 3.5, Python, Machine Learning',          '2024-02-01','2024-07-31',  500000, 'Asisten Peneliti ML',       'wa.me/6281234560001'),
('rek-02','user-02','projek', 'React, Node.js, pengalaman minimal 6 bulan',     '2024-02-15','2024-08-31',  750000, 'Full Stack Developer',      'wa.me/6281234560002'),
('rek-03','user-03','lomba',  'Mahasiswa S1 aktif, tim 3 orang, IPK min 3.0',  '2024-03-01','2024-05-31',  NULL,   'Anggota Tim Lomba Statistik','wa.me/6281234560003'),
('rek-04','user-04','riset',  'Arduino, sensor IoT, basic C/C++',               '2024-03-15','2024-09-30',  600000, 'Peneliti IoT',              'wa.me/6281234560004'),
('rek-05','user-05','projek', 'OpenCV, Python, Computer Vision dasar',          '2024-04-01','2024-10-31',  800000, 'Computer Vision Engineer',  'wa.me/6281234560005'),
('rek-06','user-06','lomba',  'Keamanan jaringan, pengalaman CTF',              '2024-04-15','2024-06-30',  NULL,   'CTF Team Member',           'wa.me/6281234560006'),
('rek-07','user-07','riset',  'Power BI atau Tableau, analisis data bisnis',    '2024-05-01','2024-11-30',  550000, 'Analis Data Bisnis',        'wa.me/6281234560007'),
('rek-08','user-08','projek', 'Algoritma lanjut, C++ atau Java',                '2024-05-15','2024-12-31',  700000, 'Backend Developer',         'wa.me/6281234560008'),
('rek-09','user-09','lomba',  'Robotika, Arduino atau ROS, tim 2-4 orang',     '2024-06-01','2024-08-31',  NULL,   'Anggota Tim Robot',         'wa.me/6281234560009'),
('rek-10','user-10','riset',  'Unity atau Unreal Engine, C# atau C++',          '2024-06-15','2024-12-31',  650000, 'Game Researcher',           'wa.me/6281234560010'),
('rek-11','user-01','projek', 'NLP, HuggingFace Transformers, Python',          '2024-07-01','2025-01-31',  900000, 'NLP Engineer',              'wa.me/6281234560001'),
('rek-12','user-02','lomba',  'Web hackathon 48 jam, tim 2-3 orang',           '2024-07-15','2024-09-30',  NULL,   'Anggota Tim Hackathon',     'wa.me/6281234560002'),
('rek-13','user-03','riset',  'Statistika inferensial, R atau Python',           '2024-08-01','2025-02-28',  520000, 'Asisten Riset Statistik',   'wa.me/6281234560003'),
('rek-14','user-04','projek', 'SCADA, PLC, otomasi industri dasar',             '2024-08-15','2025-03-31',  850000, 'Automation Engineer',       'wa.me/6281234560004'),
('rek-15','user-05','lomba',  'Kompetisi AI/ML nasional, track machine learning','2024-09-01','2024-11-30',  NULL,   'AI Competitor',             'wa.me/6281234560005'),
('rek-16','user-06','riset',  'Penetration testing, ethical hacking',           '2024-09-15','2025-03-31',  575000, 'Security Researcher',       'wa.me/6281234560006'),
('rek-17','user-07','projek', 'SAP ERP, business process analysis',             '2024-10-01','2025-04-30',  950000, 'SAP Junior Consultant',     'wa.me/6281234560007'),
('rek-18','user-08','lomba',  'Olimpiade matematika, kalkulus dan aljabar',     '2024-10-15','2024-12-15',  NULL,   'Peserta Olimpiade Matematika','wa.me/6281234560008'),
('rek-19','user-09','riset',  'Kontrol sistem, MATLAB dan Simulink',             '2024-11-01','2025-05-31',  610000, 'Peneliti Kontrol Sistem',   'wa.me/6281234560009'),
('rek-20','user-10','projek', 'Unity VR/AR, C#, pengalaman 3D modeling',        '2024-11-15','2025-05-31', 1000000, 'VR Developer',              'wa.me/6281234560010'),
('rek-21','user-01','lomba',  'Data science competition, pengalaman Kaggle',    '2024-12-01','2025-02-28',  NULL,   'Data Scientist Kompetisi',  'wa.me/6281234560001'),
('rek-22','user-02','riset',  'Cloud architecture, AWS atau GCP, DevOps',       '2024-12-15','2025-06-30',  720000, 'Cloud Researcher',          'wa.me/6281234560002'),
('rek-23','user-03','projek', 'Image processing, scikit-image, Python',         '2025-01-01','2025-07-31',  680000, 'Image Processing Developer','wa.me/6281234560003'),
('rek-24','user-04','lomba',  'Desain PCB, KiCad atau Altium Designer',         '2025-01-15','2025-03-31',  NULL,   'PCB Designer',              'wa.me/6281234560004'),
('rek-25','user-05','riset',  'Deep learning, PyTorch, GPU computing',           '2025-02-01','2025-08-31',  850000, 'Deep Learning Researcher',  'wa.me/6281234560005'),
('rek-26','user-06','projek', 'DevSecOps, Docker, Kubernetes, CI/CD pipeline',  '2025-02-15','2025-08-31',  900000, 'DevSecOps Engineer',        'wa.me/6281234560006'),
('rek-27','user-07','lomba',  'Business case competition, analisis pasar',      '2025-03-01','2025-05-31',  NULL,   'Business Analyst Kompetisi','wa.me/6281234560007'),
('rek-28','user-08','riset',  'Kriptografi modern, implementasi algoritma',      '2025-03-15','2025-09-30',  630000, 'Cryptography Researcher',   'wa.me/6281234560008'),
('rek-29','user-09','projek', 'Robot otonom, ROS 2, Python',                    '2025-04-01','2025-10-31',  950000, 'Robotics Developer',        'wa.me/6281234560009'),
('rek-30','user-10','lomba',  'Game jam 72 jam, Unity atau Godot, tim 2-3',    '2025-04-15','2025-06-30',  NULL,   'Game Developer Kompetisi',  'wa.me/6281234560010'),
('rek-31','user-01','riset',  'Reinforcement learning, OpenAI Gym, PyTorch',    '2025-05-01','2025-11-30',  780000, 'RL Researcher',             'wa.me/6281234560001'),
('rek-32','user-02','projek', 'Microservices, Go, gRPC, Apache Kafka',          '2025-05-15','2025-11-30', 1100000, 'Backend Microservices Eng', 'wa.me/6281234560002'),
('rek-33','user-03','lomba',  'Olimpiade statistika, inferensi dan regresi',    '2025-06-01','2025-08-31',  NULL,   'Statistics Competitor',     'wa.me/6281234560003'),
('rek-34','user-04','riset',  'Smart grid, IoT sensor network, MQTT',           '2025-06-15','2025-12-31',  700000, 'Smart Grid Researcher',     'wa.me/6281234560004'),
('rek-35','user-05','projek', 'MLOps, model deployment, FastAPI, monitoring',   '2025-07-01','2026-01-31', 1050000, 'MLOps Engineer',            'wa.me/6281234560005'),
('rek-36','user-06','lomba',  'CTF nasional, tim 4 orang, kategori web & pwn',  '2025-07-15','2025-09-30',  NULL,   'CTF Team Captain',          'wa.me/6281234560006'),
('rek-37','user-07','riset',  'Digital transformation, analisis ERP & BPM',    '2025-08-01','2026-02-28',  660000, 'Digital Transform Researcher','wa.me/6281234560007'),
('rek-38','user-08','projek', 'Compiler design, LLVM IR, bahasa pemrograman',   '2025-08-15','2026-02-28',  800000, 'Compiler Engineer',         'wa.me/6281234560008'),
('rek-39','user-09','lomba',  'KRTI drone competition, autonomous flight mode', '2025-09-01','2025-11-30',  NULL,   'Drone Pilot Researcher',    'wa.me/6281234560009'),
('rek-40','user-10','riset',  'Spatial computing, WebXR, AR/MR development',    '2025-09-15','2026-03-31',  875000, 'Spatial Computing Researcher','wa.me/6281234560010');


-- ====================================================================
-- 3. TIMS (40 rows)
-- Satu tim per rekrutmen.
-- nama_ketua = nama user pembuat rekrutmen tersebut.
-- Pola: rek-01..10 → user-01..10 berulang untuk rek-11..40
-- ====================================================================
INSERT INTO tims (tim_id, tipe_tim, rekrutmen_id, nama_ketua, created_at) VALUES
('tim-01','kelompok','rek-01','Dr. Andi Wijaya',   '2024-02-10 09:00:00'),
('tim-02','kelompok','rek-02','Dr. Budi Santoso',  '2024-02-25 09:00:00'),
('tim-03','kelompok','rek-03','Dr. Citra Dewi',    '2024-03-10 09:00:00'),
('tim-04','individu','rek-04','Dr. Dian Pratama',  '2024-03-25 09:00:00'),
('tim-05','kelompok','rek-05','Dr. Eka Susanti',   '2024-04-10 09:00:00'),
('tim-06','kelompok','rek-06','Dr. Fajar Nugroho', '2024-04-25 09:00:00'),
('tim-07','individu','rek-07','Dr. Gita Rahayu',   '2024-05-10 09:00:00'),
('tim-08','kelompok','rek-08','Dr. Hendra Kusuma', '2024-05-25 09:00:00'),
('tim-09','kelompok','rek-09','Dr. Indira Sari',   '2024-06-10 09:00:00'),
('tim-10','individu','rek-10','Dr. Joko Widodo',   '2024-06-25 09:00:00'),
('tim-11','kelompok','rek-11','Dr. Andi Wijaya',   '2024-07-10 09:00:00'),
('tim-12','kelompok','rek-12','Dr. Budi Santoso',  '2024-07-25 09:00:00'),
('tim-13','individu','rek-13','Dr. Citra Dewi',    '2024-08-10 09:00:00'),
('tim-14','kelompok','rek-14','Dr. Dian Pratama',  '2024-08-25 09:00:00'),
('tim-15','kelompok','rek-15','Dr. Eka Susanti',   '2024-09-10 09:00:00'),
('tim-16','individu','rek-16','Dr. Fajar Nugroho', '2024-09-25 09:00:00'),
('tim-17','kelompok','rek-17','Dr. Gita Rahayu',   '2024-10-10 09:00:00'),
('tim-18','kelompok','rek-18','Dr. Hendra Kusuma', '2024-10-25 09:00:00'),
('tim-19','individu','rek-19','Dr. Indira Sari',   '2024-11-10 09:00:00'),
('tim-20','kelompok','rek-20','Dr. Joko Widodo',   '2024-11-25 09:00:00'),
('tim-21','kelompok','rek-21','Dr. Andi Wijaya',   '2024-12-10 09:00:00'),
('tim-22','individu','rek-22','Dr. Budi Santoso',  '2024-12-25 09:00:00'),
('tim-23','kelompok','rek-23','Dr. Citra Dewi',    '2025-01-10 09:00:00'),
('tim-24','kelompok','rek-24','Dr. Dian Pratama',  '2025-01-25 09:00:00'),
('tim-25','individu','rek-25','Dr. Eka Susanti',   '2025-02-10 09:00:00'),
('tim-26','kelompok','rek-26','Dr. Fajar Nugroho', '2025-02-25 09:00:00'),
('tim-27','kelompok','rek-27','Dr. Gita Rahayu',   '2025-03-10 09:00:00'),
('tim-28','individu','rek-28','Dr. Hendra Kusuma', '2025-03-25 09:00:00'),
('tim-29','kelompok','rek-29','Dr. Indira Sari',   '2025-04-10 09:00:00'),
('tim-30','kelompok','rek-30','Dr. Joko Widodo',   '2025-04-25 09:00:00'),
('tim-31','individu','rek-31','Dr. Andi Wijaya',   '2025-05-10 09:00:00'),
('tim-32','kelompok','rek-32','Dr. Budi Santoso',  '2025-05-25 09:00:00'),
('tim-33','kelompok','rek-33','Dr. Citra Dewi',    '2025-06-10 09:00:00'),
('tim-34','individu','rek-34','Dr. Dian Pratama',  '2025-06-25 09:00:00'),
('tim-35','kelompok','rek-35','Dr. Eka Susanti',   '2025-07-10 09:00:00'),
('tim-36','kelompok','rek-36','Dr. Fajar Nugroho', '2025-07-25 09:00:00'),
('tim-37','individu','rek-37','Dr. Gita Rahayu',   '2025-08-10 09:00:00'),
('tim-38','kelompok','rek-38','Dr. Hendra Kusuma', '2025-08-25 09:00:00'),
('tim-39','kelompok','rek-39','Dr. Indira Sari',   '2025-09-10 09:00:00'),
('tim-40','individu','rek-40','Dr. Joko Widodo',   '2025-09-25 09:00:00');


-- ====================================================================
-- 4. PENDAFTARS (40 rows)
-- Distribusi status:
--   pend-01..10  → rek-01..10,  user-11..20, status = accepted
--   pend-11..20  → rek-01..10,  user-21..30, status = rejected
--   pend-21..30  → rek-11..20,  user-31..40, status = pending
--   pend-31..40  → rek-21..30,  user-11..20, status = accepted
-- ====================================================================
INSERT INTO pendaftars (pendaftar_id, rekrutmen_id, user_id, alasan_mendaftar, cv_url, portofolio_url, status, created_at) VALUES
-- accepted (rek-01..10, user-11..20)
('pend-01','rek-01','user-11','Ingin mendalami ML dalam lingkungan riset nyata',              'https://drive.google.com/cv/user-11','https://github.com/aldi-firmansyah',       'accepted','2024-02-15 10:00:00'),
('pend-02','rek-02','user-12','Tertarik membangun produk web full stack berskala besar',     'https://drive.google.com/cv/user-12','https://portfolio.bella-octavia.com',       'accepted','2024-03-01 10:00:00'),
('pend-03','rek-03','user-13','Ingin mengasah skill lomba bersama mentor berpengalaman',     'https://drive.google.com/cv/user-13','https://github.com/cahya-pratama',          'accepted','2024-03-10 10:00:00'),
('pend-04','rek-04','user-14','Pengalaman IoT cocok dengan topik riset ini',                 'https://drive.google.com/cv/user-14','https://portfolio.dewi-anggraeni.com',      'accepted','2024-03-20 10:00:00'),
('pend-05','rek-05','user-15','Ingin mengaplikasikan PHP ke sistem Computer Vision',         'https://drive.google.com/cv/user-15','https://github.com/eka-prasetyo',           'accepted','2024-04-05 10:00:00'),
('pend-06','rek-06','user-16','Tertarik keamanan siber dan ingin belajar dari dosen ahlinya','https://drive.google.com/cv/user-16','https://github.com/fina-maharani',          'accepted','2024-04-20 10:00:00'),
('pend-07','rek-07','user-17','Pengalaman analisis data cocok untuk riset ini',              'https://drive.google.com/cv/user-17','https://portfolio.gilang-saputra.com',      'accepted','2024-05-05 10:00:00'),
('pend-08','rek-08','user-18','Ingin berkontribusi di projek backend skala enterprise',      'https://drive.google.com/cv/user-18','https://github.com/hana-puspita',           'accepted','2024-05-20 10:00:00'),
('pend-09','rek-09','user-19','Antusias di bidang robotika dan otomasi',                     'https://drive.google.com/cv/user-19','https://github.com/ilham-maulana',          'accepted','2024-06-05 10:00:00'),
('pend-10','rek-10','user-20','Pengembangan game adalah passion saya sejak SMA',             'https://drive.google.com/cv/user-20','https://portfolio.jasmine-putri.com',        'accepted','2024-06-20 10:00:00'),

-- rejected (rek-01..10, user-21..30)
('pend-11','rek-01','user-21','Ingin mencoba riset ML meski background Go',                  'https://drive.google.com/cv/user-21','https://github.com/kevin-hidayat',          'rejected','2024-02-18 10:00:00'),
('pend-12','rek-02','user-22','Berpengalaman di SQL dan ingin belajar full stack',           'https://drive.google.com/cv/user-22','https://portfolio.laila-nur.com',            'rejected','2024-03-03 10:00:00'),
('pend-13','rek-03','user-23','Tertarik lomba statistika meski jurusan jaringan',            'https://drive.google.com/cv/user-23','https://github.com/mario-setiawan',         'rejected','2024-03-13 10:00:00'),
('pend-14','rek-04','user-24','Ingin memperluas pengalaman ke IoT dari ML',                  'https://drive.google.com/cv/user-24','https://portfolio.nadia-fitri.com',          'rejected','2024-03-23 10:00:00'),
('pend-15','rek-05','user-25',NULL,                                                          'https://drive.google.com/cv/user-25','https://github.com/omar-abdullah',          'rejected','2024-04-08 10:00:00'),
('pend-16','rek-06','user-26','Ingin belajar cybersecurity meski background matematika',     'https://drive.google.com/cv/user-26','https://github.com/putri-ayu',              'rejected','2024-04-23 10:00:00'),
('pend-17','rek-07','user-27',NULL,                                                          'https://drive.google.com/cv/user-27','https://portfolio.rizki-ramadhan.com',      'rejected','2024-05-08 10:00:00'),
('pend-18','rek-08','user-28','DevOps background ingin coba projek backend kompetitif',      'https://drive.google.com/cv/user-28','https://github.com/sari-wulandari',         'rejected','2024-05-23 10:00:00'),
('pend-19','rek-09','user-29',NULL,                                                          'https://drive.google.com/cv/user-29','https://portfolio.taufik-hidayat.com',      'rejected','2024-06-08 10:00:00'),
('pend-20','rek-10','user-30','Tertarik riset game engine dari perspektif mobile dev',       'https://drive.google.com/cv/user-30','https://github.com/ulfa-rahmawati',         'rejected','2024-06-23 10:00:00'),

-- pending (rek-11..20, user-31..40)
('pend-21','rek-11','user-31','Pengalaman React Native relevan untuk NLP UI',                'https://drive.google.com/cv/user-31','https://github.com/vino-pratama',           'pending', '2024-07-10 10:00:00'),
('pend-22','rek-12','user-32','Siap ikut hackathon, tim sudah terbentuk',                   'https://drive.google.com/cv/user-32','https://portfolio.winda-sari.com',           'pending', '2024-07-25 10:00:00'),
('pend-23','rek-13','user-33',NULL,                                                          'https://drive.google.com/cv/user-33','https://github.com/xander-putra',           'pending', '2024-08-10 10:00:00'),
('pend-24','rek-14','user-34','Ingin belajar otomasi industri dari sisi software',           'https://drive.google.com/cv/user-34','https://portfolio.yuni-astuti.com',          'pending', '2024-08-25 10:00:00'),
('pend-25','rek-15','user-35','Pengalaman Kubernetes dan cloud relevan untuk MLOps',         'https://drive.google.com/cv/user-35','https://github.com/zaki-maulana',           'pending', '2024-09-10 10:00:00'),
('pend-26','rek-16','user-36',NULL,                                                          'https://drive.google.com/cv/user-36','https://github.com/adelia-putri',           'pending', '2024-09-25 10:00:00'),
('pend-27','rek-17','user-37','Raspberry Pi background membantu memahami sistem tertanam',   'https://drive.google.com/cv/user-37','https://portfolio.bagas-nugraha.com',        'pending', '2024-10-10 10:00:00'),
('pend-28','rek-18','user-38','iOS dev ingin coba olimpiade matematika',                     'https://drive.google.com/cv/user-38','https://github.com/clara-octaviani',        'pending', '2024-10-25 10:00:00'),
('pend-29','rek-19','user-39',NULL,                                                          'https://drive.google.com/cv/user-39','https://portfolio.dimas-saputro.com',        'pending', '2024-11-10 10:00:00'),
('pend-30','rek-20','user-40','Unity experience mendukung pengembangan VR',                  'https://drive.google.com/cv/user-40','https://github.com/erni-susanti',           'pending', '2024-11-25 10:00:00'),

-- accepted (rek-21..30, user-11..20)
('pend-31','rek-21','user-11','Data science adalah bidang utama saya, siap kompetisi',       'https://drive.google.com/cv/user-11','https://github.com/aldi-firmansyah',        'accepted','2024-12-10 10:00:00'),
('pend-32','rek-22','user-12','UI/UX dan cloud saling melengkapi untuk produk',              'https://drive.google.com/cv/user-12','https://portfolio.bella-octavia.com',        'accepted','2024-12-25 10:00:00'),
('pend-33','rek-23','user-13','C++ background membantu image processing performa tinggi',    'https://drive.google.com/cv/user-13','https://github.com/cahya-pratama',           'accepted','2025-01-10 10:00:00'),
('pend-34','rek-24','user-14','Ingin menguasai desain PCB sebagai pelengkap skill IoT',      'https://drive.google.com/cv/user-14','https://portfolio.dewi-anggraeni.com',       'accepted','2025-01-25 10:00:00'),
('pend-35','rek-25','user-15','Ingin transisi dari web ke deep learning research',           'https://drive.google.com/cv/user-15','https://github.com/eka-prasetyo',            'accepted','2025-02-10 10:00:00'),
('pend-36','rek-26','user-16','Statistika dan DevOps bisa saling mendukung di MLOps',        'https://drive.google.com/cv/user-16','https://github.com/fina-maharani',           'accepted','2025-02-25 10:00:00'),
('pend-37','rek-27','user-17','Pengalaman analisis data relevan untuk business case',        'https://drive.google.com/cv/user-17','https://portfolio.gilang-saputra.com',       'accepted','2025-03-10 10:00:00'),
('pend-38','rek-28','user-18','TensorFlow background mendukung implementasi kriptografi',    'https://drive.google.com/cv/user-18','https://github.com/hana-puspita',            'accepted','2025-03-25 10:00:00'),
('pend-39','rek-29','user-19','Node.js dan ROS sama-sama event-driven, menarik',             'https://drive.google.com/cv/user-19','https://github.com/ilham-maulana',           'accepted','2025-04-10 10:00:00'),
('pend-40','rek-30','user-20','Flutter/Dart dan game jam adalah kombinasi ideal',            'https://drive.google.com/cv/user-20','https://portfolio.jasmine-putri.com',         'accepted','2025-04-25 10:00:00');


-- ====================================================================
-- 5. TIM_PARTICIPANTS (40 rows)
-- Setiap tim mendapat 1 peserta utama.
-- Peserta = mahasiswa yang status-nya accepted atau ditambahkan langsung.
--   tp-01..10  → tim-01..10,  user-11..20 (dari pend-01..10 accepted)
--   tp-11..20  → tim-11..20,  user-21..30 (ditambahkan langsung)
--   tp-21..30  → tim-21..30,  user-11..20 (dari pend-31..40 accepted)
--   tp-31..40  → tim-31..40,  user-31..40 (ditambahkan langsung)
-- ====================================================================
INSERT INTO tim_participants (id, tim_id, user_id, member_ke) VALUES
('tp-01','tim-01','user-11',1),
('tp-02','tim-02','user-12',1),
('tp-03','tim-03','user-13',1),
('tp-04','tim-04','user-14',1),
('tp-05','tim-05','user-15',1),
('tp-06','tim-06','user-16',1),
('tp-07','tim-07','user-17',1),
('tp-08','tim-08','user-18',1),
('tp-09','tim-09','user-19',1),
('tp-10','tim-10','user-20',1),
('tp-11','tim-11','user-21',1),
('tp-12','tim-12','user-22',1),
('tp-13','tim-13','user-23',1),
('tp-14','tim-14','user-24',1),
('tp-15','tim-15','user-25',1),
('tp-16','tim-16','user-26',1),
('tp-17','tim-17','user-27',1),
('tp-18','tim-18','user-28',1),
('tp-19','tim-19','user-29',1),
('tp-20','tim-20','user-30',1),
('tp-21','tim-21','user-11',1),
('tp-22','tim-22','user-12',1),
('tp-23','tim-23','user-13',1),
('tp-24','tim-24','user-14',1),
('tp-25','tim-25','user-15',1),
('tp-26','tim-26','user-16',1),
('tp-27','tim-27','user-17',1),
('tp-28','tim-28','user-18',1),
('tp-29','tim-29','user-19',1),
('tp-30','tim-30','user-20',1),
('tp-31','tim-31','user-31',1),
('tp-32','tim-32','user-32',1),
('tp-33','tim-33','user-33',1),
('tp-34','tim-34','user-34',1),
('tp-35','tim-35','user-35',1),
('tp-36','tim-36','user-36',1),
('tp-37','tim-37','user-37',1),
('tp-38','tim-38','user-38',1),
('tp-39','tim-39','user-39',1),
('tp-40','tim-40','user-40',1);


-- ====================================================================
-- 6. HISTORIES (40 rows)
-- Reviewer = dosen pembuat rekrutmen (sesuai tim → rekrutmen → user).
-- User     = peserta tim tersebut.
-- Pola reviewer mengikuti siklus user-01..10 sama seperti rekrutmen.
-- ====================================================================
INSERT INTO histories (id, user_id, reviewer_user_id, tim_id, rating, deskripsi, created_at) VALUES
('hist-01','user-11','user-01','tim-01',5,'Aldi menunjukkan kemampuan ML yang sangat baik, implementasi model tepat waktu','2024-07-25 10:00:00'),
('hist-02','user-12','user-02','tim-02',4,'Bella mampu deliver fitur frontend dan backend, perlu peningkatan di testing','2024-08-20 10:00:00'),
('hist-03','user-13','user-03','tim-03',5,'Cahya tampil luar biasa di lomba, kontribusi signifikan pada kemenangan tim','2024-06-15 10:00:00'),
('hist-04','user-14','user-04','tim-04',4,'Dewi memahami sensor IoT dengan cepat, dokumentasi perlu diperbaiki','2024-10-05 10:00:00'),
('hist-05','user-15','user-05','tim-05',3,'Eka perlu memperdalam Computer Vision, progress lambat di bulan pertama','2024-11-10 10:00:00'),
('hist-06','user-16','user-06','tim-06',5,'Fina sangat antusias di CTF, berhasil solve 3 challenge tingkat medium','2024-07-10 10:00:00'),
('hist-07','user-17','user-07','tim-07',4,'Gilang analisis data sangat tajam, presentasi perlu lebih terstruktur','2024-12-05 10:00:00'),
('hist-08','user-18','user-08','tim-08',5,'Hana menulis kode backend yang bersih dan teroptimasi','2025-01-05 10:00:00'),
('hist-09','user-19','user-09','tim-09',3,'Ilham perlu mempelajari lebih dalam ROS, komunikasi tim perlu ditingkatkan','2024-09-10 10:00:00'),
('hist-10','user-20','user-10','tim-10',4,'Jasmine mendesain game mechanics yang kreatif, timeline tepat sasaran','2025-01-10 10:00:00'),
('hist-11','user-21','user-01','tim-11',4,'Kevin mengimplementasi pipeline NLP dengan pendekatan yang efisien','2025-02-05 10:00:00'),
('hist-12','user-22','user-02','tim-12',5,'Laila menjadi backbone tim hackathon, database design sangat efisien','2024-10-15 10:00:00'),
('hist-13','user-23','user-03','tim-13',3,'Mario butuh bimbingan lebih dalam analisis statistik, coding bagus','2025-03-10 10:00:00'),
('hist-14','user-24','user-04','tim-14',5,'Nadia adaptasi ke PLC lebih cepat dari ekspektasi, kerja keras terbayar','2025-04-10 10:00:00'),
('hist-15','user-25','user-05','tim-15',4,'Omar membawa perspektif blockchain unik ke kompetisi AI','2024-12-15 10:00:00'),
('hist-16','user-26','user-06','tim-16',3,'Putri belum familiar penetration testing, butuh waktu lebih lama onboarding','2025-04-15 10:00:00'),
('hist-17','user-27','user-07','tim-17',5,'Rizki membantu automasi laporan bisnis yang menghemat waktu tim','2025-05-10 10:00:00'),
('hist-18','user-28','user-08','tim-18',4,'Sari mempersiapkan olimpiade dengan serius, latihan soal konsisten','2025-01-15 10:00:00'),
('hist-19','user-29','user-09','tim-19',4,'Taufik memahami sistem kontrol lebih cepat berkat background SAP','2025-06-10 10:00:00'),
('hist-20','user-30','user-10','tim-20',5,'Ulfa VR environment yang dibangun sangat immersive dan responsif','2025-06-15 10:00:00'),
('hist-21','user-11','user-01','tim-21',5,'Aldi memimpin analisis dataset kompetisi dengan sangat baik','2025-03-10 10:00:00'),
('hist-22','user-12','user-02','tim-22',4,'Bella membantu desain arsitektur cloud yang scalable','2025-07-10 10:00:00'),
('hist-23','user-13','user-03','tim-23',4,'Cahya mengoptimasi image processing pipeline hingga 3x lebih cepat','2025-08-05 10:00:00'),
('hist-24','user-14','user-04','tim-24',3,'Dewi baru belajar PCB, layout akhir masih perlu revisi minor','2025-03-05 10:00:00'),
('hist-25','user-15','user-05','tim-25',5,'Eka transisi ke deep learning sangat cepat, paper siap submit','2025-09-05 10:00:00'),
('hist-26','user-16','user-06','tim-26',4,'Fina pipeline CI/CD berjalan stabil, zero downtime sejak deploy','2025-09-10 10:00:00'),
('hist-27','user-17','user-07','tim-27',4,'Gilang presentasi business case di final mendapat pujian juri','2025-06-05 10:00:00'),
('hist-28','user-18','user-08','tim-28',5,'Hana implementasi algoritma RSA custom sangat tepat dan efisien','2025-10-10 10:00:00'),
('hist-29','user-19','user-09','tim-29',4,'Ilham robot navigasi otonom berhasil melewati seluruh obstacle test','2025-11-05 10:00:00'),
('hist-30','user-20','user-10','tim-30',3,'Jasmine game jam produk jadi namun polish UI masih kurang waktu','2025-07-05 10:00:00'),
('hist-31','user-31','user-01','tim-31',5,'Vino convergence RL agent sangat cepat, hypertuning sangat baik','2025-12-05 10:00:00'),
('hist-32','user-32','user-02','tim-32',4,'Winda desain Kafka topic yang bersih, latensi rendah di load test','2025-12-10 10:00:00'),
('hist-33','user-33','user-03','tim-33',3,'Xander perlu memperdalam probabilitas untuk olimpiade tingkat lanjut','2025-09-15 10:00:00'),
('hist-34','user-34','user-04','tim-34',5,'Yuni sensor dashboard smart grid yang dibangun real-time dan akurat','2026-01-05 10:00:00'),
('hist-35','user-35','user-05','tim-35',4,'Zaki MLOps pipeline Kubernetes stabil, model drift detection aktif','2026-02-05 10:00:00'),
('hist-36','user-36','user-06','tim-36',4,'Adelia MATLAB background ternyata sangat relevan untuk CTF forensic','2025-10-15 10:00:00'),
('hist-37','user-37','user-07','tim-37',5,'Bagas analisis ERP sangat komprehensif, rekomendasi diterima klien','2026-03-05 10:00:00'),
('hist-38','user-38','user-08','tim-38',4,'Clara implementasi lexer compiler Swift-style berjalan dengan benar','2026-03-10 10:00:00'),
('hist-39','user-39','user-09','tim-39',3,'Dimas penerbangan otonom masih sering off-track, perlu tuning PID','2025-12-15 10:00:00'),
('hist-40','user-40','user-10','tim-40',5,'Erni XR experience sangat smooth di headset Quest 3, luar biasa','2026-04-05 10:00:00');


-- ====================================================================
-- 7. BOOKMARKS (40 rows)
-- User men-bookmark rekrutmen yang menarik baginya.
--   bm-01..10  → user-11..20  bookmark rek-01..10
--   bm-11..20  → user-21..30  bookmark rek-11..20
--   bm-21..30  → user-31..40  bookmark rek-21..30
--   bm-31..40  → user-11..20  bookmark rek-31..40
-- Tidak ada pasangan (user_id, rekrutmen_id) yang duplikat.
-- ====================================================================
INSERT INTO bookmarks (id, rekrutmen_id, user_id) VALUES
('bm-01','rek-01','user-11'),
('bm-02','rek-02','user-12'),
('bm-03','rek-03','user-13'),
('bm-04','rek-04','user-14'),
('bm-05','rek-05','user-15'),
('bm-06','rek-06','user-16'),
('bm-07','rek-07','user-17'),
('bm-08','rek-08','user-18'),
('bm-09','rek-09','user-19'),
('bm-10','rek-10','user-20'),
('bm-11','rek-11','user-21'),
('bm-12','rek-12','user-22'),
('bm-13','rek-13','user-23'),
('bm-14','rek-14','user-24'),
('bm-15','rek-15','user-25'),
('bm-16','rek-16','user-26'),
('bm-17','rek-17','user-27'),
('bm-18','rek-18','user-28'),
('bm-19','rek-19','user-29'),
('bm-20','rek-20','user-30'),
('bm-21','rek-21','user-31'),
('bm-22','rek-22','user-32'),
('bm-23','rek-23','user-33'),
('bm-24','rek-24','user-34'),
('bm-25','rek-25','user-35'),
('bm-26','rek-26','user-36'),
('bm-27','rek-27','user-37'),
('bm-28','rek-28','user-38'),
('bm-29','rek-29','user-39'),
('bm-30','rek-30','user-40'),
('bm-31','rek-31','user-11'),
('bm-32','rek-32','user-12'),
('bm-33','rek-33','user-13'),
('bm-34','rek-34','user-14'),
('bm-35','rek-35','user-15'),
('bm-36','rek-36','user-16'),
('bm-37','rek-37','user-17'),
('bm-38','rek-38','user-18'),
('bm-39','rek-39','user-19'),
('bm-40','rek-40','user-20');


-- ====================================================================
-- 8. SETTING_DRIVES (40 rows)
-- ID di-generate otomatis oleh uuid_generate_v4().
-- Siklus bulan Januari–Desember; hari 1–31.
-- ====================================================================
INSERT INTO setting_drives (id, type_id, month, day, month_url, day_url) VALUES
(uuid_generate_v4(),'TYPE-A','January',   1, 'https://drive.google.com/folder/month/January-2024',   'https://drive.google.com/folder/day/2024-01-01'),
(uuid_generate_v4(),'TYPE-B','February',  2, 'https://drive.google.com/folder/month/February-2024',  'https://drive.google.com/folder/day/2024-02-02'),
(uuid_generate_v4(),'TYPE-C','March',     5, 'https://drive.google.com/folder/month/March-2024',     'https://drive.google.com/folder/day/2024-03-05'),
(uuid_generate_v4(),'TYPE-A','April',     7, 'https://drive.google.com/folder/month/April-2024',     'https://drive.google.com/folder/day/2024-04-07'),
(uuid_generate_v4(),'TYPE-B','May',      10, 'https://drive.google.com/folder/month/May-2024',       'https://drive.google.com/folder/day/2024-05-10'),
(uuid_generate_v4(),'TYPE-C','June',     12, 'https://drive.google.com/folder/month/June-2024',      'https://drive.google.com/folder/day/2024-06-12'),
(uuid_generate_v4(),'TYPE-A','July',     15, 'https://drive.google.com/folder/month/July-2024',      'https://drive.google.com/folder/day/2024-07-15'),
(uuid_generate_v4(),'TYPE-B','August',   17, 'https://drive.google.com/folder/month/August-2024',    'https://drive.google.com/folder/day/2024-08-17'),
(uuid_generate_v4(),'TYPE-C','September',20, 'https://drive.google.com/folder/month/September-2024', 'https://drive.google.com/folder/day/2024-09-20'),
(uuid_generate_v4(),'TYPE-A','October',  22, 'https://drive.google.com/folder/month/October-2024',   'https://drive.google.com/folder/day/2024-10-22'),
(uuid_generate_v4(),'TYPE-B','November', 25, 'https://drive.google.com/folder/month/November-2024',  'https://drive.google.com/folder/day/2024-11-25'),
(uuid_generate_v4(),'TYPE-C','December', 28, 'https://drive.google.com/folder/month/December-2024',  'https://drive.google.com/folder/day/2024-12-28'),
(uuid_generate_v4(),'TYPE-A','January',   3, 'https://drive.google.com/folder/month/January-2025',   'https://drive.google.com/folder/day/2025-01-03'),
(uuid_generate_v4(),'TYPE-B','February',  6, 'https://drive.google.com/folder/month/February-2025',  'https://drive.google.com/folder/day/2025-02-06'),
(uuid_generate_v4(),'TYPE-C','March',     9, 'https://drive.google.com/folder/month/March-2025',     'https://drive.google.com/folder/day/2025-03-09'),
(uuid_generate_v4(),'TYPE-A','April',    11, 'https://drive.google.com/folder/month/April-2025',     'https://drive.google.com/folder/day/2025-04-11'),
(uuid_generate_v4(),'TYPE-B','May',      14, 'https://drive.google.com/folder/month/May-2025',       'https://drive.google.com/folder/day/2025-05-14'),
(uuid_generate_v4(),'TYPE-C','June',     16, 'https://drive.google.com/folder/month/June-2025',      'https://drive.google.com/folder/day/2025-06-16'),
(uuid_generate_v4(),'TYPE-A','July',     18, 'https://drive.google.com/folder/month/July-2025',      'https://drive.google.com/folder/day/2025-07-18'),
(uuid_generate_v4(),'TYPE-B','August',   21, 'https://drive.google.com/folder/month/August-2025',    'https://drive.google.com/folder/day/2025-08-21'),
(uuid_generate_v4(),'TYPE-C','September',23, 'https://drive.google.com/folder/month/September-2025', 'https://drive.google.com/folder/day/2025-09-23'),
(uuid_generate_v4(),'TYPE-A','October',  26, 'https://drive.google.com/folder/month/October-2025',   'https://drive.google.com/folder/day/2025-10-26'),
(uuid_generate_v4(),'TYPE-B','November', 29, 'https://drive.google.com/folder/month/November-2025',  'https://drive.google.com/folder/day/2025-11-29'),
(uuid_generate_v4(),'TYPE-C','December', 30, 'https://drive.google.com/folder/month/December-2025',  'https://drive.google.com/folder/day/2025-12-30'),
(uuid_generate_v4(),'TYPE-A','January',   8, 'https://drive.google.com/folder/month/January-2026',   'https://drive.google.com/folder/day/2026-01-08'),
(uuid_generate_v4(),'TYPE-B','February', 13, 'https://drive.google.com/folder/month/February-2026',  'https://drive.google.com/folder/day/2026-02-13'),
(uuid_generate_v4(),'TYPE-C','March',    19, 'https://drive.google.com/folder/month/March-2026',     'https://drive.google.com/folder/day/2026-03-19'),
(uuid_generate_v4(),'TYPE-A','April',    24, 'https://drive.google.com/folder/month/April-2026',     'https://drive.google.com/folder/day/2026-04-24'),
(uuid_generate_v4(),'TYPE-B','May',       4, 'https://drive.google.com/folder/month/May-2026',       'https://drive.google.com/folder/day/2026-05-04'),
(uuid_generate_v4(),'TYPE-C','June',     27, 'https://drive.google.com/folder/month/June-2026',      'https://drive.google.com/folder/day/2026-06-27'),
(uuid_generate_v4(),'TYPE-A','July',     31, 'https://drive.google.com/folder/month/July-2026',      'https://drive.google.com/folder/day/2026-07-31'),
(uuid_generate_v4(),'TYPE-B','August',    2, 'https://drive.google.com/folder/month/August-2026',    'https://drive.google.com/folder/day/2026-08-02'),
(uuid_generate_v4(),'TYPE-C','September', 5, 'https://drive.google.com/folder/month/September-2026', 'https://drive.google.com/folder/day/2026-09-05'),
(uuid_generate_v4(),'TYPE-A','October',  11, 'https://drive.google.com/folder/month/October-2026',   'https://drive.google.com/folder/day/2026-10-11'),
(uuid_generate_v4(),'TYPE-B','November', 14, 'https://drive.google.com/folder/month/November-2026',  'https://drive.google.com/folder/day/2026-11-14'),
(uuid_generate_v4(),'TYPE-C','December', 20, 'https://drive.google.com/folder/month/December-2026',  'https://drive.google.com/folder/day/2026-12-20'),
(uuid_generate_v4(),'TYPE-A','January',  16, 'https://drive.google.com/folder/month/January-2027',   'https://drive.google.com/folder/day/2027-01-16'),
(uuid_generate_v4(),'TYPE-B','February', 22, 'https://drive.google.com/folder/month/February-2027',  'https://drive.google.com/folder/day/2027-02-22'),
(uuid_generate_v4(),'TYPE-C','March',    28, 'https://drive.google.com/folder/month/March-2027',     'https://drive.google.com/folder/day/2027-03-28'),
(uuid_generate_v4(),'TYPE-A','April',     9, 'https://drive.google.com/folder/month/April-2027',     'https://drive.google.com/folder/day/2027-04-09');