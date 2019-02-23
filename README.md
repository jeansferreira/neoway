# Projeto de Load Dados - Neoway (Teste)

Objetivo:
Desenvolva um serviço de manipulação de dados e persistência em base de dados relacional.

Requisitos:
- Criar um serviço em GO que receba um arquivo csv/txt de entrada (Arquivo Anexo)
- Este serviço deve persistir no banco de dados relacional (postgres) todos os dados contidos no arquivo
  Obs: O arquivo não possui um separador muito convencional
 
- Deve-se fazer o split dos dados em colunas no banco de dados
 Obs: pode ser feito diretamente no serviço em GO ou em sql
 
- Realizar higienização dos dados após persistência (sem acento, maiusculo, etc)
- Validar os CPFs/CNPJs contidos (validos e não validos numericamente)
- Todo o código deve estar disponível em repositório publico do GIT
 
Desejável:
- Utilização das linguagen GOLANG para o desenvolvimento do serviço
- Utilização do DB Postgres
- Docker Conpose , com orientações para executar (arquivo readme) 

Você será avaliado por:
- Utilização de melhores práticas de desenvolvimento (nomenclatura, funções, classes, etc);
- Utilização dos recursos mais recentes das linguagens;
- Boa organização lógica e documental (readme, comentários, etc);
- Cobertura de todos os requisitos obrigatórios.

## Arquivos de entrada
- Os arquivos de entrada em CSV ou TXT devem ficar na pasta "./Dados"
- A aplicação obtem os arquivos que estão na pasta definida e lê um de cada vêz e insere na base de dados.

## Tratamento realizado
- Todos os tratamentos solicitados estão no pacote "src/tratamento".

## Base de Dados
- Foi utilizado um docker do Postgres/PgAdmin4 para o desenvolvimento do Projeto.

- Caso a conexão do banco de dados não seja a mesma, favor alterar a classe "/src/repo/conexao.go".

```
const (
    host     = "localhost"
    port     = 5432
    user     = "postgres"
    password = "Postgres2019!"
    dbname   = "neoway"
)

```

### Comando para instalação do Docker e do Postgres/Pgdmin4 

```
sudo apt-get install docker.io
docker pull postgres
sudo docker pull postgres
docker pull dpage/pgadmin4
sudo docker pull dpage/pgadmin4
sudo docker network create --driver bridge postgres-network
sudo docker network ls
docker run --name teste-postgres --network=postgres-network -e "POSTGRES_PASSWORD=Postgres2019!" -p 5432:5432 -v /home/jean/Desenvolvimento/PostgreSQL:/var/lib/postgresql/data -d postgres
sudo docker run --name teste-postgres --network=postgres-network -e "POSTGRES_PASSWORD=Postgres2019!" -p 5432:5432 -v /home/jean/Desenvolvimento/PostgreSQL:/var/lib/postgresql/data -d postgres
docker run --name teste-pgadmin --network=postgres-network -p 15432:80 -e "PGADMIN_DEFAULT_EMAIL=jeansferreira@gmail.com" -e "PGADMIN_DEFAULT_PASSWORD=PgAdmin2019!" -d dpage/pgadmin4

```

### Script de Create da Table (Migracao).
- Obs.: O script é criado automáticamente pela aplicação, sendo que não necessário executar em base de dados. 

```

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

```

### Comandos para instalação dos bibliotecas da aplicação

```
Executar o arquivo "install.sh"

```

### Comandos para excecutar a aplicação

```
cd ../src/github.com/jeansferreira/
go run main.go

```
