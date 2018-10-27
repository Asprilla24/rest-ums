CREATE TABLE AppRoles (
    ID VARCHAR(100) PRIMARY KEY,
    DisplayName VARCHAR(100) NOT NULL,
    IsDefault BIT NOT NULL,
    TenantId VARCHAR(100),
    Description VARCHAR(100),
    LastUpdated DATETIME NOT NULL
)

CREATE INDEX IDX_AppRoles
ON AppRoles (ID, LastUpdated)