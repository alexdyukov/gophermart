module github.com/alexdyukov/gophermart

go 1.18

replace github.com/alexdyukov/gophermart/internal/accrualsvc => ./internal/accrualsvc

replace github.com/alexdyukov/gophermart/internal/gophermartsvc => ./internal/accrualsvc

replace github.com/alexdyukov/gophermart/internal/storage => ./internal/storage

replace github.com/alexdyukov/gophermart/internal/handler => ./internal/handler

replace github.com/alexdyukov/gophermart/internal/config => ./internal/config
