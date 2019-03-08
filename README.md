# OGRE API

Online Game Record Engine API - a version of the ORK API built on Go.

## Usage

Available routes:

### Player Data

`GET /players`: Returns a JSON array of every player in the system.

`GET /player/111`: Returns a JSON object for the player with ID #111.

`GET /player/111/classes`: Returns a JSON array of class information
for the player with ID #111.

### Kingdom Data

`GET /kingdoms`: Returns a JSON array of every kingdom in the system.

`GET /kingdom/111`: Returns a JSON object for the kingdom with ID #111.

`GET /kingdom/111/events`: Returns a JSON array of upcoming events for a 
kingdom with ID #111.

`GET /kingdom/111/officers`: Returns a JSON array of officers for a 
kingdom with ID #111.

## Getting Set Up for Development

This assumes you already have Go 1.11 or later installed.

### MySQL

For development, you'll need a MySQL database running somewhere accessible.
Create a user and a database, and grant all permissions to the user for
the database. For example:

```
CREATE USER 'ogre'@'%' IDENTIFIED BY 'supersecurepassword';
CREATE DATABASE 'ogre';
GRANT ALL ON `ogre`.* TO 'ogre'@'%';
```

Then, you'll need to import data from a SQL backup of the ORK production
database for testing purposes.

For example:

```
mysql -u ogre -p ogre < databackup.sql
```

_Eventually, there'll be a database seeder to create a bunch of dummy data
for this. At that point, the import will not be necessary._

### Running In Development

The development command you'll want to run is:

```
OGRE_DB_USERNAME=ogre \
OGRE_DB_PASSWORD=supersecurepassword \
OGRE_DB_HOSTNAME= \
OGRE_DB_NAME=ogre \
go run *.go
```

This will fire up the API running on port 3736. Note that all configuration is
done with environment variables.

When you're ready to build, just run `go build` in the root directory. It will create a binary
named `ogre`.
