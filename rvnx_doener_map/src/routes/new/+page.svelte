<script>
	import { currentUserStore, notificationContextStore } from '../../stores.js';
	import { onMount } from 'svelte';
	import { browser } from '$app/environment';
	import { goto } from '$app/navigation';
	import AutoComplete from 'simple-svelte-autocomplete';

	let notificationContext;
	notificationContextStore.subscribe((value) => {
		notificationContext = value;
	});

	let leaflet;
	let leaflet_ui;
	let map;
	let lastMarker;

	let anonymous = false;
	let selectedShopName;
	let selectedAddress;
	const apiKey = 'pk.776a0b34a86e5f7c30e80d67bde72661';

	async function searchAddress(search) {
		if (search && search.length > 3) {
			const res = await fetch(
				'https://eu1.locationiq.com/v1/search?format=json&key=' + apiKey + '&q=' + search
			).then((resp) => resp.json());
			return res;
		}

		return [];
	}

	let u;
	currentUserStore.subscribe((value) => {
		u = value;
	});

	onMount(async () => {
		if (browser) {
			leaflet = await import('leaflet');
			leaflet_ui = await import('leaflet-ui');

			map = leaflet.map('previewMap', {
				// Optional customizations
				mapTypeId: 'streets',
				mapTypeIds: ['streets', 'satellite'],
				gestureHandling: false,
				zoomControl: true,
				pegmanControl: false,
				locateControl: false,
				fullscreenControl: false,
				minimapControl: false,
				editInOSMControl: false,
				loadingControl: false,
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

			// limit to "single" earth map
			map.setMaxBounds([
				[105, 210],
				[-105, -190]
			]);
			map.setView([50.194036, 10.423515], 6, {
				renderer: leaflet.svg()
			});

			map.on('click', (e) => {
				if (!selectedAddress) {
					selectedAddress = {
						name: 'Marker',
						lat: e.latlng.lat,
						lon: e.latlng.lng
					};
				} else {
					selectedAddress.name = 'Marker';
					selectedAddress.lat = e.latlng.lat;
					selectedAddress.lon = e.latlng.lng;
				}
			});
		}
	});

	function submitForm(e) {
		e.preventDefault();

		if (!selectedShopName || !selectedShopName.trim() || !selectedAddress) return;

		if (!u) {
			notificationContext.addNotification({
				id: 'loginRequired',
				text: 'Für diese Funktion musst du eingeloggt sein',
				position: 'bottom-right',
				type: 'warning',
				removeAfter: 8000
			});
		}

		if (!u.activated) {
			notificationContext.addNotification({
				id: 'notActivatedWarning',
				text: 'Sorry, diese Funktion ist nur für Subs von RvNxSoul, RvNxMango oder Mahluna verfügbar. :/',
				position: 'bottom-right',
				type: 'danger',
				removeAfter: 8000
			});
			return;
		}

		fetch('/api/v1/kebabshops', {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify({
				name: selectedShopName,
				lat: Number(selectedAddress.lat),
				lng: Number(selectedAddress.lon),
				anonymous: anonymous
			})
		}).then((resp) => {
			if (resp.status == 200) {
				goto('/', { replaceState: false });
				notificationContext.addNotification({
					id: 'successAddingShop',
					text: 'Vielen Dank!',
					position: 'bottom-right',
					type: 'success',
					removeAfter: 8000
				});
				return;
			}

			notificationContext.addNotification({
				id: 'errorAddingShop',
				text: 'Leider ein Fehler aufgetreten. (' + resp.status + ' - ' + resp.statusText + ')',
				position: 'bottom-right',
				type: 'danger',
				removeAfter: 8000
			});
			return;
		});
	}

	function displayCords(value) {
		if (value && value.lat && value.lon) {
			map.setView([value.lat, value.lon], map.getZoom(), {
				renderer: leaflet.svg()
			});

			if (lastMarker) {
				map.removeLayer(lastMarker);
			}

			lastMarker = leaflet.marker([value.lat, value.lon]).addTo(map);
		}
	}
</script>

<div class="flex flex-col mx-auto container h-screen">
	<div class="card w-full h-1/3 md:h-1/2 xl:rounded-2xl xl:my-2" id="previewMap" />
	<div class="card p-4 xl:rounded-2xl">
		<h1 class="font-semibold text-lg dark:text-white">Neuen Dönerladen eintragen</h1>

		<form class="mt-2">
			<div class="mb-6">
				<label
					class="block mb-2 text-sm font-medium text-gray-900 dark:text-gray-300"
					for="shop_name">Name</label
				>
				<input
					bind:value={selectedShopName}
					class="bg-gray-50 border w-full border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block px-2 p-1 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
					id="shop_name"
					required
					type="text"
				/>
			</div>
			<div class="mb-6">
				<!--suppress XmlInvalidId -->
				<label class="block mb-2 text-sm font-medium text-gray-900 dark:text-gray-300" for="address"
					>Adresse</label
				>
				<AutoComplete
					bind:selectedItem={selectedAddress}
					className="w-full"
					delay={500}
					inputClassName="bg-gray-50 border w-full border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
					inputId="address"
					labelFieldName="display_name"
					localFiltering={false}
					maxItemsToShowInList={10}
					onChange={displayCords}
					required
					searchFunction={searchAddress}
					showLoadingIndicator={true}
				/>
			</div>
			<div class="flex items-start mb-6">
				<div class="flex items-center h-5">
					<input
						bind:checked={anonymous}
						class="w-4 h-4 bg-gray-50 rounded border border-gray-300 focus:ring-3 focus:ring-blue-300 dark:bg-gray-700 dark:border-gray-600 dark:focus:ring-blue-600 dark:ring-offset-gray-800"
						id="anonymous"
						required
						type="checkbox"
					/>
				</div>
				<label class="ml-2 text-sm font-medium text-gray-900 dark:text-gray-300" for="anonymous"
					>Anonym eintragen</label
				>
			</div>
			<button
				class="mt-2 text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm sm:w-auto px-5 py-2.5 text-center dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800"
				on:click={submitForm}
			>
				Eintragen
			</button>
		</form>
	</div>
</div>

<style>
	.card {
		@apply drop-shadow-lg;
	}
</style>
