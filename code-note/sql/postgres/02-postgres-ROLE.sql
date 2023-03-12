CREATE USER ningHuang WITH PASSWORD 'ningHuang';
GRANT ALL ON COMPANY TO ningHuang;
REVOKE ALL ON COMPANY FROM ningHuang;
drop user ningHuang;

-- GRANT privilege [, ...]
-- ON object [, ...]
-- TO { PUBLIC | GROUP group | username }