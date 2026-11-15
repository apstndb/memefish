CREATE PROPERTY GRAPH G
NODE TABLES (`order`.`group` AS Person)
EDGE TABLES (
  `order`.`select` AS Relation
    SOURCE KEY (Src) REFERENCES Person (Id)
    DESTINATION KEY (Dst) REFERENCES Person (Id)
)
