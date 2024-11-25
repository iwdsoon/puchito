package database

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var db *gorm.DB

func InitDb(SqlString string) {
	
	dbName := extractDatabaseName(SqlString)
	rootConnStr := removeDatabaseName(SqlString)
	
	rootConn, err := sql.Open("mysql", rootConnStr)
	if err != nil {
		log.Fatalf("Failed to connect to MySQL server: %v", err)
	}
	defer rootConn.Close()

	_, err = rootConn.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", dbName))
	if err != nil {
		log.Fatalf("Failed to create database: %v", err)
	}
	fmt.Printf("Database '%s' created.\n", dbName)


	conn, err := gorm.Open(mysql.Open(SqlString), &gorm.Config{
		SkipDefaultTransaction: false,
		PrepareStmt:            false,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	db = conn
	db.Logger = db.Logger.LogMode(logger.Info)
	fmt.Println("Database connection established.")

	
	createTables();
	insertSampleData();
}

func createTables() {
	sqlStatements := []string{
		`
		CREATE TABLE IF NOT EXISTS generos (
			id int NOT NULL AUTO_INCREMENT,
			genero varchar(30) NOT NULL,
			PRIMARY KEY (id),
			UNIQUE KEY genero_UNIQUE (genero)
		)
		`,
		`
		CREATE TABLE IF NOT EXISTS libros (
			id int NOT NULL AUTO_INCREMENT,
			nombre varchar(45) NOT NULL,
			autor varchar(45) NOT NULL,
			fecha_lanzamiento datetime NOT NULL,
			id_genero int NOT NULL,
			fecha_creado datetime DEFAULT CURRENT_TIMESTAMP,
			fecha_actualizado datetime DEFAULT CURRENT_TIMESTAMP,
			estado tinyint(1) DEFAULT 1,
			PRIMARY KEY (id),
			KEY fk_genero_idx (id_genero),
			CONSTRAINT fk_genero FOREIGN KEY (id_genero) REFERENCES generos (id)
		)
		`,
		`
		CREATE TABLE IF NOT EXISTS roles (
			id int NOT NULL AUTO_INCREMENT,
			rol varchar(10) NOT NULL,
			PRIMARY KEY (id),
			UNIQUE KEY rol_UNIQUE (rol)
		)
		`,
		`
		CREATE TABLE IF NOT EXISTS usuarios (
			id int NOT NULL AUTO_INCREMENT,
			nombre varchar(45) NOT NULL,
			apellido varchar(45) NOT NULL,
			dni varchar(8) NOT NULL,
			telefono varchar(45) NOT NULL,
			email varchar(45) NOT NULL,
			fecha_creado datetime DEFAULT CURRENT_TIMESTAMP,
			fecha_actualizado datetime DEFAULT CURRENT_TIMESTAMP,
			estado tinyint(1) DEFAULT 1,
			id_rol int NOT NULL,
			PRIMARY KEY (id),
			UNIQUE KEY dni_UNIQUE (dni),
			UNIQUE KEY email_UNIQUE (email),
			KEY fk_rol_idx (id_rol),
			CONSTRAINT fk_rol FOREIGN KEY (id_rol) REFERENCES roles (id)
		)
		`,
		`
		CREATE TABLE IF NOT EXISTS prestamos (
			id int NOT NULL AUTO_INCREMENT,
			id_usuario int NOT NULL,
			id_libro int NOT NULL,
			fecha_prestamo datetime DEFAULT CURRENT_TIMESTAMP,
			fecha_devolucion_estimada datetime DEFAULT NULL,
			fecha_devolucion datetime DEFAULT NULL,
			PRIMARY KEY (id),
			KEY fk_usuario_idx (id_usuario),
			KEY fk_libro_idx (id_libro),
			CONSTRAINT fk_libro FOREIGN KEY (id_libro) REFERENCES libros (id),
			CONSTRAINT fk_usuario FOREIGN KEY (id_usuario) REFERENCES usuarios (id)
		)
		`,
	}

	for _, sql := range sqlStatements {
		if err := db.Exec(sql).Error; err != nil {
			log.Fatalf("Failed to execute statement: %v", err)
		}
	}

	fmt.Println("Tables created.")
}

func insertSampleData() {
    sqlStatements := []string{

        `INSERT INTO generos (genero) VALUES
            ('Accion'),
            ('Romance'),
            ('Misterio'),
            ('Ficción'),
            ('Educación')
        ON DUPLICATE KEY UPDATE genero = genero;`,

        `INSERT INTO roles (rol) VALUES
            ('Admin'),
            ('Cliente')
        ON DUPLICATE KEY UPDATE rol = rol;`,

        `INSERT INTO libros (nombre, autor, fecha_lanzamiento, id_genero, estado) VALUES
            ('The Great Gatsby', 'F. Scott Fitzgerald', '1925-04-10', 1, 1),
            ('A Brief History of Time', 'Stephen Hawking', '1988-03-01', 2, 1),
            ('The History of the World', 'J. M. Roberts', '1976-06-01', 3, 1),
            ('Steve Jobs', 'Walter Isaacson', '2011-10-24', 4, 1),
            ('The Hound of the Baskervilles', 'Arthur Conan Doyle', '1902-08-01', 5, 1)`,

        `INSERT INTO usuarios (nombre, apellido, dni, telefono, email, id_rol, estado) VALUES
            ('Admin', 'Admin', '12345678', '555-1234', 'admin@admin.com', 1, 1),
            ('Juanito', 'Perez', '23456789', '555-5678', 'Juanito@email.com', 2, 1),
            ('Alicia', 'Pepita', '34567890', '555-9101', 'Alicia@email.com', 2, 1),
            ('Matias', 'Giordanno', '45678901', '555-1122', 'Matias@email.com', 2, 1),
            ('Raul', 'Montez', '56789012', '555-3344', 'Raul@email.com', 2, 1)
        ON DUPLICATE KEY UPDATE email = email;`,

        `INSERT INTO prestamos (id_usuario, id_libro, fecha_prestamo, fecha_devolucion_estimada, fecha_devolucion) VALUES
            (1, 1, NOW(), NOW() + INTERVAL 7 DAY, NULL),
            (2, 2, NOW(), NOW() + INTERVAL 7 DAY, NULL),
            (3, 3, NOW(), NOW() + INTERVAL 7 DAY, NULL),
            (4, 4, NOW(), NOW() + INTERVAL 7 DAY, NULL),
            (5, 5, NOW(), NOW() + INTERVAL 7 DAY, NULL)`,
    }

    for _, sql := range sqlStatements {
        if err := db.Exec(sql).Error; err != nil {
            log.Fatalf("Failed to insert data: %v", err)
        }
    }

    fmt.Println("Sample data inserted successfully.")
}

func extractDatabaseName(dsn string) string {

	start := strings.LastIndex(dsn, "/") + 1
	end := strings.Index(dsn[start:], "?")
	if end == -1 {

		end = len(dsn[start:])
	}
	return dsn[start : start+end]
}

func removeDatabaseName(dsn string) string {

	start := strings.LastIndex(dsn, "/") + 1
	end := strings.Index(dsn[start:], "?")
	if end == -1 {
		return dsn[:start]
	}

	return dsn[:start] + dsn[start+end:]
}

func GetDb() *gorm.DB {
	return db
}