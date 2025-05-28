# Compressor
Compress, transcode, and more

## Development Builds
1. Fill required values into `.env`
```bash
   cp .env.example .env
```
2. Create a Stripe application, put the secret key in `.env`
3. For development, you are going to want to run the frontend with NPM, not Docker
```bash
npm run dev
```
4. The backend and database should be run via Docker Compose, use the `--build` flag to ensure your changes are included
```bash
compose up --build
```
5. Remember to initialize the database with the latest init from `./backend/migrations/sql`
6. That is it for now!

## Feature Tracker
There is still way too much to do for one of these
- [ ] The entire application