# Dev Backend (Golang) - O Guardião dos Dados

Este projeto é uma API RESTful desenvolvida em Golang, utilizando o framework Gin Gonic e o ORM GORM. A API gerencia missões de aventura para uma guilda de aventureiros.

Configuração e Execução
Certifique-se de ter o Go instalado em sua máquina. Você pode baixá-lo em golang.org.
Clone este repositório para sua máquina local.
Instale as dependências do projeto.
Configure as variáveis de ambiente no arquivo .env conforme necessário. Você pode usar o arquivo .env como modelo, nele contém o endereço banco de dados préviamente configurado, não possuindo informações sensíveis.

Execute o seguinte comando para iniciar o servidor:
```
go run main.go
```
A API estará disponível em http://localhost:8080.

## Documentação
A documentação da API foi gerada automaticamente usando o Swagger. Para acessar a documentação, inicie o servidor e navegue até [http://localhost:8080/swagger/index.html](http://localhost:8080/docs/index.html) em seu navegador.
