# First stage: build the executable.
FROM golang:1.16 AS builder

# Set the working directory outside $GOPATH to enable the support for modules.
WORKDIR /src

# Fetch dependencies first; they are less susceptible to change on every build
# and will therefore be cached for speeding up the next build
COPY ./go.mod ./go.sum ./
RUN go mod download

# Import the code from the context.
COPY ./ ./

# Build the executable to `/app`. Mark the build as statically linked.
RUN CGO_ENABLED=0 go build \
  -installsuffix 'static' \
  -o /kouyi_palm_treasure .

# Final stage: the running container.
FROM scratch AS final

# Import the compiled executable from the first stage.
COPY --from=builder /kouyi_palm_treasure /kouyi_palm_treasure

COPY ./configs ./

# Declare the port on which the webserver will be exposed.
EXPOSE 8080

# Run the compiled binary.
ENTRYPOINT ["/kouyi_palm_treasure"]