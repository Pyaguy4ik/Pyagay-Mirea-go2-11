package main

import (
    "log"
    "net/http"
    "os"

    "example.com/pz11-graphql/graph"
    "github.com/99designs/gqlgen/graphql/handler"
    "github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8081"

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = defaultPort
    }

    srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

    http.Handle("/", playground.Handler("GraphQL playground", "/query"))
    http.Handle("/query", srv)

    // Слушаем на всех интерфейсах для доступа извне
    addr := "0.0.0.0:" + port
    log.Printf("GraphQL server started on %s", addr)
    log.Printf("Playground available at http://<server-ip>:%s/", port)
    log.Fatal(http.ListenAndServe(addr, nil))
}
