# 🚀 desafio_stress

Um **load tester (teste de carga)** simples em Go, desenvolvido para estressar endpoints HTTP com múltiplas requisições simultâneas. Ideal para medir desempenho, latência e distribuição de códigos de resposta de serviços web.

---

## 📦 Sobre o Projeto

Este utilitário permite que você:

- Defina o número total de requisições
- Controle o nível de concorrência
- Obtenha métricas básicas como tempo total e distribuição de status HTTP

---

## 🧰 Tecnologias

- [Go (Golang)](https://golang.org/) 1.24+
- Docker (opcional)

---

## ⚙️ Como usar

### ✅ Pré-requisitos

- Go instalado (1.24+)
- OU Docker

### 🏃‍♂️ Execução local

```bash
go run main.go --url=http://localhost:8080 --requests=100 --concurrency=10
````

Parâmetros:

| Parâmetro       | Descrição                                   |
| --------------- | ------------------------------------------- |
| `--url`         | URL do endpoint a ser testado               |
| `--requests`    | Número total de requisições                 |
| `--concurrency` | Número de requisições simultâneas (workers) |

---

## 🐳 Usando via Docker

### Build da imagem

```bash
docker build -t loadtester .
```

### Execução

```bash
docker run --rm loadtester \
  --url=http://localhost:8080 \
  --requests=100 \
  --concurrency=10
```

---

## 📊 Exemplo de saída

```txt
Iniciando teste de carga:
URL: http://localhost:8080
Requests: 100
Concorrência: 10

Relatório de Teste:
Tempo total: 3.20121463s
Total de requests: 100
HTTP 200: 100
Distribuição de status HTTP:
  200: 100
```

---

## 📂 Estrutura

```
.
├── Dockerfile         # Build container
├── go.mod             # Gerenciamento de dependências Go
├── LICENSE            # MIT License
├── main.go            # Código-fonte do load tester
└── README.md          # Este arquivo
```

---

## 📄 Licença

Distribuído sob a licença MIT. Veja `LICENSE` para mais informações.

---

## ✍️ Autor

Alefe Serafim — [LinkedIn](https://www.linkedin.com/in/alefeserafim)
