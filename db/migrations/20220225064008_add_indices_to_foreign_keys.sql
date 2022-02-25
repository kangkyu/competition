-- migrate:up transaction:false
CREATE INDEX IF NOT EXISTS index_scores_on_contestant_id ON scores (contestant_id);
CREATE INDEX IF NOT EXISTS index_scores_on_user_id ON scores (user_id);

-- migrate:down transaction:false
DROP INDEX IF EXISTS index_scores_on_user_id;
DROP INDEX IF EXISTS index_scores_on_contestant_id;
