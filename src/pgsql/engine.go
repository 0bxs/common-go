package pgsql

import "xorm.io/xorm"

type Engine struct {
	*xorm.Engine
}

func new(engine *xorm.Engine) *Engine {
	return &Engine{Engine: engine}
}

func (client *Engine) Transaction(f func(*xorm.Session) (any, error)) any {
	session := client.NewSession()
	defer func(session *xorm.Session) {
		_ = session.Close()
	}(session)

	if err := session.Begin(); err != nil {
		panic(err)
	}

	result, err := f(session)
	if err != nil {
		_ = session.Rollback()
		panic(err)
	}

	if err = session.Commit(); err != nil {
		_ = session.Rollback()
		panic(err)
	}

	return result
}

func (client *Engine) TransactionCatch(fn func(*xorm.Session) any) any {
	session := client.NewSession()

	defer func(session *xorm.Session) {
		if err := recover(); err != nil {
			_ = session.Rollback()
			panic(err)
		}
	}(session)

	defer func(session *xorm.Session) {
		_ = session.Close()
	}(session)

	if err := session.Begin(); err != nil {
		panic(err)
	}

	res := fn(session)
	if err, ok := res.(error); ok && err != nil {
		_ = session.Rollback()
		panic(err)
	}
	if err := session.Commit(); err != nil {
		_ = session.Rollback()
		panic(err)
	}

	return res
}
