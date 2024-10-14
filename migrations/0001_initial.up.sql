CREATE TABLE player(
  id SERIAL PRIMARY KEY,
  name VARCHAR NOT NULL,
  level INT NOT NULL,
  class VARCHAR NOT NULL,
  created_at TIMESTAMPTZ NOT NULL,
  updated_at TIMESTAMPTZ NOT NULL
);

CREATE UNIQUE INDEX on player(name);

INSERT INTO player (id, name, level, class, created_at, updated_at) VALUES
(1, 'Sephiro', 32, 'Warrior', '2024-01-15 14:23:45Z', '2024-02-10 08:30:15Z'),
(2, 'SpellQueen', 47, 'Mage', '2023-12-05 10:15:12Z', '2024-01-22 16:45:22Z'),
(3, 'StormcallerX', 29, 'Druid', '2023-11-20 09:20:34Z', '2024-01-10 12:10:58Z'),
(4, 'ShadowMaster', 50, 'Rogue', '2023-10-08 19:50:14Z', '2023-12-18 11:35:45Z'),
(5, 'FireFurry', 55, 'Sorcerer', '2023-09-03 17:45:23Z', '2023-11-14 14:20:30Z');
