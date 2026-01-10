<script>
	import { lazyLoad } from '$lib/lazyload';
	import { env } from '$env/dynamic/public';

	let { data } = $props();

	console.log(data);

	const cardClass = 'w-full h-auto min-h-96 object-scale-down rounded-lg shadow-md bg-gray-800';
</script>

<h1 class="p-8 text-center text-4xl">{env.PUBLIC_TITLE}</h1>

{#snippet photo(path, alt)}
	<img use:lazyLoad={path} alt={alt} class={cardClass} />
{/snippet}

<div class="grid grid-cols-1 gap-4 md:hidden">
	{#each data.imageData as image}
		{@render photo(image.src, image.alt)}
	{/each}
</div>

<div class="hidden grid-cols-3 gap-4 md:grid">
	<div class="col-span-1">
		<div class="grid grid-cols-1 gap-4">
			{#each data.imageData as image, i}
				{#if i % 3 == 0}
					{@render photo(image.src, image.alt)}
				{/if}
			{/each}
		</div>
	</div>
	<div class="col-span-1">
		<div class="grid grid-cols-1 gap-4">
			{#each data.imageData as image, i}
				{#if i % 3 == 1}
					{@render photo(image.src, image.alt)}
				{/if}
			{/each}
		</div>
	</div>
	<div class="col-span-1">
		<div class="grid grid-cols-1 gap-4">
			{#each data.imageData as image, i}
				{#if i % 3 == 2}
					{@render photo(image.src, image.alt)}
				{/if}
			{/each}
		</div>
	</div>
</div>

<style>
	img {
		opacity: 0;
		transition: all 1s ease;
	}
</style>
