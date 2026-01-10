export async function load({ fetch }) {
	const response = await fetch('/api/imageData');
    return { imageData: await response.json() };
}