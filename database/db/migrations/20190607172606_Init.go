
package main

import (
	"database/sql"
	"log"
)

// Up is executed when this migration is applied
func Up_20190607172606(txn *sql.Tx) {
	_,err:=txn.Exec(`
		CREATE TABLE way_api."user"
		(
			user_id bigserial NOT NULL,
    		first_name character varying(100),
    		last_name character varying(100) NOT NULL,
    		phone_number character varying(12) NOT NULL UNIQUE,
    		verified boolean NOT NULL DEFAULT false ,
    		temporary_pin character varying(6),
    		temporary_pin_expiry timestamp with time zone,
    		created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    		updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    		PRIMARY KEY (user_id)
		)
		WITH (
    		OIDS = FALSE
		);

		ALTER TABLE way_api."user"
    	OWNER to way;
		`)

	if err != nil {
		log.Println(err)
		return
	}

	_,err=txn.Exec(`
		CREATE TABLE way_api."login"
		(
			login_id bigserial NOT NULL,
			user_id INTEGER NOT NULL REFERENCES way_api.user(user_id),
    		username character varying(100) NOT NULL,
    		password character varying(200) NOT NULL,
    		created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    		updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    		PRIMARY KEY (login_id)
		)
		WITH (
    		OIDS = FALSE
		);

		ALTER TABLE way_api."login"
    	OWNER to way;
		`)

	if err != nil {
		log.Println(err)
		return
	}

	_,err=txn.Exec(`
		CREATE TABLE way_api."relationship"
		(
			relationship_id bigserial NOT NULL,
			user_1 INTEGER NOT NULL REFERENCES way_api.user(user_id),
    		user_2 INTEGER NOT NULL REFERENCES way_api.user(user_id),
    		status character varying(100) NOT NULL,
    		last_actor INTEGER NOT NULL REFERENCES way_api.user(user_id),
    		user_1_trusted boolean NOT NULL default false,
    		user_2_trusted boolean NOT NULL default false,
    		created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    		updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    		PRIMARY KEY (relationship_id)
		)
		WITH (
    		OIDS = FALSE
		);

		ALTER TABLE way_api."relationship"
    	OWNER to way;
		`)

	if err != nil {
		log.Println(err)
		return
	}
}

// Down is executed when this migration is rolled back
func Down_20190607172606(txn *sql.Tx) {
	_,err:=txn.Exec(`
		DROP TABLE way_api."relationship",way_api."login",way_api."user"`)

	if err != nil {
		log.Println(err)
		return
	}
}
