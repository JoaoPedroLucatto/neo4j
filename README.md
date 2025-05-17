# 📘 Neo4j ETL + API

## 📌 Visão Geral

Este projeto consiste em dois microsserviços independentes e desacoplados, responsáveis por realizar a **carga de dados estruturados em CSV no banco de dados Neo4j** e por **expor uma API para consulta desses dados**. A arquitetura visa flexibilidade, manutenibilidade e facilidade de extensão futura.

![alt text](<./neo4j-api/docs/arquitetura.png>)

---

## 🧱 Microsserviços

### `neo4j-etl`
Responsável por:
- Leitura de arquivos CSV
- Conversão dos dados para o modelo de grafo
- Carga dos dados no Neo4j

### `neo4j-api`
Responsável por:
- Exposição de endpoints HTTP para consulta de dados no Neo4j
- Aplicação da lógica de negócio
- Implementação de regras de consulta e agregação

---

## 🧠 Abordagem Técnica

### Clean Architecture
O projeto foi estruturado seguindo os princípios da **Clean Architecture**, visando:
- Separação clara entre lógica de negócio, controle de fluxo e detalhes de infraestrutura

### Clean Code
A base de código segue práticas de legibilidade e padronização, como:
- Métodos com responsabilidades únicas
- Nomenclatura clara e objetiva
- Redução de dependências acopladas diretamente

---

## ⚙️ Processamento e Carga de Dados (`neo4j-etl`)

1. O serviço importa os arquivos `.csv`.
2. Os dados são lidos, convertidos para o formato apropriado de grafo (nós e relacionamentos), e enviados para o Neo4j.

### Melhorias planejadas:
- Verificar se o arquivo já foi importado.
- Leitura periódica automática do diretório onde os arquivos são salvos.
- Renomear o arquivo CSV após importação para evitar duplicidade.
---

## 🌐 API — Consulta de Dados (`neo4j-api`)

A API é responsável por receber requisições HTTP e traduzi-las em consultas Cypher ao banco Neo4j. As respostas são formatadas em JSON.

A API é documentada utilizando o padrão **OpenAPI (Swagger)**, facilitando testes e integração com outras aplicações.

### Enpoint API

A duas formas de testar os enpoint GET da API.

- Pelo arquivo **resquest.http**
- ou pelo arquivo **openapi.yml**

---

## 🔑 Decisões Técnicas Importantes

- **Separação de responsabilidades**: ETL e API são serviços independentes e escaláveis.
- **Desacoplamento via interface**: Permite trocar o banco Neo4j por outro ou até por uma API, modificando apenas a camada do banco de dados.

---

## 🚀 Subindo o Projeto

Precisa conter docker instalado na sua máquina.

```bash
make build up
```

Caso precise matar os containers e limpar o volume criado.

```bash
make clean
```

---

## 💡 Arquitetura Sugerida

A carga de dados será feita por meio de um endpoint POST /import, mantendo a API como única interface com o banco. Isso isola a infraestrutura e permite escalar os serviços de forma independente.

![alt text](<./neo4j-api/docs/arquitetura-sugerida.png>)

Haverá poucas mudanças no projeto atual. Como utilizamos uma **interface para abstração do banco de dados**, basta implementarmos um novo adaptador (por exemplo, uma integração via REST) para que o serviço `neo4j-etl` envie os dados via HTTP para a `neo4j-api`.  
Não será necessário reescrever toda a lógica de negócio — apenas criar uma nova camada de comunicação com o banco.