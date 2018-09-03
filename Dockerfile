FROM gcr.io/silvafabio-codelabs/cicdbase:0.1
ENV SOURCE=/go/src/github.com/viglesiasce/gke-info
COPY . $SOURCE
WORKDIR $SOURCE/cmd/gke-info
RUN go build -o gke-info
CMD ./gke-info
