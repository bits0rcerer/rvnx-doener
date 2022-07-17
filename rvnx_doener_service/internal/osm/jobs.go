package osm

import (
	"log"
	"rvnx_doener_service/internal/services"
)

func SyncOSMKebabShops(kebabShopService *services.KebabShopService) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("[!] Sync OSM kebab shops job failed: ", r)
		}
	}()

	shops, err := GetOSMKebabShops()
	if err != nil {
		panic(err)
	}

	for _, s := range shops {
		_, err := kebabShopService.UpdateOrInsertKebabShop(&s)
		if err != nil {
			panic(err)
		}
	}

	log.Println("[*] Sync OSM kebab shops job done")
}
