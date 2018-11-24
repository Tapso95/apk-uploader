DROP TABLE type_application;
CREATE TABLE type_application (
	id_type_app INT(7) AUTO_INCREMENT,
	nom_type_app varchar(255),
 	description_type_app varchar(500)
	);

CREATE TABLE profil_utilisateur(
	id_profil INT(11) AUTO_INCREMENT,
	nom_profil varchar(255),
	code_profil varchar(255),
	description_profil varchar(500),
	date_creation_profil date
	);
CREATE TABLE evaluation_application(
	id_evaluation INT(11) AUTO_INCREMENT,
	evaluation DECIMAL(1,2),
	commentaire VARCHAR(500),
	date_evaluation DATETIME
);
CREATE TABLE categorie_application(
	id_categorie_app INT(11)AUTO_INCREMENT,
	nom_categorie_app VARCHAR(255),
	description_categorie_app VARCHAR(500)
);
CREATE TABLE admin(
	id_admin INT(11),
	nom_admin VARCHAR(255),
	prenom_admin VARCHAR(255),
	email_admin VARCHAR(100),
	password_admin VARCHAR(255)
);
ALTER TABLE detail_telechargement(
	id_detail_tele INT(11),
	date_tele DATETIME,
	id_application INT(11),
	id_utilisateur INT(11)

);

select nom_application,