ALTER TABLE folders
ADD CONSTRAINT folders_owner_sub_name_unique
UNIQUE (owner_sub, name);