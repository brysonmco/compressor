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
4. The api and database should be run via Docker Compose, use the `--build` flag to ensure your changes are included
```bash
compose up --build
```
5. Remember to initialize the database with the latest init from `./backend/migrations/sql`
6. Set GHCR credentials on the Compression Service server, make sure the token can read containers
```bash
export GHCR_USERNAME=
export GHCR_TOKEN=

```
7. The Compression Service is not designed to run in Docker as it manages containers itself, so you will need to run it
manually
```bash
go run ./compression-service/main.go
```

## Feature Tracker
The application is not complete, these features I plan to add later
- [ ] The entire application
- [ ] 2FA Verification Codes
- [ ] Social Login (Google and Apple)
- [ ] Progress Bar for Compression Jobs