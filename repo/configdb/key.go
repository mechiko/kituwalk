package configdb

import (
	"errors"
	"fmt"
	"strings"

	"github.com/upper/db/v4"
)

func (c *DbConfig) Key(k string) (out string, err error) {
	if c == nil {
		return "", fmt.Errorf("%s: receiver is nil", modError)
	}
	if c.dbSession == nil {
		return "", fmt.Errorf("%s: dbSession is nil (did you call Check()?)", modError)
	}
	k = strings.TrimSpace(k)
	if k == "" {
		return "", fmt.Errorf("%s: key is empty", modError)
	}
	param := &Parameters{}
	coll := c.dbSession.Collection("parameters")
	if err := coll.Find(db.Cond{"name": k}).One(param); err != nil {
		if errors.Is(err, db.ErrNoMoreRows) {
			return "", fmt.Errorf("%s: key %q not found", modError, k)
		}
		return "", fmt.Errorf("%s: query key %q: %w", modError, k, err)
	}
	return param.Value, nil

}
