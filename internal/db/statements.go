package db

var statements = map[string]string{
  // inserts a new row into the mustangs table
  "create-mustang": `
  INSERT INTO mustangs (mustang_id, name)
    values(UUID_TO_BIN(?), ?)
  `,
  // soft deletes a mustang by id
  "delete-mustang": `
  UPDATE
    mustangs
  SET
    deleted_at = NOW()
  WHERE
    mustang_id = UUID_TO_BIN(?)
    AND deleted_at IS NULL
  `,
  // gets a single mustang row by id
  "get-mustang": `
  SELECT
    mustang_id, name
  FROM
    mustangs
  WHERE
    mustang_id = UUID_TO_BIN(?)
    AND deleted_at IS NULL
  `,
  // update a single mustang row by ID
  "update-mustang": `
  UPDATE
    mustangs
  SET
    name = ?
  WHERE
    mustang_id = UUID_TO_BIN(?)
    AND deleted_at IS NULL
  `,
}
