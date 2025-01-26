# Task Manager

Task Manager é uma aplicação para gerenciar tarefas, construída com Go, Gin e MongoDB.

## Pré-requisitos

- Docker
- Docker Compose

## Configuração

1. Clone o repositório:

    ```sh
    git clone https://github.com/igorpdev/taskmanager.git
    cd taskmanager
    ```

2. Configure o arquivo `config.yaml` com as configurações desejadas:

    ```yaml
    server:
      port: 8080

    database:
      uri: mongodb://mongo:27017
      name: taskmanager
    ```

## Executando a aplicação

1. Construa e inicie os contêineres Docker:

    ```sh
    docker-compose -f deploy/docker-compose.yaml up --build 
    ```

2. Acesse a aplicação em [http://localhost:8080](http://localhost:8080).

## Endpoints

- `GET /tasks` - Lista todas as tarefas
- `POST /tasks` - Cria uma nova tarefa
- `GET /tasks/:id` - Obtém uma tarefa pelo ID
- `PUT /tasks/:id` - Atualiza uma tarefa pelo ID
- `DELETE /tasks/:id` - Deleta uma tarefa pelo ID
- `GET /metrics` - Retorna as métricas da aplicação (usado para monitoramento)

## Observações

Certifique-se de que o endpoint `/metrics` está habilitado no código e configurado corretamente no ambiente para expor as métricas da aplicação.