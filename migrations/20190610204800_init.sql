CREATE TABLE univers.region (
  id INT(11) NOT NULL AUTO_INCREMENT,
  region_name VARCHAR(256) NOT NULL,
  PRIMARY KEY (id)
)
  ENGINE innoDB
  CHARACTER SET utf8
  COLLATE utf8_general_ci;

CREATE TABLE univers.faculty (
  id INT(11) NOT NULL AUTO_INCREMENT,
  faculty_name VARCHAR(256) NOT NULL,
  PRIMARY KEY (id)
)
  ENGINE innoDB
  CHARACTER SET utf8
  COLLATE utf8_general_ci;

CREATE TABLE univers.rector (
  id INT(11) NOT NULL AUTO_INCREMENT,
  first_name VARCHAR(256) NOT NULL,
  family_name VARCHAR(256) NOT NULL,
  patronymic VARCHAR(256),
  PRIMARY KEY (id)
)
  ENGINE innoDB
  CHARACTER SET utf8
  COLLATE utf8_general_ci;

CREATE TABLE univers.condition (
  id INT(11) NOT NULL AUTO_INCREMENT,
  admission_condition TEXT NOT NULL,
  PRIMARY KEY (id)
)
  ENGINE innoDB
  CHARACTER SET utf8
  COLLATE utf8_general_ci;

CREATE TABLE univers.m2m_faculty_condition (
  id INT(11) NOT NULL AUTO_INCREMENT,
  faculty_id INT(11) NOT NULL,
  condition_id INT(11) NOT NULL,
  PRIMARY KEY (id)
)
  ENGINE innoDB
  CHARACTER SET utf8
  COLLATE utf8_general_ci;

CREATE TABLE univers.m2m_faculty_university (
  id INT(11) NOT NULL AUTO_INCREMENT,
  faculty_id INT(11) NOT NULL,
  university_id INT(11) NOT NULL,
  PRIMARY KEY (id)
)
  ENGINE innoDB
  CHARACTER SET utf8
  COLLATE utf8_general_ci;

CREATE TABLE univers.university (
  id INT(11) NOT NULL AUTO_INCREMENT,
  university_name VARCHAR(256) NOT NULL,
  address VARCHAR(256) NOT NULL,
  region_id INT(11) NOT NULL,
  status ENUM('state', 'commercial'),
  shape VARCHAR(256) NOT NULL,
  rector_id INT(11) NOT NULL,
  student_number INT(11) NOT NULL,
  PRIMARY KEY (id)
)
  ENGINE innoDB
  CHARACTER SET utf8
  COLLATE utf8_general_ci;

alter table univers.university
  add constraint university_region_id_fk
    foreign key (region_id) references region (id);

alter table univers.university
  add constraint university_rector_id_fk
    foreign key (rector_id) references rector (id);

alter table univers.m2m_faculty_university
  add constraint m2m_faculty_university_faculty_id_fk
    foreign key (faculty_id) references faculty (id);

alter table univers.m2m_faculty_university
  add constraint m2m_faculty_university_university_id_fk
    foreign key (university_id) references university (id);

alter table univers.m2m_faculty_condition
  add constraint m2m_faculty_condition_condition_id_fk
    foreign key (condition_id) references `condition` (id);

alter table univers.m2m_faculty_condition
  add constraint m2m_faculty_condition_faculty_id_fk
    foreign key (faculty_id) references faculty (id);






