Repositórios para referência pessoal com objetivo de auto aprendizado

## postgres
- Projeto com exemplo de conexão ao postgres
- Conexão baseada em variáveis declaradas no arquivo .env
    
Criar um container postgeres:
docker run --name some-postgres -e POSTGRES_PASSWORD=mysecretpassword -d postgres

Acessar um container:
docker exec -it some-postgres /bin/bash


CREATE USER nome_do_usuario;


## simple-package
- Usa o projeto anterior, simple, como base
- Separa server em um pacote diferente
- Cria um diretório para server
  
## simple
- Exibe o conteúdo de um arquivo html enviando um título que será substituido na página.
- Carrega página de estilos, bootstrap, armazenada localmente.
