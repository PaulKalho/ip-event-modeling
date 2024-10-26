CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS tasks (
	id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
	name text NOT NULL,
	description text DEFAULT ''::text NOT NULL,
	status integer NOT NULL,
	public boolean DEFAULT false NOT NULL,
	created_by uuid NOT NULL,
	due_at timestamp DEFAULT CURRENT_TIMESTAMP,
	created_at timestamp DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS subtasks (
	id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
	task_id uuid NOT NULL,
	name text NOT NULL,
	done boolean NOT NULL DEFAULT FALSE,
	FOREIGN KEY (task_id)
		REFERENCES tasks(id)
		ON DELETE CASCADE
);
