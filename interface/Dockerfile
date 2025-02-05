FROM --platform=linux/amd64 node:18-alpine AS builder
WORKDIR /app
COPY . .
COPY package*.json ./

# Install dependencies
RUN npm ci
RUN npm run build

# Stage 2: Serve the SvelteKit app
FROM --platform=linux/amd64 node:18-alpine AS runner

#WORKDIR /app
COPY --from=builder /app/build ./build
COPY package*.json ./
RUN npm install --prod

EXPOSE 3000
CMD ["node", "build"]

#EXPOSE 4173
#CMD ["serve", "-s", "build", "-l", "4173"]
