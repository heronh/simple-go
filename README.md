Repositórios para referência pessoal com objetivo de auto aprendizado

## Postgres
- Comando para criar o servidor:
  - docker run --name db -e POSTGRES_PASSWORD=mysecretpassword -d postgres
- Acessar servidor postgres:
  - docker run -it --rm --network some-network postgres psql -h some-postgres -U postgres

## postgres
- Projeto com exemplo de conexão ao postgres
- Conexão baseada em variáveis declaradas no arquivo .env

## simple-package
- Usa o projeto anterior, simple, como base
- Separa server em um pacote diferente
- Cria um diretório para server
  
## simple
- Exibe o conteúdo de um arquivo html enviando um título que será substituido na página.
- Carrega página de estilos, bootstrap, armazenada localmente.
