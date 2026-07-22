GRAPH FinGraph
CALL PageRank(node_labels => ['Account'], edge_labels => ['Transfers']) YIELD node, score
RETURN node.id, score AS page_rank
