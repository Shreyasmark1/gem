package db

import (
	"database/sql"
	"errors"
)

// WrapNotFound turns sql.ErrNoRows into a nil pointer
func DbExe[T any](val T, err error) (*T, error) {
	if err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		// if pgErr, ok := err.(*pq.Error); ok {
		//     switch pgErr.Code {
		//     case "23505": // unique_violation
		//         return fmt.Errorf("author name already exists (code %s)", pgErr.Code)
		//     case "23503": // foreign_key_violation
		//         return fmt.Errorf("associated record not found")
		//     }
		// }
		return nil, err
	}
	return &val, nil
}
