# API-CRUD com Go e PostgreSQL

Este é um projeto de exemplo que implementa uma aplicação web simples utilizando Go para o back-end, PostgreSQL para persistência de dados, e uma interface web simples para interagir com a aplicação. O projeto é containerizado utilizando Docker e Docker Compose, permitindo uma configuração e execução simples e consistente.

## Diagrama da Arquitetur

Aqui está um diagrama básico para ilustrar a arquitetura do projeto:

![Arquitetura](https://github.com/user-attachments/assets/15f3eab7-e17a-46ef-b170-59df28c17c75)

- **Frontend:** Interface HTML/CSS simples servida pelo Go.
- **Backend:** Aplicação Go que gerencia a lógica de negócios e a comunicação com o banco de dados.
- **Banco de Dados:** PostgreSQL para persistência dos dados.

## Pré-requisitos

Certifique-se de ter os seguintes softwares instalados na sua máquina:

- **Docker:** Para construir e rodar os containers.
- **Docker Compose:** Para orquestrar os containers.
- **Go:** Para desenvolvimento local, caso você queira fazer alterações no código.

## Configuração do Ambiente

1. **Clone o repositório:**

   ```bash
   git clone https://github.com/johnpaulnasc/api-crud.git
   cd api-cru
   
2. **Configure as variáveis de ambiente:**

   Crie um arquivo `.env` na raiz do projeto com as seguintes variáveis:
   ```bash
   DATABASE_USER=youruser
   DATABASE_PASSWORD=yourpassword
   DATABASE_HOST=postgres
   DATABASE_PORT=5432
   DATABASE_NAME=yourdbname

## Como Rodar o Projeto

1. **Construa e inicie os containers:**

   Utilize o Docker Compose para construir a imagem da aplicação e iniciar os containers:
   ```bash
   docker-compose up --build

2. **Acesse a aplicação:**

   A aplicação estará disponível em `http://localhost:8080`.
   
## Endpoints

- **GET /items:** Lista todos os itens.
- **POST /items:** Cria um novo item.
- **PUT /items:** Atualiza um item existente (a ser implementado).
- **DELETE /items:** Deleta um item existente (a ser implementado).

## Funcionalidades

- **CRUD Completo:** A aplicação permite criar, ler, atualizar e deletar itens no banco de dados PostgreSQL.
- **Interface Web Simples:** Um frontend minimalista permite a interação com a API através do navegador.
- **Containerização Completa:** Tudo roda em containers, garantindo consistência no ambiente de desenvolvimento.

## Como Contribuir

Sinta-se à vontade para contribuir com este projeto. Aqui estão algumas maneiras:
- **Reporte Bugs:** Use a seção de "Issues" no GitHub para reportar problemas.
- **Submeta PRs:** Envie pull requests para adicionar funcionalidades ou corrigir problemas.
- **Sugira Melhorias:** Qualquer sugestão é bem-vinda!

## Licença
Este projeto é distribuído sob a licença MIT. Veja o arquivo `LICENSE` para mais detalhes.
