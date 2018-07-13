
CREATE TABLE networks (
  id serial primary key,
  created_at timestamptz DEFAULT current_timestamp,
  updated_at timestamptz DEFAULT NULL,
  addr varchar(255) DEFAULT NULL,
  cidr integer DEFAULT NULL,
  fixed varchar(255)[] DEFAULT NULL,
  comment text DEFAULT NULL,
);

CREATE FUNCTION set_updated_at()
  RETURNS trigger as $$
  BEGIN
    NEW.updated_at := current_timestamp;
    RETURN NEW;
  END $$ language plpgsql;

CREATE TRIGGER networks_updated_at
  BEFORE UPDATE OR INSERT ON networks
  FOR EACH ROW
  EXECUTE PROCEDURE set_updated_at();
