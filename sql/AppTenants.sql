CREATE TABLE AppTenants (
    ID VARCHAR(100) PRIMARY KEY,
    IsActive BIT NOT NULL,
    Name VARCHAR(100) NOT NULL,
    LastUpdated DATETIME NOT NULL
)

CREATE INDEX IDX_AppTenants
ON AppTenants (ID, LastUpdated)