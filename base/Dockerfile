# FROM asia-southeast1-docker.pkg.dev/coffee-dev-351011/coffee/api-builder AS builder

# Use the official Alpine image for a lean production container.
# https://hub.docker.com/_/alpine
# https://docs.docker.com/develop/develop-images/multistage-build/#use-multi-stage-builds
FROM alpine:3.15
RUN apk add --no-cache ca-certificates

# Copy the binary to the production image from the builder stage.
# COPY --from=builder /app/server /app/server

# Run the web service on container startup.
CMD ["/app/server"]
