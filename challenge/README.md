# Backend Assessment

Olá! 🖖🏽

Nossa intenção é, através deste (breve) desafio, avaliar a habilidade técnica percebida ao empregar e desenvolver uma solução para o problema aqui descrito.

## Domínio Problema

Uma instituição financeira contratou os serviços da T10 buscando maior **agilidade dos dados** através da metrificação de processos que, até então, não eram _observados_ (apropriadamente). Um dos processos é a solicitação do produto débito automático de empresas parceiras.
A operação é realizada manualmente e vai ser automatizada por este serviço, que vai permitir que outros serviços consumam, de forma livre, de seus eventos operacionais.

# Escopo

## Casos de Uso

As entidades conhecidas são:

- `ExternalApp`, representa uma aplicação externa e
- `Customer`, identificado por `customer_mid`, representa um cliente de `ExternalApp`
- `SuperUser`, representa um analista da mesa de integração

Glossário:

- "Solicitação de ativação" é traduzido para "Activation request"

### 1. Acesso

**Premissa**: Dado que um `ExternalApp` ou `SuperUser` possui um conjunto de credenciais de acesso válido, um novo token é gerado

- Dado que um novo token é gerado, então a lista de tokens ativos é atualizada

**~Premissa**: Dado que um `ExternalApp` ou `SuperUser` não possui um conjunto de credenciais de acesso válido, um erro é retornado e nenhum token é gerado

### 2. Acesso à recursos

**Premissa**: Dado que um `ExternalApp` ou `SuperUser` possui um token ativo e possui permissão para acessar um recurso específico, a ação é executada

Relação de acesso:

1. Commands

   - `RequestToken: ExternalApp, SuperUser`
   - `IssueProductActivation: ExternalApp`
   - `RejectActivation: SuperUser`
   - `ApproveActivation: SuperUser`

1. Read Model
   - `ActivationRequests: ExternalApp, SuperUser`

**~Premissa**: Dado que um `ExternalApp` ou `SuperUser` possui um token ativo, **não** possui permissão para acessar um recurso específico, um erro é retornado à aplicação e nenhuma ação é executada

**~Premissa**: Dado que um `ExternalApp` **não** possui um token ativo e solicita acesso à um recurso, um erro é retornado à aplicação e nenhuma ação é executada

### 2. Solicitação

**Premissa**: Dado que um `ExternalApp` possui um token ativo, permissão para solicitar uma ativação de produto e solicita uma ativação para o `customer_mid`, então uma solicitação de ativação é despachada

- Dado que uma solicitação de ativação é despachada, então um email de confirmação é enviada ao `customer_mid`

### 3. Avaliação

**Premissa**: Dado que um `SuperUser` possui um token ativo, permissão para avaliar uma ativação de produto e

1. rejeita uma determinada ativação, então o cancelamento desta ativação é despachado
   - Dado que um cancelamento de uma ativação é despachado, então o read model de solicitações é atualizada
   - Dado que um cancelamento de uma ativação é despachado, então um email de cancelamento é enviada ao `customer_mid`
1. aprova uma determinada ativação, então a aprovação desta ativação é despachada
   - Dado que uma aprovação de uma ativação é despachada, então o read model de solicitações é atualizada
   - Dado que uma aprovação de uma ativação é despachada, então um email é enviada ao `customer_mid`

Diagrama do [modelo de eventos](img/model.jpg). Note que é uma representação do domínio _exclusivamente_.

## Requisitos

Especifica o contexto em que a aplicação será operacionalizada

### Não funcionais

1. 30 empresas parceiras
1. 10 super-users
1. 1M reqs/dia
1. Eventos operacionais disponibilizados em streams para consumo externo

### Funcionais

#### Tecnologias

- implementação: `golang | elixir | python`
- armazenamento: `postgres | mongodb`
- broker: `kafka | rabbitmq`

#### Protocolos

- pontos de entrada: `http`
- autenticação: `simple jwt`

#### Padrões

Preferencialmente:

- arquitetural: `cqrs & hexagonal`
- design: `ddd & solid`

Bonus points:

- message bus as stream

### 3rd parties

O uso de bibliotecas externas é **livre**.

### Deployment

A forma como a aplicação será disponibilizada é **livre**. Fica a critério do candidato, por exemplo, usar algum PaaS a fim de reduzir a complexidade bem como utilizar receitas prontas através de ferramentas de automatização e.g. `ansible+dockercompose`.

No entanto, é esperado bom senso na documentação caso sejam usadas soluções @ `localhost`.

# Entrega

A _Release_ 0.1 🚀 consiste na implementação de um servidor web que implementa os casos de uso listados acima respeitando os requisitos funcionais e não funcionais. Fica a critério do desenvolvedor como os testes serão escritos, os scripts de _data migration_, os _schemas_ de entrada e saída da api e todas as outras definições que não foram listadas neste documento.

## Avaliação

Critérios ordenados por ordem de peso decrescente:

1. Correção (_correctness_) da solução

   - a fim de solucionar o [domínio-problema](#domínio-problema)
   - a fim de cumprir os [casos de uso](#casos-de-uso)
   - ao implementar os [requisitos](#requisitos) especificados

1. Testes
1. Organização, documentação e clareza na estruturação do projeto
1. Estilo, legibilidade e simplicidade no código
1. Escolhas e uso de 3rd parties
1. Padrões de segurança

#### Bonus points 🏆

1. Teste de stress
1. Boas práticas na modelagem e armazenamento de dados

## Eliminatórios

1. Copiar ou "se inspirar" em código alheio é _veementemente_ vetado ✋

## Submissão

Ao finalizar a implementação, o diretório da solução pode ser submetido de duas formas:

1. através de um _fork_ e um _pull request_ neste repositório ou
1. por email, compactado, para `it@t10.digital` com o assunto `Backend Assessment`

Feito 🤘
