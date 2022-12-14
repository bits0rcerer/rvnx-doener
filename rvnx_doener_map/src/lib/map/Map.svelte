<script>
	import { onMount } from 'svelte';
	import { browser } from '$app/environment';
	import ClusterPopUp from './ClusterMarkerPopup.svelte';
	import ShopPopUp from './ShopMarkerPopup.svelte';
	import { isMobile, isApple } from '../common/device-detection.js';
	import KebabFilter from './KebabFilter.svelte';

	let leaflet;
	let leaflet_ui;

	const biggestMarkerRadius = 50;

	let rvnxFilter = false;
	let map;
	let lastMarkers = [];
	let suspendMarkerReload = false;

	// Create a popup with a Svelte component inside it and handle removal when the popup is torn down.
	// `createFn` will be called whenever the popup is being created, and should create and return the component.
	async function bindPopup(marker, createFn) {
		let popupComponent;
		marker.bindPopup(() => {
			let container = leaflet.DomUtil.create('div');
			popupComponent = createFn(container);
			return container;
		});

		marker.on('popupopen', () => {
			suspendMarkerReload = true;
			setTimeout(() => {
				suspendMarkerReload = false;
			}, 1000);
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
		const marker = leaflet
			.circleMarker([cluster.lat, cluster.lng])
			.setRadius(cluster.norm * biggestMarkerRadius);

		bindPopup(marker, (m) => {
			return new ClusterPopUp({
				target: m,
				props: {
					cluster: cluster,
					zoomCallback: (bounds) => {
						suspendMarkerReload = false;
						map.fitBounds([
							[bounds.min_lat, bounds.min_lng],
							[bounds.max_lat, bounds.max_lng]
						]);
					}
				}
			});
		});

		return marker;
	}

	async function createShopMarker(shop) {
		const marker = leaflet.marker([shop.lat, shop.lng]);

		bindPopup(marker, (m) => {
			return new ShopPopUp({
				target: m,
				props: {
					shopID: shop.id
				}
			});
		});

		return marker;
	}

	let loadingShops = false;

	async function loadShops() {
		if (suspendMarkerReload || loadingShops) return;
		loadingShops = false;

		const bounds = map.getBounds();
		const ne = bounds.getNorthEast();
		const sw = bounds.getSouthWest();

		fetch(
			'/api/v1/kebabshops/auto?ltm=' +
				sw.lat +
				'&ltx=' +
				ne.lat +
				'&lnm=' +
				sw.lng +
				'&lnx=' +
				ne.lng +
				'&rvnx_only=' +
				rvnxFilter +
				(isMobile() ? '&cc=8&ct=30' : '')
		)
			.then((resp) => resp.json())
			.then((data) => {
				if (data.clusters || data.cords) {
					for (const s of lastMarkers) {
						map.removeLayer(s);
					}
					lastMarkers = [];

					if (data.clusters) {
						data.clusters.forEach(async (cluster) => {
							const marker = await createClusterMarker(cluster);
							marker.addTo(map);

							lastMarkers.push(marker);
						});
					}

					if (data.cords) {
						data.cords.forEach(async (shop) => {
							const marker = await createShopMarker(shop);
							marker.addTo(map);

							lastMarkers.push(marker);
						});
					}
				}
			})
			.finally(() => {
				loadingShops = false;
			});
	}

	onMount(async () => {
		if (browser) {
			leaflet = await import('leaflet');
			leaflet_ui = await import('leaflet-ui');

			leaflet.Control.KebabFilterControl = leaflet.Control.extend({
				onAdd: () => {
					var container = leaflet.DomUtil.create('div', 'leaflet-bar');
					const kf = new KebabFilter({
						target: container,
						props: {
							checked: rvnxFilter,
							onChange: (e) => {
								if (rvnxFilter != e.target.checked) {
									rvnxFilter = e.target.checked;
									loadShops();
								}
							}
						}
					});
					return container;
				}
			});

			leaflet.control.kebabFilterControl = function (opts) {
				return new leaflet.Control.KebabFilterControl(opts);
			};

			map = leaflet.map('map', {
				// Optional customizations
				mapTypeId: 'streets',
				mapTypeIds: ['streets', 'satellite'],
				gestureHandling: false,
				zoomControl: true,
				pegmanControl: true,
				locateControl: true,
				fullscreenControl: false,
				minimapControl: false,
				editInOSMControl: false,
				loadingControl: true,
				disableDefaultUI: false,
				rotate: true,

				layersControl: {
					inline: true,
					position: 'bottomleft'
				},

				searchControl: {
					markerLocation: false,
					autoCollapse: false,
					position: 'topleft'
				}
			});

			leaflet.control
				.kebabFilterControl({
					position: 'topright'
				})
				.addTo(map);

			// limit to "single" earth map
			map.setMaxBounds([
				[105, 210],
				[-105, -190]
			]);

			if (isApple().isSafari) {
				// Safari doing weird stuff...
				map.setView([50.194036, 10.423515], 6, {
					renderer: leaflet.svg()
				});
			} else {
				// zoom to user's location
				map.locate({ setView: true });
			}

			map.on('locationfound', () => {
				// location found
			});
			map.on('locationerror', (e) => {
				console.log('location denied: ', e);
				map.setView([50.194036, 10.423515], 6, {
					renderer: leaflet.svg()
				});
			});

			map.on('moveend', loadShops);
			loadShops();
		}
	});
</script>

<div class="w-auto h-full" id="map" />
