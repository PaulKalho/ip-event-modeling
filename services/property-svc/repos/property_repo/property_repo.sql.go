// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: property_repo.sql

package property_repo

import (
	"context"

	"github.com/google/uuid"
)

const createProperty = `-- name: CreateProperty :exec
INSERT INTO properties
	(id, subject_type, field_type, name)
VALUES ($1, $2, $3, $4)
`

type CreatePropertyParams struct {
	ID          uuid.UUID
	SubjectType int32
	FieldType   int32
	Name        string
}

func (q *Queries) CreateProperty(ctx context.Context, arg CreatePropertyParams) error {
	_, err := q.db.Exec(ctx, createProperty,
		arg.ID,
		arg.SubjectType,
		arg.FieldType,
		arg.Name,
	)
	return err
}

const createSelectData = `-- name: CreateSelectData :one
INSERT INTO select_datas
	(allow_freetext)
VALUES ($1)
RETURNING id
`

func (q *Queries) CreateSelectData(ctx context.Context, allowFreetext bool) (uuid.UUID, error) {
	row := q.db.QueryRow(ctx, createSelectData, allowFreetext)
	var id uuid.UUID
	err := row.Scan(&id)
	return id, err
}

const createSelectOption = `-- name: CreateSelectOption :exec
INSERT INTO select_options
	(id, name, description, is_custom, select_data_id)
VALUES ($1, $2, coalesce($3, ''), $4, $5)
`

type CreateSelectOptionParams struct {
	ID           uuid.UUID
	Name         string
	Description  interface{}
	IsCustom     bool
	SelectDataID uuid.UUID
}

func (q *Queries) CreateSelectOption(ctx context.Context, arg CreateSelectOptionParams) error {
	_, err := q.db.Exec(ctx, createSelectOption,
		arg.ID,
		arg.Name,
		arg.Description,
		arg.IsCustom,
		arg.SelectDataID,
	)
	return err
}

const deleteSelectDataByPropertyID = `-- name: DeleteSelectDataByPropertyID :exec
DELETE FROM select_datas
    WHERE id IN (
        SELECT properties.select_data_id
        FROM properties
        WHERE properties.id = $1
	)
`

func (q *Queries) DeleteSelectDataByPropertyID(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.Exec(ctx, deleteSelectDataByPropertyID, id)
	return err
}

const getPropertiesWithSelectDataAndOptionsBySubjectTypeOrID = `-- name: GetPropertiesWithSelectDataAndOptionsBySubjectTypeOrID :many
SELECT
	properties.id, properties.subject_type, properties.field_type, properties.name, properties.description, properties.is_archived, properties.set_id, properties.select_data_id,
	select_options.id as select_option_id,
	select_options.name as select_option_name,
	select_options.description as select_option_description,
	select_options.is_custom as select_option_is_custom,
	select_datas.id as select_datas_id,
	select_datas.allow_freetext as select_datas_allow_freetext
	FROM properties
	LEFT JOIN select_datas ON properties.select_data_id = select_datas.id
	LEFT JOIN select_options ON select_options.select_data_id = select_datas.id
 	WHERE subject_type = $1 OR properties.id = $2
`

type GetPropertiesWithSelectDataAndOptionsBySubjectTypeOrIDParams struct {
	SubjectType *int32
	ID          uuid.NullUUID
}

type GetPropertiesWithSelectDataAndOptionsBySubjectTypeOrIDRow struct {
	Property                 Property
	SelectOptionID           uuid.NullUUID
	SelectOptionName         *string
	SelectOptionDescription  *string
	SelectOptionIsCustom     *bool
	SelectDatasID            uuid.NullUUID
	SelectDatasAllowFreetext *bool
}

func (q *Queries) GetPropertiesWithSelectDataAndOptionsBySubjectTypeOrID(ctx context.Context, arg GetPropertiesWithSelectDataAndOptionsBySubjectTypeOrIDParams) ([]GetPropertiesWithSelectDataAndOptionsBySubjectTypeOrIDRow, error) {
	rows, err := q.db.Query(ctx, getPropertiesWithSelectDataAndOptionsBySubjectTypeOrID, arg.SubjectType, arg.ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetPropertiesWithSelectDataAndOptionsBySubjectTypeOrIDRow{}
	for rows.Next() {
		var i GetPropertiesWithSelectDataAndOptionsBySubjectTypeOrIDRow
		if err := rows.Scan(
			&i.Property.ID,
			&i.Property.SubjectType,
			&i.Property.FieldType,
			&i.Property.Name,
			&i.Property.Description,
			&i.Property.IsArchived,
			&i.Property.SetID,
			&i.Property.SelectDataID,
			&i.SelectOptionID,
			&i.SelectOptionName,
			&i.SelectOptionDescription,
			&i.SelectOptionIsCustom,
			&i.SelectDatasID,
			&i.SelectDatasAllowFreetext,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getPropertyById = `-- name: GetPropertyById :one
SELECT id, subject_type, field_type, name, description, is_archived, set_id, select_data_id FROM properties WHERE id = $1
`

func (q *Queries) GetPropertyById(ctx context.Context, id uuid.UUID) (Property, error) {
	row := q.db.QueryRow(ctx, getPropertyById, id)
	var i Property
	err := row.Scan(
		&i.ID,
		&i.SubjectType,
		&i.FieldType,
		&i.Name,
		&i.Description,
		&i.IsArchived,
		&i.SetID,
		&i.SelectDataID,
	)
	return i, err
}

const updateProperty = `-- name: UpdateProperty :exec
UPDATE properties
SET subject_type = coalesce($2, subject_type),
    field_type = coalesce($3, field_type),
    name = coalesce($4, name),
    description = coalesce($5, description),
    is_archived = coalesce($6, is_archived)
WHERE id = $1
`

type UpdatePropertyParams struct {
	ID          uuid.UUID
	SubjectType *int32
	FieldType   *int32
	Name        *string
	Description *string
	IsArchived  *bool
}

func (q *Queries) UpdateProperty(ctx context.Context, arg UpdatePropertyParams) error {
	_, err := q.db.Exec(ctx, updateProperty,
		arg.ID,
		arg.SubjectType,
		arg.FieldType,
		arg.Name,
		arg.Description,
		arg.IsArchived,
	)
	return err
}

const updatePropertySelectDataID = `-- name: UpdatePropertySelectDataID :exec
UPDATE properties
SET select_data_id = $2
WHERE id = $1
`

type UpdatePropertySelectDataIDParams struct {
	ID           uuid.UUID
	SelectDataID uuid.NullUUID
}

func (q *Queries) UpdatePropertySelectDataID(ctx context.Context, arg UpdatePropertySelectDataIDParams) error {
	_, err := q.db.Exec(ctx, updatePropertySelectDataID, arg.ID, arg.SelectDataID)
	return err
}

const updatePropertySetID = `-- name: UpdatePropertySetID :exec
UPDATE properties
SET set_id = $1
WHERE id = $2
`

type UpdatePropertySetIDParams struct {
	SetID uuid.NullUUID
	ID    uuid.UUID
}

func (q *Queries) UpdatePropertySetID(ctx context.Context, arg UpdatePropertySetIDParams) error {
	_, err := q.db.Exec(ctx, updatePropertySetID, arg.SetID, arg.ID)
	return err
}

const updateSelectData = `-- name: UpdateSelectData :exec
UPDATE select_datas
SET allow_freetext = $1
WHERE id = $2
`

type UpdateSelectDataParams struct {
	AllowFreetext bool
	ID            uuid.UUID
}

func (q *Queries) UpdateSelectData(ctx context.Context, arg UpdateSelectDataParams) error {
	_, err := q.db.Exec(ctx, updateSelectData, arg.AllowFreetext, arg.ID)
	return err
}
