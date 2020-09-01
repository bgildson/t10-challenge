/**
 * SUPERUSER USERS
 */
INSERT INTO users
	(id, name, email, hashed_password, role, created_at)
VALUES
	(uuid_in(md5(random()::text || clock_timestamp()::text)::cstring),
	 'Gildson Bezerra',
	 'gildson@email.com',
	 '$2a$10$maUTf7nv6nZ5dyYzMwHMAu9fGNlEf3GYSb7a8f46WpTqVwWjG8/.m', -- 123456
	 'superuser',
	 NOW());

INSERT INTO users
	(id, name, email, hashed_password, role, created_at)
VALUES
	(uuid_in(md5(random()::text || clock_timestamp()::text)::cstring),
	 'administrator',
	 'admin@email.com',
	 '$2a$10$maUTf7nv6nZ5dyYzMwHMAu9fGNlEf3GYSb7a8f46WpTqVwWjG8/.m', -- 123456
	 'superuser',
	 NOW());


/**
 * EXTERNALAPP USERS
 */
INSERT INTO users
	(id, name, email, hashed_password, role, created_at) 
VALUES
	(uuid_in(md5(random()::text || clock_timestamp()::text)::cstring),
	 'Green Bank',
	 'greenbank@email.com',
	 '$2a$10$maUTf7nv6nZ5dyYzMwHMAu9fGNlEf3GYSb7a8f46WpTqVwWjG8/.m', -- 123456
	 'externalapp',
	 NOW());

INSERT INTO users
	(id, name, email, hashed_password, role, created_at) 
VALUES
	(uuid_in(md5(random()::text || clock_timestamp()::text)::cstring),
	 'Blue Bank',
	 'bluebank@email.com',
	 '$2a$10$maUTf7nv6nZ5dyYzMwHMAu9fGNlEf3GYSb7a8f46WpTqVwWjG8/.m', -- 123456
	 'externalapp',
	 NOW());

INSERT INTO users
	(id, name, email, hashed_password, role, created_at) 
VALUES
	(uuid_in(md5(random()::text || clock_timestamp()::text)::cstring),
	 'Purple Bank',
	 'purplebank@email.com',
	 '$2a$10$maUTf7nv6nZ5dyYzMwHMAu9fGNlEf3GYSb7a8f46WpTqVwWjG8/.m', -- 123456
	 'externalapp',
	 NOW());
