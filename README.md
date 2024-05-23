### Web Application Commands
```sh
go run ./cmd/web
```

```sh
go get github.com/go-sql-driver/mysql
```

```sh
sudo apt install mysql-server
```

```sh
sudo /etc/init.d/mysql start
```

```sh
sudo mysql_secure_installation
```

```sh
sudo mysql
```

```sql
CREATE DATABASE snippetbox CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
```

```sql
USE snippetbox;
```

```sql
CREATE TABLE snippets (
 id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
 title VARCHAR(100) NOT NULL,
 content TEXT NOT NULL,
 created DATETIME NOT NULL,
 expires DATETIME NOT NULL
);
```

```sql
CREATE INDEX idx_snippets_created ON snippets(created);
```

```sql
INSERT INTO snippets (title, content, created, expires) VALUES (
 'An old silent pond',
 'An old silent pond...\nA frog jumps into the pond,\nsplash! Silence again.\n\n– Matsuo Bashō',
 UTC_TIMESTAMP(),
 DATE_ADD(UTC_TIMESTAMP(), INTERVAL 365 DAY)
);
```

```sql
INSERT INTO snippets (title, content, created, expires) VALUES (
 'Over the wintry forest',
 'Over the wintry\nforest, winds howl in rage\nwith no leaves to blow.\n\n– Natsume Soseki',
 UTC_TIMESTAMP(),
 DATE_ADD(UTC_TIMESTAMP(), INTERVAL 365 DAY)
);
```

```sql
INSERT INTO snippets (title, content, created, expires) VALUES (
 'First autumn morning',
 'First autumn morning\nthe mirror I stare into\nshows my father''s face.\n\n– Murakami Kijo',
 UTC_TIMESTAMP(),
 DATE_ADD(UTC_TIMESTAMP(), INTERVAL 7 DAY)
);
```

```sql
CREATE USER 'web'@'localhost';
GRANT SELECT, INSERT, UPDATE, DELETE ON snippetbox.* TO 'web'@'localhost';
ALTER USER 'web'@'localhost' IDENTIFIED BY 'pass';
```

```sh
exit
```

```sh
sudo chmod -R 755 /var/run/mysqld
```

```sh
sudo systemctl restart mysql
```

```sh
mysql -D snippetbox -u web -p
```

```sql
SELECT id, title, expires FROM snippets;
```


