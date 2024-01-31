
### Documentação do Projeto CRUD Go MVC

Este é um projeto Go (Golang) criado durante a aprendizagem da linguagem Go, seguindo a arquitetura MVC para operações CRUD em um banco de dados MongoDB.

### Instalação

1. **Clone o Repositório:**
   ```bash
   git clone git@github.com:MelvimJones/SimpleGoCRUD.git
   cd crudGo_mvc
   ```

2. **Configuração do Ambiente:**
   - Certifique-se de ter o Go instalado em sua máquina.
   - Instale as dependências do projeto:
     ```bash
     go get
     ```

3. **Configuração do Banco de Dados:**
   - Configure a URI do MongoDB no arquivo `.env`.
   - Sim, é uma boa prática incluir instruções sobre como criar o arquivo `.env` no seu README. Aqui estão algumas instruções simples que você pode adicionar:

### Criando o Arquivo `.env`

O arquivo `.env` é utilizado para armazenar variáveis de ambiente sensíveis, como chaves de API ou configurações específicas do ambiente. Siga estas etapas para criar o arquivo `.env` para o seu projeto:

1. **Navegue até a Pasta do Projeto:**
   Certifique-se de estar no diretório do seu projeto antes de criar o arquivo `.env`.

   ```bash
   cd crudGo_mvc\crudapp
   ```

2. **Crie o Arquivo `.env`:**
   Utilize o seu editor de texto preferido para criar um novo arquivo chamado `.env`.

   ```bash
   touch .env
   ```

   Se estiver usando o Windows e não tiver o comando `touch`, pode criar manualmente usando o Bloco de Notas ou um editor de texto similar.

3. **Adicione Configurações:**
   Abra o arquivo `.env` no seu editor de texto e adicione as variáveis de ambiente necessárias. Por exemplo:

   ```env
   MONGODB_URI=your_mongodb_uri_here
   ```

   Substitua `your_mongodb_uri_here` pela URI real do seu banco de dados MongoDB.

4. **Salve o Arquivo:**
   Salve as alterações no arquivo `.env`.


### Utilização

1. **Executar a Aplicação:**
   ```bash
   go run main.go
   ```
   - O servidor estará disponível em http://localhost:8080.

2. **Endpoints CRUD:**

   - **Listar Todos os Produtos:**
     ```bash
     GET http://localhost:8080/products
     ```

   - **Criar um Novo Produto:**
     ```bash
     POST http://localhost:8080/products/create
     ```

   - **Atualizar um Produto:**
     ```bash
     PUT http://localhost:8080/products/update?id=ID_DO_PRODUTO
     ```

   - **Deletar um Produto:**
     ```bash
     DELETE http://localhost:8080/products/delete?id=ID_DO_PRODUTO
     ```

### Exemplos de Solicitações HTTP

- **Listar Todos os Produtos:**
  ```http
  GET http://localhost:8080/products
  ```

- **Criar um Novo Produto:**
  ```http
  POST http://localhost:8080/products/create
  Content-Type: application/json

  {
    "descricao": "Nome do Produto",
    "preco": 19.99,
    "imagem": "url_da_imagem",
    "quant": 10
  }
  ```

- **Atualizar um Produto:**
  ```http
  PUT http://localhost:8080/products/update?id=ID_DO_PRODUTO
  Content-Type: application/json

  {
    "descricao": "Novo Nome do Produto",
    "preco": 29.99,
    "imagem": "nova_url_da_imagem",
    "quant": 15
  }
  ```

- **Deletar um Produto:**
  ```http
  DELETE http://localhost:8080/products/delete?id=ID_DO_PRODUTO
  ```

### Observações

- Certifique-se de ter um MongoDB em execução e configurado corretamente.
- Substitua `ID_DO_PRODUTO` pelos IDs reais dos produtos ao fazer atualizações ou exclusões.

