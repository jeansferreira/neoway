package main


//Biblioteca em uso na aplicação

//go get github.com/jmoiron/sqlx
//go get github.com/lib/pq
//go get github.com/acroca
//go get github.com/go-delve
//go get github.com/jeansferreira
//go get github.com/jmiron
//go get github.com/jmoiron
//go get github.com/karrick
//go get github.com/lib
//go get github.com/nsf
//go get github.com/pkg
//go get github.com/ramya-rao-a
//go get github.com/rogpeppe
//go get github.com/sqs
//go get github.com/uudashr

import (
	"io/ioutil"
	"os"
	"bufio"
	"fmt"
	"regexp"
	"log"

	"github.com/jeansferreira/neoway/repo"
	"github.com/jeansferreira/neoway/tratamento"
	"github.com/jeansferreira/neoway/model"
)

func main() {

	dir := "./Dados/"
    files, err := ioutil.ReadDir(dir)
    if err != nil {
        log.Fatal("[main] Ocorreu na pasta de diretório dos arquivos. Erro: ",err.Error())
    }

    for _, f := range files {
		fmt.Println("Carga do Arquivo ", f.Name())
		processaArquivo(f.Name())
	}
}

func processaArquivo(caminhaArquivo string ) {

	file, err := os.Open("./Dados/" + caminhaArquivo)
    if err != nil {
		log.Fatal("[main] Ocorreu erro ao abrir o arquivo. Erro: ",err.Error())
		return
	}
	defer file.Close()
	
    scanner := bufio.NewScanner(file)
    r := regexp.MustCompile("[^\\s]+")

    list_movimento:=[]model.Compra{}

    for scanner.Scan() {
		results := r.FindAllString(scanner.Text(), -1)
		
		mov := model.Compra{}
		mov.Cpf_cnpj_comprador = results[0]
		mov.Flg_private = results[1]
		mov.Flg_incompleto = results[2]
		mov.Dt_ultima_compra = results[3]
		mov.Vl_ticket_medio = results[4]
		mov.Vl_ticket_ult_compra = results[5]
		mov.Cnpj_loja_freq = results[6]
		mov.Cnpj_loja_ultima = results[7]

		if tratamento.ValidateCpfCnpj(results[0]) && tratamento.ValidateCpfCnpj(results[7]) && tratamento.ValidateCpfCnpj(results[7]) {
				list_movimento=append(list_movimento,mov)    
			}
    }

    fmt.Println("CPF/CNPJ que estão válidos para inserir na Base de Dados: ", len(list_movimento))
    for i := 0; i < len(list_movimento); i++ {
        //fmt.Printf("Movimento: %+v\r\n ", list_movimento[i])
		}

		//Nesta etapa o arquivo já está tratado e em objetos dentro do Array, e vai ser enviado para a classe persistir
	repo.InsereDados(list_movimento)
}