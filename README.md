Repositórios para referência pessoal com objetivo de auto aprendizado

## jwt 
- Teste de autenticação e login usando gin e jwt
- Para criar um usuário:
  - curl -X POST -H "Content-Type: application/json" -d '{"email":"heron@gmail.com","password":"123456"}' localhost:4004/signup
- Login:
  - curl -X POST -H "Content-Type: application/json" -d '{"email":"heron@gmail.com","password":"123456"}'  localhost:4004/login
- Validar uma rota:
  - curl -X GET localhost:4004/validate 
## Templates
- Teste com template
- Executa a migração criando uma tabela de usuários 
- Preenche dados em uma tabela de usuários

## postgres
- Projeto com exemplo de conexão ao postgres
- Conexão baseada em variáveis declaradas no arquivo

## simple-package
- Usa o projeto anterior, simple, como base
- Separa server em um pacote diferente
- Cria um diretório para server
  
## simple
- Exibe o conteúdo de um arquivo html enviando um título que será substituido na página.
- Carrega página de estilos, bootstrap, armazenada localmente.
