# API para orçamento familiar - Challenge Alura

## O que é um challenge
São 4 semanas de desafios propostos pela plataforma de ensino Alura com o objetivo de praticar construindo um projeto. Toda semana são disponibilizados desafios e o aluno deve usar o material de apoio fornecido a cada semana para resolver o desafio proposto. 

### Projeto
Essa edição tem como objetivo construir uma api REST de controle financeiro. 

### Desafios de cada semana
<b>1ª semana</b> - CRUD de despesas e receitas e testes de api utilizando Postman

<b>2ª semana</b> - Categorização de despesas, filtro de despesas/receitas por descrição, listagem de despesas/receitas por mês, resumo do mês e testes automatizados

<b>3ª e 4ª semana</b> - Deploy e autenticação

## Tecnologias utilizadas
Go(Golang)<br>
Docker <br>
Gin Gonic <br>
gORM <br>
Heroku <br>

## URL Base
 > https://api-orcamento-familiar.herokuapp.com/
## Rotas

### Cadastro e Autenticação
| Rota | Método | Descrição | BODY PARAMS | QUERY PARAMS |
| --- | --- | --- | --- | --- |
| / | POST | Cadastra novo usuário para acessar todas as outras requisições | <pre>{<br>"email": "thay@thay.com",<br>"senha": "123456"<br>}</pre> | - |
| /login | POST | Retorna Bearer token obrigatório em todas as outras requisições | <pre>{<br>"email": "thay@thay.com",<br>"senha": "123456"<br>}</pre> | - |
 
 É necessário cadastrar e fazer o login antes de acessar as outras requisições do app.

### Receitas
| Rota | Método | Descrição | BODY PARAMS | QUERY PARAMS |
| --- | --- | --- | --- | --- |
| /receitas | POST | Cadastra uma receita | <pre> {<br> "descricao": "Salário",<br> "valor": 1200,<br> "data": "09/2022"<br>} </pre> | - |
| /receitas | GET | Retorna todas as receitas | - | descricao (opcional) |
| /receitas/{ano}/{mes} | GET | Retorna todas as receitas do mês | - | - |
| /receitas/{id} | GET | Retorna receita por id | - | - |
| /receitas/{id} | PUT | Atualiza receita por id | <pre> {<br> "descricao": "Salário",<br> "valor": 1200,<br> "data": "09/2022"<br>} </pre> | - |
| /receitas/{id} | DELETE | Remove receita por id | - | - |

### Despesas
| Rota | Método | Descrição | BODY PARAMS | QUERY PARAMS |
| --- | --- | --- | --- | --- |
| /despesas | POST | Cadastra uma despesa |  <pre> {<br> "descricao": "Netflix",<br> "valor": 55.90,<br> "data": "09/2022",<br> "categoria": "lazer"<br>} </pre> O campo categoria é opcional (ver ids correspondentes na tabela Categoria) | - |
| /despesas | GET | Retorna todas as despesas | - | descricao (opcional) |
| /despesas/{ano}/{mes} | GET | Retorna todas as despesas do mês | - | - |
| /despesas/{id} | GET | Retorna despesas por id | - | - |
| /despesas/{id} | PUT | Atualiza despesa por id |  <pre> {<br> "descricao": "Netflix",<br> "valor": 55.90,<br> "data": "09/2022",<br> "categoria": "lazer"<br>} </pre> O campo id_categoria é opciona (ver categorias correspondentes na tabela Categoria) | - |
| /despesas/{id} | DELETE | Remove despesa por id | - | - |

### Resumo
| Rota | Método | Descrição | BODY PARAMS | QUERY PARAMS |
| --- | --- | --- | --- | --- |
| /resumo/{ano}/{mes} | GET | Retorna resumo do mês | - | - |

### Categorias disponíveis
| Nome |
| --- |
| Alimentação |
| Saúde |
| Moradia |
| Transporte |
| Educação |
| Lazer |
| Imprevistos |
| Outras |

##Makefile 
```
Para fazer o build da aplicação usar o comando:
<code>make docker-image-build</code>

Para rodar o projeto, utilizar após o build o comando:
<code>make run</code>
