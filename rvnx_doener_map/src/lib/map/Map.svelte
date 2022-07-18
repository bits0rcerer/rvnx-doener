<script>
    import {onMount} from 'svelte';
    import {browser} from '$app/env';

    const biggestMarkerRadius = 50

    let map;
    let lastMarkers = []

    async function loadShops() {
        const leaflet = await import('leaflet');

        const bounds = map.getBounds();
        const ne = bounds.getNorthEast();
        const sw = bounds.getSouthWest();

        fetch("/api/v1/kebabshops/combined?ltm=" + sw.lat + "&ltx=" + ne.lat + "&lnm=" + sw.lng + "&lnx=" + ne.lng)
            .then(resp => resp.json())
            .then(data => {
                if (data.clusters || data.cords) {
                    for (const s of lastMarkers) {
                        map.removeLayer(s)
                    }
                    lastMarkers = []

                    if (data.clusters) {
                        data.clusters.forEach((cluster) => {
                            lastMarkers.push(
                                leaflet.circleMarker([cluster.lat, cluster.lng])
                                    .setRadius(cluster.norm * biggestMarkerRadius)
                                    .addTo(map)
                                    .bindPopup(cluster.shops))
                        })
                    }

                    if (data.cords) {
                        data.cords.forEach((shop) => {
                            lastMarkers.push(
                                leaflet.marker([shop.lat, shop.lng])
                                    .addTo(map)
                                    .bindPopup(shop.id))
                        })
                    }
                }
            })
    }

    onMount(async () => {
        if (browser) {
            const leaflet = await import('leaflet');

            map = leaflet.map('map').setView(
                [50.194036, 10.423515], 6, {
                    renderer: leaflet.svg()
                }
            );

            // limit to "single" earth map
            map.setMaxBounds([
                [85, 190],
                [-85, -170],
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


<main>
    <div id="map"></div>
</main>

<style>
    #map {
        height: 100vh;
        width: 100vw;
    }
</style>