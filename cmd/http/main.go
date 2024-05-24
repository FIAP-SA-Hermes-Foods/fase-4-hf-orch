package main

import (
	"context"
	l "fase-4-hf-orch/external/logger"
	httpExt "fase-4-hf-orch/internal/adapters/driver/http"
	clientrpc "fase-4-hf-orch/internal/adapters/driver/rpc/client"
	orderrpc "fase-4-hf-orch/internal/adapters/driver/rpc/order"
	productrpc "fase-4-hf-orch/internal/adapters/driver/rpc/product"
	voucherrpc "fase-4-hf-orch/internal/adapters/driver/rpc/voucher"
	"fase-4-hf-orch/internal/core/application"
	httpH "fase-4-hf-orch/internal/handler/http"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/marcos-dev88/genv"
)

func init() {
	if err := genv.New(); err != nil {
		l.Errorf("error set envs %v", " | ", err)
	}
}

func main() {

	router := http.NewServeMux()

	ctx := context.Background()

	urlAPI := fmt.Sprintf("http://%s:%s/%s",
		os.Getenv("MERCADO_PAGO_API_HOST"),
		os.Getenv("MERCADO_PAGO_API_PORT"),
		os.Getenv("MERCADO_PAGO_API_URI"),
	)

	headersAPI := map[string]string{
		"Content-type": "application/json",
	}

	du, err := time.ParseDuration(os.Getenv("MERCADO_PAGO_API_TIMEOUT"))

	if err != nil {
		log.Fatalf("error: %v", err)
	}

	paymentAPI := httpExt.NewMercadoPagoAPI(urlAPI, headersAPI, du)
	clientRPC := clientrpc.NewClientRPC(ctx, os.Getenv("HOST_CLIENT"), os.Getenv("PORT_CLIENT"))
	orderRPC := orderrpc.NewOrderRPC(ctx, os.Getenv("HOST_ORDER"), os.Getenv("PORT_ORDER"))
	productRPC := productrpc.NewProductRPC(ctx, os.Getenv("HOST_PRODUCT"), os.Getenv("PORT_PRODUCT"))
	voucherRPC := voucherrpc.NewVoucherRPC(ctx, os.Getenv("HOST_PRODUCT"), os.Getenv("PORT_PRODUCT"))

	app := application.NewApplication(clientRPC, orderRPC, productRPC, voucherRPC, paymentAPI)

	h := httpH.NewHandler(app)

	router.Handle("/hermes_foods/health", http.StripPrefix("/", httpH.Middleware(h.HealthCheck)))
	router.Handle("/hermes_foods/client/", http.StripPrefix("/", httpH.Middleware(h.HandlerClient)))
	router.Handle("/hermes_foods/client", http.StripPrefix("/", httpH.Middleware(h.HandlerClient)))

	router.Handle("/hermes_foods/order/", http.StripPrefix("/", httpH.Middleware(h.HandlerOrder)))
	router.Handle("/hermes_foods/order", http.StripPrefix("/", httpH.Middleware(h.HandlerOrder)))

	router.Handle("/hermes_foods/product/", http.StripPrefix("/", httpH.Middleware(h.HandlerProduct)))
	router.Handle("/hermes_foods/product", http.StripPrefix("/", httpH.Middleware(h.HandlerProduct)))

	router.Handle("/hermes_foods/voucher/", http.StripPrefix("/", httpH.Middleware(h.HandlerVoucher)))
	router.Handle("/hermes_foods/voucher", http.StripPrefix("/", httpH.Middleware(h.HandlerVoucher)))

	log.Fatal(http.ListenAndServe(":"+os.Getenv("API_HTTP_PORT"), router))
}
