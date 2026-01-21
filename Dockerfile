FROM golang:1.25 AS tool-builder
WORKDIR /app
COPY ./tools/bento-gallery-pre-runner/go.mod ./tools/bento-gallery-pre-runner/go.sum ./
RUN go mod download
COPY ./tools/bento-gallery-pre-runner/*.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o ./bento-gallery-pre-runner

# Use a Node.js Alpine image for the builder stage
FROM node:25-alpine AS site-builder
WORKDIR /app
COPY package*.json ./
RUN npm ci
COPY . .
ENV IMAGE_DIR="/app/static/images"
RUN npm run build
RUN npm prune --production

# Use another Node.js Alpine image for the final stage
FROM node:25-alpine
WORKDIR /app
COPY --from=site-builder /app/build build/
COPY --from=site-builder /app/node_modules node_modules/
COPY package.json .
EXPOSE 3000

COPY --from=tool-builder /app/bento-gallery-pre-runner tools/
RUN chmod +x tools/bento-gallery-pre-runner

ENV NODE_ENV=production
ENV IMAGE_DIR="/app/build/client/images"
ENV MEDIA_DIR="/media"
ENV STATIC_DIR="/app/build/client"
COPY docker-start.sh .
RUN chmod +x ./docker-start.sh
CMD [ "./docker-start.sh" ]