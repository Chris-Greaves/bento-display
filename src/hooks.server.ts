import { building } from '$app/environment';
import { env } from '$env/dynamic/private';
import { imageSizeFromFile } from 'image-size/fromFile'
import type { ISizeCalculationResult } from 'image-size/types/interface';
import fs from 'fs';
import path from 'path';

console.log(env.IMAGE_DIR);

if (!building) {
	// Generate json of all the images in the static folder for use in the app
    const imageDir = path.resolve(env.IMAGE_DIR);
    const files = fs.readdirSync(imageDir);
    const images = files.filter((file: any) => {
        const ext = path.extname(file).toLowerCase();
        return ext === '.jpg' || ext === '.jpeg' || ext === '.png' || ext === '.gif';
    });
    console.log(files);
    console.log(`Found ${images.length} images in ${imageDir}`);

    // Loop through images and create an array of objects with src, alt text and dimensions
    const imageData = images.map(async (image: any) => {
        const filePath = path.join(imageDir, image);
        const dimensions = await sizeOfImage(filePath);
        return {
            src: `/images/${image}`,
            alt: `Image ${image}`,
            width: dimensions.width,
            height: dimensions.height,
            orientation: getOrientation(dimensions)
        };
    });

    // Resolve all promises
    const resolvedImageData = await Promise.all(imageData);

    console.log(resolvedImageData);

    // Write to a json file in the src/lib folder
    const outputPath = path.resolve(imageDir, 'imageData.json');
    fs.writeFileSync(outputPath, JSON.stringify(resolvedImageData, null, 2));
    console.log(`Image data written to ${outputPath}`);
}

function sizeOfImage(filePath: string) {
    return imageSizeFromFile(filePath);
}

function getOrientation(dimensions: ISizeCalculationResult) {
    if (dimensions.orientation) {
        return dimensions.orientation < 5 ? 'landscape' : 'portrait';
    }
    return dimensions.width >= dimensions.height ? 'landscape' : 'portrait';
}