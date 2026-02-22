GRAPH FinGraph
-- OPTIONAL MATCH
OPTIONAL MATCH (n:Person)
-- MATCH with Hints and multiple paths with top-level WHERE
MATCH @{join_method=hash_join} (src:Account)-[:Transfers]->(dst:Account), (src)<-[:Owns]-(p:Person) 
WHERE p.id > 0
-- Search Prefixes
MATCH ALL (p1:Person)-[:Knows]->+(f1:Person)
MATCH ANY (p2:Person)-[:Knows]->+(f2:Person)
MATCH ANY SHORTEST (p3:Person)-[:Knows]->+(f3:Person)
MATCH ANY CHEAPEST (p4:Person)-[:Knows]->+(f4:Person)
-- Path Modes (singular/plural)
MATCH WALK PATH (aw)-[ew]->(bw)
MATCH WALK PATHS (aws)-[ews]->(bws)
MATCH TRAIL PATH (atvar)-[et]->(btvar)
MATCH SIMPLE PATH (asvar)-[es]->(bsvar)
MATCH ACYCLIC PATH (aa)-[ea]->(ba)
-- Edge Directions and Fillers
MATCH (ad1)-[:Knows]-(bd1)
MATCH (ad2)<-[:Knows]-(bd2)
MATCH (ad3)-[:Knows]->(bd3)
MATCH (ad4)<-[:Knows]->(bd4)
MATCH (ad5)-[]-(bd5)
MATCH (ad6)-(bd6)
MATCH (ad7)<-(bd7)
MATCH (ad8)->(bd8)
MATCH (ad9)<->(bd9)
-- Label Expressions
MATCH (nl1:Person|Employee)
MATCH (nl2:Person&Employee)
MATCH (nl3:!Person)
MATCH (nl4:%)
MATCH (nl5:(Person|Employee)&!Student)
MATCH (nl6 IS Person)
MATCH (nl7:Person)
-- Element Filler with Variable, Properties, WHERE, and COST
MATCH path = (ne:Person {age: 20} WHERE ne.name = 'Alice')-[ee:Knows WHERE ee.since > 2020 COST ee.since]->(me)
-- Quantifiers
MATCH (aq1)-[:Knows]->* (bq1)
MATCH (aq2)-[:Knows]->+ (bq2)
MATCH (aq4)-[:Knows]->{2} (bq4)
MATCH (aq5)-[:Knows]->{2,} (bq5)
MATCH (aq6)-[:Knows]->{,5} (bq6)
MATCH (aq7)-[:Knows]->{2,5} (bq7)
-- Subpath Patterns with everything
MATCH ( @{a=1} WALK (asub)-[esub]->(bsub) WHERE asub.id > 0 ){2,3}
RETURN 1
