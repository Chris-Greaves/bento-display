import { env } from '$env/dynamic/private';
import fs from 'fs';
import path from 'path';

export async function load() {
    try {
        const data = fs.readFileSync(path.resolve(env.IMAGE_DIR, "imageData.json"));
        return { imageData: data.toString() ? JSON.parse(data.toString()) : [] };
    } catch (err) {
        console.error('Error reading imageData.json:', err);
        return { imageData: [] };
    }
}