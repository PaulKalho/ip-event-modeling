-- name: CreateTask :exec
INSERT INTO tasks
	(id, name, status, created_by, created_at)
VALUES ($1, $2, $3, $4, $5);

-- name: UpdateTask :exec
UPDATE tasks
SET name = coalesce(sqlc.narg('name'), name),
	description = coalesce(sqlc.narg('description'), description),
	status = coalesce(sqlc.narg('status'), status),
	public = coalesce(sqlc.narg('public'), public),
	created_by = coalesce(sqlc.narg('created_by'), created_by),
	due_at = coalesce(sqlc.narg('due_at'), due_at)
WHERE id = $1;

-- name: CreateSubtask :exec
INSERT INTO subtasks
	(id, task_id, name, done)
VALUES (@id, @task_id, @name, @done);

-- name: UpdateSubtask :exec
UPDATE subtasks
SET name = coalesce(sqlc.narg('name'), name),
    done = coalesce(sqlc.narg('done'), done)
WHERE id = @id;

-- name: DeleteSubtask :exec
DELETE FROM subtasks WHERE id = @id;

-- name: RemoveTaskDueAt :exec
UPDATE tasks
SET due_at = NULL
WHERE id = @id;

-- name: DeleteTask :exec
DELETE FROM tasks WHERE id = $1;
