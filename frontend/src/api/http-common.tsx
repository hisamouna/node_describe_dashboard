import axios from 'axios';

const server = axios.create({
  baseURL: "http://localhost:8090/v1",
  headers: {
    "Content-type": "application/json",
  }
})

const httpObject = {
  server
}

export default httpObject
