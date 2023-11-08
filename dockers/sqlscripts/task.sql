BEGIN;
CREATE TABLE IF NOT EXISTS Employess (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    age DATE NOT NULL,
    sex VARCHAR(10) NOT NULL
);

INSERT INTO Employess (name, age, sex) VALUES ('Петров Петр Петрович', DATE '1976-02-03', 'М');
INSERT INTO Employess (name, age, sex) VALUES ('Иванов Иван Иванович', DATE '1997-04-15', 'М');
INSERT INTO Employess (name, age, sex) VALUES ('Яичко Лиза Сергеевна', DATE '2000-11-21', 'Ж');

CREATE TABLE IF NOT EXISTS Tool (
    id SERIAL PRIMARY KEY,
    descryption VARCHAR(100) NOT NULL
);

INSERT INTO Tool (descryption) VALUES ('Лопата');
INSERT INTO Tool (descryption) VALUES ('Молоток');
INSERT INTO Tool (descryption) VALUES ('Метла');

CREATE TABLE IF NOT EXISTS Users (
    id SERIAL PRIMARY KEY,
    login VARCHAR(100) NOT NULL,
    password VARCHAR(100) NOT NULL
);

INSERT INTO Users (login, password) VALUES ('user', 'user');

CREATE TABLE IF NOT EXISTS Tasks (
    id SERIAL PRIMARY KEY,
    descryption VARCHAR(100) NOT NULL,
    employeId INTEGER,
    ToolId INTEGER
);

ALTER TABLE IF EXISTS Tasks 
    ADD FOREIGN KEY (employeId)
    REFERENCES Employess (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;

ALTER TABLE IF EXISTS Tasks 
    ADD FOREIGN KEY (ToolId)
    REFERENCES Tool (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;
    
END;


