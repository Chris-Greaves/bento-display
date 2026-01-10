import { building } from '$app/environment';
import { env } from '$env/dynamic/private';
console.log(env.IMAGE_DIR);

if (!building) {
	// Generate json of all the images in the static folder for use in the app
}