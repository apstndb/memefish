CREATE PROPERTY GRAPH G
NODE TABLES (
  sch.People AS Person,
  sch.Accounts AS Account
)
EDGE TABLES (
  sch.Owns AS Own
    SOURCE KEY (PersonId) REFERENCES Person (Id)
    DESTINATION KEY (AccountId) REFERENCES Account (Id)
)
