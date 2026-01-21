<script>
	import { lazyLoad } from '$lib/lazyload';
	import { env } from '$env/dynamic/public';
	import { Modal, Content, Trigger } from 'sv-popup';

	let { data } = $props();

	const cardClass = 'w-full h-auto object-scale-down rounded-lg shadow-md bg-gray-800';
</script>

<h1 class="p-8 text-center text-4xl">{env.PUBLIC_TITLE}</h1>

{#snippet photo(path_small, path_large, alt)}
	<Modal wrapper={false} big={true} button={true}>
		<Content class="m-0 items-center justify-center md:h-9/10 md:p-8">
			<img src={path_large} alt={alt} class="h-full w-full object-scale-down" style="opacity: 1;" />
		</Content>
		<Trigger>
			<img use:lazyLoad={path_small} alt={alt} class={cardClass} style="height: 100vh;" />
		</Trigger>
	</Modal>
{/snippet}

<div class="grid grid-cols-1 gap-4 md:hidden">
	{#each data.imageData as image}
		{@render photo(image.web_optimised_path, image.web_path, image.filename)}
	{/each}
</div>

<div class="hidden grid-cols-3 gap-4 md:grid">
	<div class="col-span-1">
		<div class="grid grid-cols-1 gap-4">
			{#each data.imageData as image, i}
				{#if i % 3 == 0}
					{@render photo(image.web_optimised_path, image.web_path, image.filename)}
				{/if}
			{/each}
		</div>
	</div>
	<div class="col-span-1">
		<div class="grid grid-cols-1 gap-4">
			{#each data.imageData as image, i}
				{#if i % 3 == 1}
					{@render photo(image.web_optimised_path, image.web_path, image.filename)}
				{/if}
			{/each}
		</div>
	</div>
	<div class="col-span-1">
		<div class="grid grid-cols-1 gap-4">
			{#each data.imageData as image, i}
				{#if i % 3 == 2}
					{@render photo(image.web_optimised_path, image.web_path, image.filename)}
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
