ALTER TABLE foo ALTER COLUMN bar TIMESTAMP NOT NULL ON UPDATE (pending_commit_timestamp()) DEFAULT (pending_commit_timestamp())
