package constvar

// tables
var (
	TablesSQLString = []string{
		`CREATE TABLE IF NOT EXISTS auths(
			id          INT UNSIGNED NOT NULL AUTO_INCREMENT,
			username    VARCHAR(50) NOT NULL,
			password    VARCHAR(512) NOT NULL,
			PRIMARY KEY (id)
		  )ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;`,

		`CREATE TABLE IF NOT EXISTS tags(
			id             INT UNSIGNED NOT NULL AUTO_INCREMENT,
			name           VARCHAR(100) NOT NULL,
			created_on     TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			created_by     VARCHAR(100) NOT NULL,
			updated_on      TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
			updated_by      VARCHAR(100) DEFAULT '',
			state          BOOLEAN DEFAULT TRUE,
			deleted_on     INT(11) UNSIGNED DEFAULT 0,
			PRIMARY KEY (id)
		  )ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;`,

		`CREATE TABLE IF NOT EXISTS articles(
			id             INT UNSIGNED NOT NULL AUTO_INCREMENT,
			tag_id         INT DEFAULT 0,
			title          VARCHAR(100) NOT NULL,
			description    VARCHAR(255) DEFAULT '',
			content        TEXT,
			created_on     TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			created_by     VARCHAR(100) NOT NULL,
			updated_on      TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
			updated_by      VARCHAR(100) DEFAULT '',
			state          BOOLEAN DEFAULT TRUE,
			deleted_on     INT(11) UNSIGNED DEFAULT 0,
			PRIMARY KEY (id)
		  )ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;`,

		`CREATE TABLE IF NOT EXISTS files (
			article_id  INTEGER UNSIGNED NOT NULL,
			path         VARCHAR(512) DEFAULT ' ',
			md5        VARCHAR(512) DEFAULT ' ',
			created_on  DATETIME DEFAULT CURRENT_TIMESTAMP,
			PRIMARY KEY (md5)
		) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;`,
	}
)
