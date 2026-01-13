import { env } from '$env/dynamic/private';
import fs from 'fs';
import path from 'path';

export async function GET() {
	// Synchronously read 'input2.txt'
	try {
		const data = fs.readFileSync(path.resolve(env.IMAGE_DIR, "imageData.json"));
		return Response.json(JSON.parse(data.toString()));
	} catch (err) {
			console.error('Error reading imageData.json:', err);
			return Response.json({imageData: []});
	}
}