# Databases - database/sql
*sql.Register*
Databases drivers should be registered in func init() when using third party dbs. It can be helpful to use a ```var drivers = make(map[string]driver.Driver)``` to map table names to drivers.
*driver.Driver*
The Driver interface has an Open(name string) method that returns the connections of the database that it has connected to.
*driver.Conn*
The connection can Prepare(query string) - Close() and Begin()
