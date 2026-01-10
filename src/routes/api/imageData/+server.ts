import { env } from '$env/dynamic/private';
import fs from 'fs';
import path from 'path';

const imageData = fs.readFileSync(path.resolve(env.IMAGE_DIR, "imageData.json"));

export function GET() {
	return Response.json(JSON.parse(imageData.toString()));
}