package cmd

import (
	"encoding/json"
	"log"
	"os"
	"os/signal"
	"syscall"
	
	"github.com/katatrina/go12-service/shared/datatype"
	
	"github.com/google/uuid"
	"github.com/nats-io/nats.go"
	"github.com/spf13/cobra"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var consumerCmd = &cobra.Command{
	Use:   "consumer",
	Short: "Start consumer",
}

var consumerIncreaseLikeCountRestaurantCmd = &cobra.Command{
	Use:   "increase-like",
	Short: "Start consumer increase like count restaurant",
	Run: func(cmd *cobra.Command, args []string) {
		dsn := os.Getenv("DB_DSN")
		dbMaster, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		
		if err != nil {
			log.Fatal("failed to connect database", err)
		}
		
		db := dbMaster.Debug()
		
		nc, err := nats.Connect(os.Getenv("NATS_URL"))
		
		if err != nil {
			log.Fatal("failed to connect nats", err)
		}
		
		nc.Subscribe(datatype.EvtUserLikedRestaurant, func(msg *nats.Msg) {
			type msgData struct {
				RestaurantID uuid.UUID `json:"restaurantID"`
				UserID       uuid.UUID `json:"userID"`
			}
			
			var data msgData
			
			if err := json.Unmarshal(msg.Data, &data); err != nil {
				log.Println("failed to unmarshal data", err)
				return
			}
			
			if err := db.Table("restaurants").
				Where("id = ?", data.RestaurantID).
				Update("liked_count", gorm.Expr("liked_count + 1")).Error; err != nil {
				log.Println("failed to update like count", err)
				return
			}
			
			log.Println("Update like count success for restaurantID:", data.RestaurantID)
		})
		
		// Setup graceful shutdown
		signalChan := make(chan os.Signal, 1)
		signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
		
		// Block until we receive a signal
		log.Println("Consumer started. Press Ctrl+C to exit...")
		<-signalChan
		
		log.Println("Shutting down consumer...")
		
		// Drain connection (process pending messages before closing)
		if err := nc.Drain(); err != nil {
			log.Printf("Error draining NATS connection: %v", err)
		}
		
		// Close NATS connection
		nc.Close()
		
		log.Println("Consumer shutdown complete")
	},
}

var consumerDecreaseLikeCountRestaurantCmd = &cobra.Command{
	Use:   "decrease-like",
	Short: "Start consumer decrease like count restaurant",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Start consumer decrease like count restaurant")
	},
}

func setupConsumerCmd() {
	consumerCmd.AddCommand(consumerIncreaseLikeCountRestaurantCmd)
	consumerCmd.AddCommand(consumerDecreaseLikeCountRestaurantCmd)
}
