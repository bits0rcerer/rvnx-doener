<script>
	export var onChange;
	export var checked;
</script>

<div class="bg-slate-100 px-2 py-1">
	<label for="default-toggle" class="inline-flex relative items-center cursor-pointer mt-1 -ml-1">
		<input
			on:change={onChange}
			type="checkbox"
			bind:checked
			id="default-toggle"
			class="sr-only peer"
		/>
		<div
			class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 
            rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] 
            after:absolute after:top-[2px] after:left-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full 
            after:h-5 after:w-5 after:transition-all peer-checked:bg-blue-600"
		/>
		<span class="ml-2 text-xs font-medium text-gray-900">RvNx Comunity</span>
	</label>
</div>
