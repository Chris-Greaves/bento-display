<script>
	import { lazyLoad } from '$lib/lazyload.js';

	const cardClass = 'w-full h-auto min-h-96 object-scale-down rounded-lg shadow-md bg-gray-800';
</script>

<h1 class="p-8 text-center text-4xl">Congratulations Mr & Mrs Greaves!</h1>

<div class="grid grid-cols-1 gap-4 md:hidden">
	<img use:lazyLoad={"/images/photo1.JPG"} class={cardClass} />
	<img use:lazyLoad={"/images/photo2.JPG"} class={cardClass} />
	<img use:lazyLoad={"/images/photo3.JPG"} class={cardClass} />
	<img use:lazyLoad={"/images/photo4.JPG"} class={cardClass} />
	<img use:lazyLoad={"/images/photo5.JPG"} class={cardClass} />
</div>

<div class="hidden grid-cols-3 gap-4 md:grid">
	<div class="col-span-1">
		<div class="grid grid-cols-1 gap-4">
			<img use:lazyLoad={"/images/photo0.JPG"} alt="Wedding 1" class={cardClass} />
			<img use:lazyLoad={"/images/photo3.JPG"} alt="Wedding 2" class={cardClass} />
			<img use:lazyLoad={"/images/photo6.JPG"} alt="Wedding 3" class={cardClass} />
			<img use:lazyLoad={"/images/photo9.JPG"} alt="Wedding 9" class={cardClass} />
		</div>
	</div>
	<div class="col-span-1">
		<div class="grid grid-cols-1 gap-4">
			<img use:lazyLoad={"/images/photo1.JPG"} alt="Wedding 2" class={cardClass} />
			<img use:lazyLoad={"/images/photo4.JPG"} alt="Wedding 3" class={cardClass} />
			<img use:lazyLoad={"/images/photo7.JPG"} alt="Wedding 1" class={cardClass} />
		</div>
	</div>
	<div class="col-span-1">
		<div class="grid grid-cols-1 gap-4">
			<img use:lazyLoad={"/images/photo2.JPG"} alt="Wedding 3" class={cardClass} />
			<img use:lazyLoad={"/images/photo5.JPG"} alt="Wedding 1" class={cardClass} />
			<img use:lazyLoad={"/images/photo8.JPG"} alt="Wedding 2" class={cardClass} />
		</div>
	</div>
</div>

<style>
	img {
		opacity: 0;
		transition: all 1s ease;
	}
</style>
