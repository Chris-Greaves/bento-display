<script lang="ts">
	import './layout.css';
	import favicon from '$lib/assets/favicon.svg';

	let { children } = $props();
</script>

<svelte:head><link rel="icon" href={favicon} /></svelte:head>
<div class="min-h-screen flex flex-col items-center justify-start p-4 bg-gray-900 text-gray-100">
	{@render children()}
</div>
