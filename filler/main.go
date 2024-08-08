package main

import (
	"avito-test/internal/config"
	"avito-test/internal/models"
	"avito-test/internal/storage/house"
	"context"
	"fmt"
	"log"
	"math/rand"
	"sync"
)

func main() {
	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatal(err)
	}
	repo, err := house.NewRepository(cfg)
	if err != nil {
		log.Fatal(err)
	}

	houses := 100
	year := 2024
	dev := "mayor's office"
	ctx := context.WithValue(context.Background(), models.User{}, models.User{Role: models.ModeratorRole})
	wg := new(sync.WaitGroup)
	for i := 1; i < 51; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for j := 1; j < houses; j++ {
				log.Println(j)
				address := fmt.Sprintf("street # %d, %d, Moscow", i, j)
				house, err := repo.CreateHouse(ctx, models.HouseBuilder{
					Address:          &address,
					ConstructionDate: &year,
					Developer:        &dev,
				})
				if err != nil {
					log.Println(err)
					continue
				}
				flantNum := 100 + rand.Int()%50
				for k := 0; k < flantNum; k++ {
					price := 10000 + rand.Uint64()%3000
					rooms := 1 + byte(rand.Int()%5)
					_, err := repo.CreateFlat(ctx, models.FlatBuilder{
						HouseId: &house.UUID,
						Price:   &price,
						Rooms:   &rooms,
					}, "created")
					if err != nil {
						log.Println(err)
						break
					}
				}
			}
		}(i)
	}
	wg.Wait()
}
