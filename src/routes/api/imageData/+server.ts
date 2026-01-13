import { env } from '$env/dynamic/private';
import fs from 'fs';
import path from 'path';

export async function GET() {
	// Asynchronously read 'input1.txt'
	fs.readFile(path.resolve(env.IMAGE_DIR, "imageData.json"), { encoding: 'utf8', flag: 'r' }, (err: any, data: any) => {
		if (err) {
			console.error('Error reading imageData.json:', err);
			return Response.json({imageData: []});
		} else {
			return Response.json(JSON.parse(data.toString()));
		}
	});
}