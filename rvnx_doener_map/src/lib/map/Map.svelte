<script>
    import {onMount} from 'svelte';
    import {browser} from '$app/env';
    import ClusterPopUp from "./ClusterPopUp.svelte";
    import ShopPopUp from "./ShopPopUp.svelte";

    let leaflet;

    const biggestMarkerRadius = 50

    let map;
    let lastMarkers = []

    // Create a popup with a Svelte component inside it and handle removal when the popup is torn down.
    // `createFn` will be called whenever the popup is being created, and should create and return the component.
    async function bindPopup(marker, createFn) {
        let popupComponent;
        marker.bindPopup(() => {
            let container = leaflet.DomUtil.create('div');
            popupComponent = createFn(container);
            return container;
        });

        marker.on('popupclose', () => {
            if (popupComponent) {
                let old = popupComponent;
                popupComponent = null;
                // Wait to destroy until after the fadeout completes.
                setTimeout(() => {
                    old.$destroy();
                }, 500);

            }
        });
    }

    async function createClusterMarker(cluster) {
        const marker = leaflet.circleMarker([cluster.lat, cluster.lng])
            .setRadius(cluster.norm * biggestMarkerRadius)

        bindPopup(marker, (m) => {
            return new ClusterPopUp({
                target: m,
                props: {
                    cluster: cluster,
                    zoomCallback: (bounds) => {
                        console.log(bounds)

                        map.fitBounds([
                            [bounds.min_lat, bounds.min_lng],
                            [bounds.max_lat, bounds.max_lng]
                        ])
                    }
                }
            });
        });

        return marker
    }

    async function createShopMarker(shop) {
        const marker = leaflet.marker([shop.lat, shop.lng])

        bindPopup(marker, (m) => {
            return new ShopPopUp({
                target: m,
                props: {
                    shopID: shop.id,
                }
            });
        });

        return marker
    }

    async function loadShops() {
        const bounds = map.getBounds();
        const ne = bounds.getNorthEast();
        const sw = bounds.getSouthWest();

        fetch("/api/v1/kebabshops/auto?ltm=" + sw.lat + "&ltx=" + ne.lat + "&lnm=" + sw.lng + "&lnx=" + ne.lng)
            .then(resp => resp.json())
            .then(data => {
                if (data.clusters || data.cords) {
                    for (const s of lastMarkers) {
                        map.removeLayer(s)
                    }
                    lastMarkers = []

                    if (data.clusters) {
                        data.clusters.forEach(async (cluster) => {
                            const marker = await createClusterMarker(cluster);
                            marker.addTo(map);

                            lastMarkers.push(marker);
                        })
                    }

                    if (data.cords) {
                        data.cords.forEach(async (shop) => {
                            const marker = await createShopMarker(shop);
                            marker.addTo(map);

                            lastMarkers.push(marker);
                        })
                    }
                }
            })
    }

    onMount(async () => {
        if (browser) {
            leaflet = await import("leaflet")

            map = leaflet.map('map').setView(
                [50.194036, 10.423515], 6, {
                    renderer: leaflet.svg()
                }
            );

            // limit to "single" earth map
            map.setMaxBounds([
                [105, 210],
                [-105, -190],
            ]);

            leaflet.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
                attribution: '&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
            }).addTo(map);

            // zoom to user's location
            map.locate({setView: true, maxZoom: 13});

            map.on("zoomend", loadShops)
            map.on("moveend", loadShops)
            loadShops()
        }
    });
</script>

<div class="w-auto h-full" id="map"></div>