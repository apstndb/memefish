ALTER TABLE foo ADD COLUMN bar TIMESTAMP NOT NULL ON UPDATE (pending_commit_timestamp()) DEFAULT (pending_commit_timestamp())
