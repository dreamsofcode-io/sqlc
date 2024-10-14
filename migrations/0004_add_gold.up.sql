ALTER TABLE player ADD COLUMN gold bigint default 0;
ALTER TABLE player ADD CONSTRAINT chk_gold_non_negative CHECK (gold >= 0);
