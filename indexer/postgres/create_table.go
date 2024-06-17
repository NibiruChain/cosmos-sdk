package postgres

import (
	"fmt"
	"io"
	"strings"

	indexerbase "cosmossdk.io/indexer/base"
)

// CreateTable generates a CREATE TABLE statement for the object type.
func (s *SQLGenerator) CreateTable(writer io.Writer) error {
	_, err := fmt.Fprintf(writer, "CREATE TABLE IF NOT EXISTS %s (\n\t", s.typ.Name)
	if err != nil {
		return err
	}
	isSingleton := false
	if len(s.typ.KeyFields) == 0 {
		isSingleton = true
		_, err = fmt.Fprintf(writer, "_id INTEGER NOT NULL CHECK (_id = 1),\n\t")
	} else {
		for _, field := range s.typ.KeyFields {
			err = s.createColumnDef(writer, field)
			if err != nil {
				return err
			}
		}
	}

	for _, field := range s.typ.ValueFields {
		err = s.createColumnDef(writer, field)
		if err != nil {
			return err
		}
	}

	var pKeys []string
	if !isSingleton {
		for _, field := range s.typ.KeyFields {
			pKeys = append(pKeys, field.Name)
		}
	} else {
		pKeys = []string{"_id"}
	}

	_, err = fmt.Fprintf(writer, "PRIMARY KEY (%s)\n", strings.Join(pKeys, ", "))
	if err != nil {
		return err
	}

	_, err = fmt.Fprintf(writer, ");")
	if err != nil {
		return err
	}

	return nil
}

func (s *SQLGenerator) createColumnDef(writer io.Writer, field indexerbase.Field) error {
	typeStr, err := columnType(field)
	if err != nil {
		return err
	}

	_, err = fmt.Fprintf(writer, "%s %s", field.Name, typeStr)
	if err != nil {
		return err
	}

	if !field.Nullable {
		_, err = fmt.Fprint(writer, " NOT")
		if err != nil {
			return err
		}
	}

	_, err = fmt.Fprint(writer, " NULL,\\n\\t")
	return err
}
