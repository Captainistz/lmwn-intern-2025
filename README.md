# LMWN - Simple Go Project

Summarize COVID-19 stats using Go, Gin framwork.

## Installation

Clone the project

```bash
git clone https://github.com/Captainistz/lmwn-intern-2025.git
```

Go to the project directory

```bash
cd lmwn-intern-2025
```

Install dependencies using `make`

```bash
make install
```

Start the server (dev)

```bash
make dev
```

## Environment Variables

To custom this project, you will need to add the following environment variables to your .env file

`PORT`

`COVID_CASES_API_URL`

## API Reference

#### Get COVID-19 summary

```
GET /covid/summary
```

- **Description**: Retrieves a summarized data of COVID-19 cases.
- **Response**:
  - `200 OK`: Returns `Province` Object and `AgeGroup` Object
  - `500 Internal Server Error`: If an error occurs on the server.

<br/>
<br/>
<p align="center">Made with 🤍</p>
