package shutdown

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Quit(ser *http.Server) (err error) {
	quit := make(chan os.Signal, 1)
	// NOTES: SIGINT(2) denotes signal sent by Ctrl+C, SIGTERM(15) denotes signal can be blocked,
	// SIGQUIT(3) denotes signal sent by Ctrl+\, SIGKILL(9) denotes signal non-catchable.
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err = ser.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}
	log.Println("Server exiting...")

	return
}
