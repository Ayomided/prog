// og_image_generator.ts

import { Resvg } from "npm:@resvg/resvg-js";
import { walk } from "jsr:@std/fs";
import { dirname, join } from "jsr:@std/path";

async function fetchSVG(path: string): Promise<string> {
	const url = new URL(`http://localhost:8080/og-image/${path}`);
	const response = await fetch(url);
	if (!response.ok) {
		throw new Error(`Failed to fetch SVG for ${path}: ${response.statusText}`);
	}
	return await response.text();
}

async function generateOGImage(
	svgContent: string,
	outputPath: string,
): Promise<void> {
	const opts = {
		width: 1200,
		height: 630,
		dpi: 96,
	};
	const resvg = new Resvg(svgContent, opts);
	const pngData = resvg.render();
	const pngBuffer = pngData.asPng();
	console.info("Original SVG Size:", `${resvg.width} x ${resvg.height}`);
	console.info("Output PNG Size  :", `${pngData.width} x ${pngData.height}`);

	await Deno.writeFile(outputPath, pngBuffer);
}

async function pregenerateOGImages(
	postsDir: string,
	outputDir: string,
): Promise<void> {
	for await (const entry of walk(postsDir, { exts: [".md"] })) {
		if (entry.isFile) {
			const relativePath = entry.path;
			const fileNameWithoutExt = relativePath.substring(
				0,
				relativePath.lastIndexOf("."),
			);
			const outPathName = fileNameWithoutExt.split("/")[1];

			// const decoder = new TextDecoder("utf-8");
			// const unintContent = await Deno.readFile("./static/base.svg");
			// const svgContent = decoder.decode(unintContent);
			const svgContent = await fetchSVG(outPathName);

			console.log(svgContent);

			const outputPath = join(outputDir, `${outPathName}.png`);
			await Deno.mkdir(dirname(outputPath), { recursive: true });

			await generateOGImage(svgContent, outputPath);
			console.log(`Generated OG image: ${outputPath}`);
		}
	}
}

async function main() {
	const postsDir = "./posts";
	const outputDir = "./static/og-images";

	try {
		await pregenerateOGImages(postsDir, outputDir);
		console.log("OG image pre-generation completed successfully.");
	} catch (error) {
		console.error("Error pre-generating OG images:", error);
	}
}

if (import.meta.main) {
	main();
}
