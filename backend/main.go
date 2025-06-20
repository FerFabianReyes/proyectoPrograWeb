package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type vmIdentifica struct {
	Nominal float32 `json:"nominal"`
	Minimo  float32 `json:"minimo"`
	Maximo  float32 `json:"maximo"`
	UM      string  `json:"um"`
}

type vmOpera struct {
	Valor     float32 `json:"valor"`
	IdEstatus int32   `json:"idEstatus"`
}

type identifica struct {
	ID          int32        `json:"identificador"`
	Nombre      string       `json:"nombre"`
	Ubicacion   string       `json:"ubicacion"`
	Coordenadas string       `json:"coordenadas"`
	Potencia    vmIdentifica `json:"potencia"`
	Voltaje     vmIdentifica `json:"voltaje"`
	Corriente   vmIdentifica `json:"corriente"`
	Caudal      vmIdentifica `json:"caudal"`
	Fecha       string       `json:"fechaRegistro"`
}

type opera struct {
	Potencia  vmOpera `json:"potencia"`
	Voltaje   vmOpera `json:"voltaje"`
	Corriente vmOpera `json:"corriente"`
	Caudal    vmOpera `json:"caudal"`
	IdEstatus int32   `json:"idEstatus"`
	Estatus   string  `json:"estatus"`
	Fecha     string  `json:"fechaRegistro"`
}

type dispositivo struct {
	Identifica identifica `json:"identifica"`
	Opera      opera      `json:"opera"`
	Estado     int32      `json:"estado"`
}

type mensaje struct {
	Codigo      int32  `json:"codigo"`
	Descripcion string `json:"descripcion"`
}

// var dispositivos []dispositivo
var db *sql.DB
var msg mensaje

func main() {
	var err error

	db, err = sql.Open("postgres", "postgres://adminuno:pwod.2025@localhost/simulacion?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	router := gin.Default()
	router.GET("/dispositivos", getDispositivosDb)
	//router.GET("/dispositivos/:id", getDispositivoByID)
	router.POST("/dispositivos", createDispositivosDb)
	router.POST("/dispositivos/delete", deleteDispositivoByID)

	router.Run("localhost:8080")
}

func getDispositivosDb(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	rows, err := db.Query(`
		select id, nombre, ubicacion, coordenadas,
			(potencia->>'nominal')::decimal(10,2) as pnominal,(potencia->>'minimo')::decimal(10,2) as pminimo, (potencia->>'maximo')::decimal(10,2) as pmaximo, (potencia->>'um')::text as pum,
			(voltaje->>'nominal')::decimal(10,2) as vnominal,(voltaje->>'minimo')::decimal(10,2) as vminimo, (voltaje->>'maximo')::decimal(10,2) as vmaximo, (voltaje->>'um')::text as vum,
			(corriente->>'nominal')::decimal(10,2) as anominal,(corriente->>'minimo')::decimal(10,2) as aminimo, (corriente->>'maximo')::decimal(10,2) as cmaximo, (corriente->>'um')::text as aum,
			(caudal->>'nominal')::decimal(10,2) as cnominal,(caudal->>'minimo')::decimal(10,2) as cminimo, (caudal->>'maximo')::decimal(10,2) as cmaximo, (caudal->>'um')::text as cum,
			to_char(registro_fecha,'DD/MM/YYYY') as registro_fecha,estado
		from sistemas.dispositivos
	`)

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var dispositivos []dispositivo
	for rows.Next() {
		var d dispositivo
		/*
			d.Identifica.ID = 1
			d.Identifica.Nombre = "Represa Uno"
			d.Identifica.Ubicacion = "En la compuerta de la Represa Uno"
			d.Identifica.Coordenadas = "19.7060° N, 101.1950° W"
			d.Identifica.Potencia.Nominal = 7.200
			d.Identifica.Potencia.Minimo = 6.200
			d.Identifica.Potencia.Maximo = 8.600
			d.Identifica.Potencia.UM = "KW"
			d.Identifica.Voltaje.Nominal = 240
			d.Identifica.Voltaje.Minimo = 230
			d.Identifica.Voltaje.Maximo = 250
			d.Identifica.Voltaje.UM = "Volts"
			d.Identifica.Corriente.Nominal = 30
			d.Identifica.Corriente.Minimo = 25
			d.Identifica.Corriente.Maximo = 35
			d.Identifica.Corriente.UM = "Amperes"
			d.Identifica.Caudal.Nominal = 1
			d.Identifica.Caudal.Minimo = 0.90
			d.Identifica.Caudal.Maximo = 1.20
			d.Identifica.Caudal.UM = "m3/minuto"
			d.Identifica.Fecha = ""
		*/
		d.Opera.Potencia.Valor = 7.200
		d.Opera.Potencia.IdEstatus = 1
		d.Opera.Voltaje.Valor = 240
		d.Opera.Voltaje.IdEstatus = 1
		d.Opera.Corriente.Valor = 30
		d.Opera.Corriente.IdEstatus = 1
		d.Opera.Caudal.Valor = 1
		d.Opera.Caudal.IdEstatus = 1
		d.Opera.IdEstatus = 1
		d.Opera.Estatus = "Operacion Normal"
		d.Opera.Fecha = ""
		d.Estado = 1

		err := rows.Scan(
			&d.Identifica.ID,
			&d.Identifica.Nombre,
			&d.Identifica.Ubicacion,
			&d.Identifica.Coordenadas,
			&d.Identifica.Potencia.Nominal,
			&d.Identifica.Potencia.Minimo,
			&d.Identifica.Potencia.Maximo,
			&d.Identifica.Potencia.UM,
			&d.Identifica.Voltaje.Nominal,
			&d.Identifica.Voltaje.Minimo,
			&d.Identifica.Voltaje.Maximo,
			&d.Identifica.Voltaje.UM,
			&d.Identifica.Corriente.Nominal,
			&d.Identifica.Corriente.Minimo,
			&d.Identifica.Corriente.Maximo,
			&d.Identifica.Corriente.UM,
			&d.Identifica.Caudal.Nominal,
			&d.Identifica.Caudal.Minimo,
			&d.Identifica.Caudal.Maximo,
			&d.Identifica.Caudal.UM,
			&d.Identifica.Fecha,
			&d.Estado,
		)
		if err != nil {
			log.Fatal(err)
		}
		dispositivos = append(dispositivos, d)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	setOpera()
	c.IndentedJSON(http.StatusOK, dispositivos)
}

func createDispositivosDb(c *gin.Context) {
	var err error
	fmt.Println("POST Request")
	//c.Header("Content-Type", "application/json")
	c.Header("Content-type", "application/x-www-form-urlencoded")
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	var (
		nombre, ubicacion, coordenadas string
		pNom, pMin, pMax               float64
		vNom, vMin, vMax               float64
		aNom, aMin, aMax               float64
		cNom, cMin, cMax               float64
		pU, vU, aU, cU                 string
	)

	nombre = c.PostForm("nombre")
	ubicacion = c.PostForm("ubicacion")
	coordenadas = c.PostForm("coordenadas")
	pNom, _ = strconv.ParseFloat(c.PostForm("pNom"), 32)
	pMin, _ = strconv.ParseFloat(c.PostForm("pMin"), 32)
	pMax, _ = strconv.ParseFloat(c.PostForm("pMax"), 32)
	pU = c.PostForm("pU")
	vNom, _ = strconv.ParseFloat(c.PostForm("vNom"), 32)
	vMin, _ = strconv.ParseFloat(c.PostForm("vMin"), 32)
	vMax, _ = strconv.ParseFloat(c.PostForm("vMax"), 32)
	vU = c.PostForm("vU")
	aNom, _ = strconv.ParseFloat(c.PostForm("aNom"), 32)
	aMin, _ = strconv.ParseFloat(c.PostForm("aMin"), 32)
	aMax, _ = strconv.ParseFloat(c.PostForm("aMax"), 32)
	aU = c.PostForm("aU")
	cNom, _ = strconv.ParseFloat(c.PostForm("cNom"), 32)
	cMin, _ = strconv.ParseFloat(c.PostForm("cMin"), 32)
	cMax, _ = strconv.ParseFloat(c.PostForm("cMax"), 32)
	cU = c.PostForm("cU")

	query := "insert into sistemas.dispositivos (nombre,ubicacion,coordenadas,potencia,voltaje,corriente,caudal) values" +
		"($1,$2,$3," +
		fmt.Sprintf("'{\"nominal\": %f, \"minimo\": %f, \"maximo\": %f, \"um\":\"%s\" }',", pNom, pMin, pMax, pU) +
		fmt.Sprintf("'{\"nominal\": %f, \"minimo\": %f, \"maximo\": %f, \"um\":\"%s\" }',", vNom, vMin, vMax, vU) +
		fmt.Sprintf("'{\"nominal\": %f, \"minimo\": %f, \"maximo\": %f, \"um\":\"%s\" }',", aNom, aMin, aMax, aU) +
		fmt.Sprintf("'{\"nominal\": %f, \"minimo\": %f, \"maximo\": %f, \"um\":\"%s\" }');", cNom, cMin, cMax, cU)

	stmt, err := db.Prepare(query)
	msg.Codigo = 201
	msg.Descripcion = "Registro insertado"
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ubicacion)
	defer stmt.Close()
	if _, err := stmt.Exec(nombre, ubicacion, coordenadas); err != nil {
		msg.Codigo = 401
		msg.Descripcion = err.Error()
		log.Fatal(err)
	}

	if err != nil {
		msg.Codigo = 401
		msg.Descripcion = err.Error()
		log.Fatal(err)
	}
	fmt.Println(msg)
	c.IndentedJSON(http.StatusOK, msg)
}

func deleteDispositivoByID(c *gin.Context) {
	var err error
	fmt.Println("DELETE Request")
	//c.Header("Content-Type", "application/json")
	c.Header("Content-type", "application/x-www-form-urlencoded")
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	var (
		id int
	)

	id, _ = strconv.Atoi(c.Param("id"))

	query := "delete from sistemas.dispositivos where id=$1;"

	stmt, err := db.Prepare(query)
	msg.Codigo = 201
	msg.Descripcion = "Registro eliminado"
	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()
	if _, err := stmt.Exec(id); err != nil {
		msg.Codigo = 401
		msg.Descripcion = err.Error()
		log.Fatal(err)
	}

	if err != nil {
		msg.Codigo = 401
		msg.Descripcion = err.Error()
		log.Fatal(err)
	}
	fmt.Println(msg)
	c.IndentedJSON(http.StatusOK, msg)
}

func setOpera() {
	r := rand.New(rand.NewSource(99))
	fmt.Println(r.NormFloat64(), r.Float64(), r.NormFloat64(), r.Float64())
}

//
//func getDispositivos(c *gin.Context) {
//	c.IndentedJSON(http.StatusOK, dispositivos)
//}
