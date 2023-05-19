CREATE TABLE IF NOT EXISTS subtasks (
	id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
	task_id uuid NOT NULL,
	name text NOT NULL,
	description text  NOT NULL DEFAULT '',
	done boolean NOT NULL DEFAULT FALSE,
	FOREIGN KEY (task_id)
	    REFERENCES tasks(id)
		ON DELETE CASCADE
);
