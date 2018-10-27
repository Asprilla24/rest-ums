CREATE TABLE AppUsers (
    ID VARCHAR(100) PRIMARY KEY,
    AccessFailedCount INT NOT NULL,
    EmailAddress VARCHAR(100) UNIQUE NOT NULL,
    EmailConfirmationCode VARCHAR(20) NOT NULL,
    IsActive TINYINT(1) NOT NULL,
    IsEmailConfirmed TINYINT(1) NOT NULL,
    IsPhoneNumberConfirmed TINYINT(1) NOT NULL,
    LastLoginTime DATETIME NOT NULL,
    Password VARCHAR(200) NOT NULL,
    PasswordResetCode VARCHAR(50) NOT NULL,
    PhoneNumber VARCHAR(20) NOT NULL,
    FirstName VARCHAR(20) NOT NULL,
    LastName VARCHAR(20) NOT NULL,
    TenantId VARCHAR(100) NOT NULL,
    Username VARCHAR(100) UNIQUE NOT NULL,
    Gender VARCHAR(10) NOT NULL,
    LastUpdated DATETIME NOT NULL
)

CREATE INDEX IDX_AppUsers
ON AppUsers (ID, LastUpdated)

INSERT INTO AppUsers
VALUES (
    'Al241',
    0,
    'aldesenaasprilla@gmail.com',
    '',
    true,
    true,
    true,
    '2018-10-28 00:00:01',
    'AldeGanteng',
    '',
    '082213810518',
    'Alde',
    'Asprilla',
    '01',
    'Asprilla24',
    'MALE',
    '2018-10-29 20:06:00'
)