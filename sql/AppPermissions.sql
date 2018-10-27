CREATE TABLE AppPermissions (
    ID VARCHAR(100) PRIMARY KEY,
    IsGranted BIT NOT NULL,
    Name VARCHAR(100) NOT NULL,
    RoleId VARCHAR(100) NOT NULL,
    UserId VARCHAR(100) NOT NULL,
    LastUpdated DATETIME NOT NULL
)

CREATE INDEX IDX_AppPermissions
ON AppPermissions (ID, LastUpdated)