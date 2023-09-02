create database restaurant;

CREATE TABLE "users" (
  "id" uuid PRIMARY KEY NOT NULL,
  "email" varchar NOT NULL UNIQUE,
  "password" varchar NOT NULL,
  "full_name" varchar NULL,
  "phone" varchar NULL,
  "profile_picture" varchar NULL,
  "address" varchar NULL,
  "role" smallint NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "coupons" (
  "id" uuid PRIMARY KEY NOT NULL,
  "code" varchar NOT NULL UNIQUE,
  "discount" int NOT NULL,
  "valid_until" timestamp NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" timestamp NULL
);

CREATE TABLE "user_coupons" (
  "user_id" uuid NOT NULL,
  "coupon_id" uuid NOT NULL,
  "expired_at" timestamp NOT NULL,
  "qty" int NOT NULL,
  PRIMARY KEY ("user_id", "coupon_id")
);

CREATE TABLE "categories" (
  "id" SERIAL PRIMARY KEY NOT NULL,
  "name" varchar NOT NULL
);

CREATE TABLE "menus" (
  "id" SERIAL PRIMARY KEY NOT NULL,
  "name" varchar NOT NULL,
  "price" int NOT NULL,
  "photo" varchar NOT NULL,
  "category_id" int NOT NULL,
  "rating" int NULL,
  "total_review" int NULL,
  "description" varchar NULL,
  "is_available" boolean NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" timestamp NULL
);

CREATE TABLE "user_favorites" (
  "user_id" uuid NOT NULL,
  "menu_id" int NOT NULL,
  PRIMARY KEY ("user_id", "menu_id")
);

CREATE TABLE "payments" (
  "id" SERIAL PRIMARY KEY NOT NULL,
  "description" varchar NOT NULL
);

CREATE TABLE "orders" (
  "id" SERIAL PRIMARY KEY NOT NULL,
  "user_id" uuid NOT NULL,
  "coupon_id" uuid NULL,
  "payment_id" int NOT NULL,
  "status" varchar NOT NULL,
  "subtotal" int NOT NULL,
  "total_price" int NOT NULL,
  "order_date" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "order_details" (
  "id" SERIAL NOT NULL PRIMARY KEY,
  "menu_id" int NOT NULL,
  "order_id" int NOT NULL,
  "qty" int NOT NULL,
  "option_id" INT NULL
);

CREATE TABLE "reviews" (
  "id" SERIAL PRIMARY KEY NOT NULL,
  "user_id" uuid NOT NULL,
  "menu_id" int NOT NULL,
  "description" varchar NOT NULL,
  "rating" int NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "menu_options" (
  "id" SERIAL PRIMARY KEY NOT NULL,
  "menu_id" INT NOT NULL,
  "name" varchar NOT NULL,
  "price" int NOT NULL
);

CREATE TABLE "questions" (
  "id" SERIAL PRIMARY KEY NOT NULL,
  "question" varchar NOT NULL,
  "correct_answer" varchar NOT NULL,
  "option_one" varchar NOT NULL,
  "option_two" varchar NOT NULL,
  "option_three" varchar NOT NULL,
  "option_four" varchar NOT NULL
);

CREATE TABLE "games" (
  "id" SERIAL PRIMARY KEY NOT NULL,
  "user_id" uuid NOT NULL,
  "score" int NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "leaderboards" (
  "id" SERIAL PRIMARY KEY NOT NULL,
  "user_id" uuid NOT NULL,
  "accumulated_score" int NOT NULL DEFAULT 0,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "promotions" (
	id serial PRIMARY KEY NOT NULL,
	menu_id int NOT NULL,
	"name" varchar NOT NULL,
	discount int NOT NULL,
	created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at timestamp NULL
);


INSERT INTO users (id,email,"password","role") VALUES
	('3af191c4-6792-4eb2-902b-cc020c0e2f2e','admin@email.com','$2a$04$LoG0NFPuaobuUM0X9j4PGOU8PNxlcdaO7CHckjOcNPPotR6EKNWQq',0),
	('2034cc93-3b40-4dc1-821a-eeb80560321f','aldo@test.com','$2a$04$LoG0NFPuaobuUM0X9j4PGOU8PNxlcdaO7CHckjOcNPPotR6EKNWQq',1),
	('c3364a20-0dba-43a0-ac4d-ce2af4e14889','aldoo@test.com','$2a$04$LoG0NFPuaobuUM0X9j4PGOU8PNxlcdaO7CHckjOcNPPotR6EKNWQq',1),
	('51930fb7-7687-43c6-bcab-5bb852f0caef','test@test.com','$2a$04$LoG0NFPuaobuUM0X9j4PGOU8PNxlcdaO7CHckjOcNPPotR6EKNWQq',1);
	
INSERT INTO categories ("name") VALUES 
	('Beverage'), ('Dessert'), ('Fast Food'),
	('Chicken'), ('Western'), ('Asian'), ('Noodles');

INSERT INTO menus ("name", price, photo, category_id, rating, total_review, is_available) VALUES 
	('Sweet Tea', 10000, 'https://res.cloudinary.com/dcdexrr4n/image/upload/v1670229226/pexels-barbara-webb-792613_w5xjyr.jpg', 1, 0, 0, true),
	('Coke', 8000, 'https://res.cloudinary.com/dcdexrr4n/image/upload/v1670229619/pexels-pixabay-50593_e61f55.jpg', 1, 0, 0, true),
	('Coffee', 7000, 'https://res.cloudinary.com/dcdexrr4n/image/upload/v1670229622/pexels-tirachard-kumtanom-544113_irguad.jpg', 1, 0, 0, true),
	('Mineral Water', 5000, 'https://res.cloudinary.com/dcdexrr4n/image/upload/v1670229629/pexels-francesco-paggiaro-593099_gmwrjm.jpg', 1, 0, 0, false),
	('Grilled Gyoza', 20000, '', 6, 0, 0, false),
	('Siomay', 3000, '', 6, 0, 0, true),
	('Mendoan', 2000, '', 6, 0, 0, true),
	('Salad', 10000, 'https://res.cloudinary.com/dcdexrr4n/image/upload/v1670229633/pexels-roman-odintsov-4871111_r7knbo.jpg', 2, 0, 0, true),
	('Bakmi', 22000, '', 7, 0, 0, true),
	('Chicken Satay', 25000, 'https://res.cloudinary.com/dcdexrr4n/image/upload/v1670229635/pexels-mouktik-joshi-9646858_lieubt.jpg', 4, 0, 0, true),
	('Seafood Fried Rice', 30000, 'https://res.cloudinary.com/dcdexrr4n/image/upload/v1670229822/pexels-zn_s-food_natureart-8992927_izvxme.jpg', 6, 0, 0, true),
	('Kebab', 23000, 'https://res.cloudinary.com/dcdexrr4n/image/upload/v1670229826/pexels-dhiraj-jain-12737663_aqnauo.jpg', 5, 0, 0, true),
	('Burger', 33000, 'https://res.cloudinary.com/dcdexrr4n/image/upload/v1670229829/pexels-ata-ebem-10831651_ecyofh.jpg', 3, 0, 0, true),
	('Special Fried Rice', 20000, 'https://res.cloudinary.com/dcdexrr4n/image/upload/v1670229822/pexels-zn_s-food_natureart-8992927_izvxme.jpg', 6, 0, 0, true),
	('Fettucine Carbonara', 27000, 'https://res.cloudinary.com/dcdexrr4n/image/upload/v1670229830/pexels-ali-nafezarefi-4161714_oyxvhx.jpg', 7, 0, 0, true),
	('Banana Split', 17000, '', 2, 0, 0, true),
	('Mochi',5000, '', 2, 0, 0, true),
	('Ice Cream',13000, 'https://res.cloudinary.com/dcdexrr4n/image/upload/v1670229835/pexels-%E5%8A%89%E8%8A%B7%E5%AE%89-11719086_plmhgj.jpg', 2, 0, 0, true);

INSERT INTO coupons (id, code, valid_until, discount) VALUES 
	('f1e35304-4b89-4301-8ec4-b5d5f81a5495', 'FRIYAY50', '2022-12-30', 50000),
	('cc6fe9c1-a1aa-46b9-8c7c-77b662c55395', 'MONDAYPAYDAY', '2022-12-30', 100000),
	('4a5985d7-6897-41c6-96ad-d27c85c9cda0', 'DECEMBERDEAL20', '2022-12-30', 20000);

INSERT INTO user_coupons (user_id, coupon_id, expired_at, qty) VALUES 
	('2034cc93-3b40-4dc1-821a-eeb80560321f', 'cc6fe9c1-a1aa-46b9-8c7c-77b662c55395', '2022-12-30', 1),
	('2034cc93-3b40-4dc1-821a-eeb80560321f', '4a5985d7-6897-41c6-96ad-d27c85c9cda0', '2022-12-30', 1),
	('c3364a20-0dba-43a0-ac4d-ce2af4e14889', '4a5985d7-6897-41c6-96ad-d27c85c9cda0', '2022-12-30', 1);

INSERT INTO payments (description) VALUES 
	('DANA'),
	('ShopeePay');

INSERT INTO menu_options (menu_id, "name", price) VALUES
	(18, 'Vanilla With Cone', 0),
	(18, 'Vanilla With Cup', 1000),
	(18, 'Chocolate With Cone', 0),
	(18, 'Chocolate With Cup', 1000);

INSERT INTO questions ("question", "correct_answer", option_one, option_two, option_three, option_four) VALUES 
	('What country in the world produces the most coffee?', 'Brazil', 'Brazil', 'Argentina', 'USA', 'Indonesia'),
	('What does the Japanese word Sayonara mean?', 'Goodbye', 'Welcome', 'Hai', 'Goodbye', 'How Are You?'),
	('What country is sushi from?', 'Japan', 'Vietnam', 'Thailand', 'Korea', 'Japan'),
	('What U.S. state is known as the Empire State?', 'New York State', 'Washington DC', 'New York State', 'Texas', 'Hollywood'),
	('In what country will one find the Leaning Tower of Pisa?', 'Italy', 'Portugal', 'Spain', 'Italy', 'England'),
	('What is the name of the fictional city Batman calls home?', 'Gotham', 'Bikini Bottom', 'Gotham', 'Konoha', 'Crypton'),
	('Which chess piece cannot move in a straight line?', 'Knight', 'Pawn', 'Bishop', 'Rook', 'Knight'),
	('Which world continent is the largest?', 'Asia', 'Africa', 'Europe', 'Asia', 'Antartica'),
	('On what continent is the country of Egypt?', 'Africa', 'Africa', 'Asia', 'South America', 'Europe'),
	('What is the name of the world’s largest desert?', 'Sahara', 'Arabian', 'Sahara', 'Gobi', 'Patagonian');


INSERT INTO promotions (menu_id, "name", discount) VALUES
	(18, 'Ice Cream You Scream', 5000),
	(12, 'Kebab Party', 10000);

ALTER TABLE "user_coupons" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
ALTER TABLE "user_coupons" ADD FOREIGN KEY ("coupon_id") REFERENCES "coupons" ("id");

ALTER TABLE "menus" ADD FOREIGN KEY ("category_id") REFERENCES "categories" ("id");

ALTER TABLE "user_favorites" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
ALTER TABLE "user_favorites" ADD FOREIGN KEY ("menu_id") REFERENCES "menus" ("id");

ALTER TABLE "orders" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
ALTER TABLE "orders" ADD FOREIGN KEY ("payment_id") REFERENCES "payments" ("id");

ALTER TABLE "order_details" ADD FOREIGN KEY ("menu_id") REFERENCES "menus" ("id");
ALTER TABLE "order_details" ADD FOREIGN KEY ("order_id") REFERENCES "orders" ("id");
ALTER TABLE "order_details" ADD FOREIGN KEY ("option_id") REFERENCES "menu_options" ("id");

ALTER TABLE "menu_options" ADD FOREIGN KEY ("menu_id") REFERENCES "menus" ("id");

ALTER TABLE "reviews" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
ALTER TABLE "reviews" ADD FOREIGN KEY ("menu_id") REFERENCES "menus" ("id");

ALTER TABLE "games" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "leaderboards" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "promotions" ADD FOREIGN KEY ("menu_id") REFERENCES "menus" ("id");



