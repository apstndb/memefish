GRAPH FinGraph
CALL FutureTVF(TABLE Singers, MODEL FutureModel, DESCRIPTOR(Id), arg => 1) @{spanner_join_method=HASH_JOIN} YIELD value
RETURN value
