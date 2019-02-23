package repo

import (
	"github.com/jeansferreira/neoway/tratamento"
	"github.com/jeansferreira/neoway/model"
    "fmt"
    
    _ "github.com/lib/pq"
    "github.com/jmoiron/sqlx"
)

var schema = `
CREATE TABLE IF NOT EXISTS public."Mercado"
(
    cpf_cnpj_comprador character varying(14) COLLATE pg_catalog."default" NOT NULL,
    flg_private boolean,
    flg_incompleto boolean,
    dt_ultima_compra date,
    vl_ticket_medio numeric(10,2),
    vl_ticket_ult_compra numeric(10,2),
    cnpj_loja_freq character varying(14) COLLATE pg_catalog."default" NOT NULL,
    cnpj_loja_ultima character varying(14) COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT "Mercado_pkey" PRIMARY KEY (cpf_cnpj_comprador)
)
`
var sqlInsert = `INSERT INTO public."Mercado"(cpf_cnpj_comprador, flg_private, flg_incompleto, 
                                                dt_ultima_compra, vl_ticket_medio, vl_ticket_ult_compra, 
                                                cnpj_loja_freq, cnpj_loja_ultima) 
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
    
const (
    host     = "teste-postgres-compose"
    port     = 16543
    user     = "postgres"
    password = "Postgres2019!"
    dbname   = "postgres"
  )
  
//InsereDados metodo que insere os dados dos arquivo na base de dados
func InsereDados(mov []model.Compra) {

    connStr := "host="+host+" dbname="+dbname+" sslmode=disable user="+user+" password="+password
    
    db, err := sqlx.Connect("postgres", connStr)
    
    if err != nil {
        fmt.Println("[repo] Erro no conexão com banco de dados. Erro:", err.Error())
    }
    
    // Cria a tabela na base de dados
    db.MustExec(schema)
    
    tx := db.MustBegin()

    for i := 0; i < len(mov); i++ {
        fmt.Println(">>> ", mov[i].Cpf_cnpj_comprador)
        tx.MustExec(sqlInsert,  tratamento.RemoveCaracteres(mov[i].Cpf_cnpj_comprador),
                                mov[i].GetFlgPrivate(), 
                                mov[i].GetFlgIncompleto(),
                                mov[i].Dt_ultima_compra, 
                                mov[i].GetVlTicketMedio(), 
                                mov[i].GetVlTicketUltCompra(),
                                tratamento.RemoveCaracteres(mov[i].Cnpj_loja_freq),
                                tratamento.RemoveCaracteres(mov[i].Cnpj_loja_ultima))
    }
    
    tx.Commit()

}