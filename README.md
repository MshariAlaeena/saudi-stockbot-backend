# RAG Patient Chatbot

A multi-tenant medical assistant chatbot that uses Retrieval-Augmented Generation (RAG) to answer patient queries, extract information from documents/images, and schedule appointments.

## Features

* **Semantic Q\&A** via Pinecone
* **Multimodal Extraction** using llama-4-scout-17b-16e-instruct
* **Appointment Scheduling** integration (configurable per tenant)
* **Document Ingestion**: upload PDFs, images; extract and chunk text
* **Per-tenant Customization**: tone, prompt templates, calendar creds

## Tech Stack

* **Backend:** Go (Gin)
* **Embeddings and Vector DB:** Pinecone (integrated inference)
* **LLM API:** Groq (llama-3.3-70b-versatile)
* **LLM API:** Groq (llama-4-scout-17b-16e-instruct)

## Architecture

```
Client ⇄ Gin API ⇄ Services:
  • VectorDB    (Pinecone)
  • LLM         (Groq)
```

## Getting Started

### Prerequisites

* Go 1.20+
* Pinecone project (with integrated inference index)
* Groq API keys

### Installation

```bash
git clone https://github.com/your-org/rag-patient-chatbot.git
cd rag-patient-chatbot
go mod download
go build -o patient-chatbot ./cmd
```

### Configuration

Copy `.env.example` to `.env` and fill in with your values:

```dotenv
PINECONE_NAMESPACE=your_pinecone_namespace
PINECONE_API_KEY=…
PINECONE_INDEX=…
PINECONE_HOST=…
GROQ_API_KEY=…
LLM_MODEL=…
ARABIC_LLM_MODEL=…
MULTIMODAL_LLM_MODEL=…
DB_HOST=…
DB_PORT=…
DB_USER=…
DB_PASSWORD=…
DB_NAME=…
```

## Running

With the Makefile and `.env` in place, you have two options:

**1. Build & run via Makefile**

```bash
cp .env.example .env     # one-time only
make run                  # builds and runs the server
```

**2. Development mode (no binary)**

```bash
make run-dev              # loads .env and runs via `go run`
```

Server listens on **:8080** by default.

## API Endpoints

### Upload Document

```
POST /api/v1/upload
Content-Type: multipart/form-data
Fields:
  - org_id: UUID
  - file: binary
Response: 200 Success
{
  "doc_id": "<uuid>",
  "ingestion_status": "pending"
}
```

### Chat

```
POST /api/v1/chat
Content-Type: application/json
Body:
{
  "model": "<model-name>",
  "messages": [ ... ],
  "temperature": 1.0,
  "max_completion_tokens": 1024,
  "top_p": 1.0,
  "stream": false
}
Response: 200 OK
{
  "answer": "...",
  "sources": ["<chunk_id>", ...]
}
```

### Health Check

```
GET /api/v1/health
Response: 200 OK
{ "status": "ok" }
```

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/awesome`)
3. Commit your changes (`git commit -m "Add awesome feature"`)
4. Push to your branch (`git push origin feature/awesome`)
5. Open a Pull Request

## License

This project is licensed under the MIT License.
