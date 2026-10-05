`WALK` = (`WALK`)-[`TRAIL`]->(`SIMPLE`),
`ACYCLIC` = (`ACYCLIC`)-[`PATH`]->(`PATHS`),
`PATH` = (`walk` IS Person)-[`Trail`:Knows]->(`Simple`),
`PATHS` = (`path`),
`TRAIL` = (`trail`),
`SIMPLE` = (`simple`),
ordinary = TRAIL PATHS (a:WALK)-[e:PATH]->(b)
