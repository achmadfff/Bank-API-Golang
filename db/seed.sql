INSERT INTO roles (id, name) VALUES ('1ca78238-8e40-4763-a20f-59b5b41791b1', 'Super Admin');
INSERT INTO roles (id, name) VALUES ('1ca78238-8e40-4763-a20f-59b5b41791b2', 'Customer');

INSERT INTO users (id, role_id, name, position, email, phone, encrypted_password) VALUES ('2ca78238-8e40-4763-a20f-59b5b41791b1', '1ca78238-8e40-4763-a20f-59b5b41791b1', 'Superadmin', 'Jendral', 'superadmin@mail.com', '081111111111', '$2a$04$GtAyqUr8G1jriNAWhYklQe8yJ2ei54i8Oaq3SyaLHmLT0rP/TZeWu');

INSERT into merchants(id,name,phone,address,balance)
VALUES ('bfdd663e-8167-45f3-91b2-b4cb11bf46c6','Ayam Geprek',08321389312,'Cibinong',0),
       ('87c43149-8757-4ecb-b99b-e60affd1eeb9','Teguk',098312983,'Cibinong',0),
       ('789c5424-e7ab-49da-9abd-0f1ff6a5c4ac','Es Teh Indonesia',09830217301,'Cibinong',0),
       ('71bb2e49-5c3a-4180-a567-ff1f41aa3d9f','Street Boba',09832109312,'Cibinong',0),
       ('d7049405-6a71-47a8-bf06-0c025d2e7a33','Haus Indoneisa',0812938103,'Cibinong',0)