<script>
	export let cluster;
	export let zoomCallback;
</script>

<div>
	<span>Dönerläden: {cluster.shops}</span>
	<br />
	<span class="cursor-zoom-in text-blue-500" on:click={zoomCallback(cluster.bounds)}>🔎 Zoom</span>
</div>
