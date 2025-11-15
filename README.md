# 🎬 MagicStreamMovies

**MagicStreamMovies** is a modern movie streaming web application built with **React**, **Go (Gin)**, and **MongoDB**. It allows users to browse movies, stream trailers, review films, and get **AI-powered personalized movie recommendations**.

---

## 🛠 Features

- **User Authentication**: Register, login, and logout securely.  
- **Movie Browsing**: View all movies with search functionality.  
- **Movie Reviews**: Submit and update reviews for movies.  
- **AI-Powered Recommendations**: Personalized movie suggestions using AI based on user activity and preferences.  
- **Streaming**: Stream movies directly via YouTube embedded links.  
- **Responsive UI**: Works on desktop and mobile devices.  
- **Secure API**: Backend built with Go (Gin) and MongoDB.  
- **Protected Routes**: Only authenticated users can access certain features.  

---

## 🧰 Tech Stack

| Frontend | Backend | Database | AI |
|----------|---------|----------|----|
| React, React Router v6 | Go (Gin Framework) | MongoDB Atlas | AI recommendation engine (e.g., TensorFlow, OpenAI API, or custom ML model) |
| Axios for HTTP requests | JWT for authentication | Mongoose driver for Go | Personalized recommendations based on user reviews and watch history |

---


---

## 🚀 Getting Started

### Prerequisites

- [Go](https://golang.org/dl/) >= 1.20  
- [Node.js & npm](https://nodejs.org/) >= 18  
- MongoDB Atlas account (or local MongoDB instance)  

---

### Backend Setup (Go + MongoDB)

1. Navigate to the backend folder:

```bash
cd server/movie-stream-api
```
2. Create a .env file with the following variables:

```bash
DATABASE_NAME = movie-stream
MONGODB_URI=<Your_database_address>
OPENAI_API_KEY=<your_OPENAI_secret_key_here>
SECRET_KEY=<your_secret_key_here>
SECRET_REFRESH_KEY=<your_refresh_secret_key_here>
BASE_PROMPT_TEMPLATE=Return a response using one of these words: {rankings}. The response should be a single word and should not contain any other text. The response should be based on any of the following review: 
RECOMENDED_MOVIE_LIMIT=5
ALLOWED_ORIGINS=http://localhost:3000,http://localhost:5173,http://localhost:8080
GIN_MODE=release
```

3. Install dependencies::

```bash
go mod tidy
```

4. Run the backend server:

```bash
go run main.go
```
- Server runs on http://localhost:8080 by default.

### Frontend Setup (React)

1. Navigate to the frontend folder:

```bash
cd cd client
```
2. Create a .env file with the following variables:

```bash
VITE_API_BASE_URL=http://localhost:8080
```

3. Install dependencies::

```bash
npm install
```

4. Start the React development server:

```bash
npm run dev
```
- Frontend runs on http://localhost:5173

## Author

[Paul Nnaji](https://github.com/paularinzee)

## License

[MIT](./LICENSE)
